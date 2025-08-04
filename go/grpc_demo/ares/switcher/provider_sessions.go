package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"sync"
	"time"
)

type ProviderSessions struct {
	node ares.INode
	sync.Mutex
	providerSessions sync.Map
}

func NewProviderSessions(node ares.INode) *ProviderSessions {
	return &ProviderSessions{
		node: node,
	}
}

func (pss *ProviderSessions) CheckAlive() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			pss.providerSessions.Range(func(k, v interface{}) bool {
				providerSession := v.(*ProviderSession)
				if !providerSession.Alive() {
					providerSession.Close()
				} else {
					providerSession.Check()
				}
				return true
			})
		}
	}()
}

func (pss *ProviderSessions) OnAddSession(session ares.ISession) {

}

func (pss *ProviderSessions) OnRemoveSession(session ares.ISession) {
	providerSession := session.(*ProviderSession)
	pss.unbindSession(providerSession)
}

func (pss *ProviderSessions) GetSession(pvId int32) ares.ISession {
	value, ok := pss.providerSessions.Load(pvId)
	if !ok {
		return nil
	}
	return value.(*ProviderSession)
}

func (pss *ProviderSessions) AllSessions() []ares.ISession {
	pss.Lock()
	defer pss.Unlock()
	result := make([]ares.ISession, 0)
	pss.providerSessions.Range(func(k, v interface{}) bool {
		result = append(result, v.(*ProviderSession))
		return true
	})

	return result
}

func (pss *ProviderSessions) unbindSession(providerSession *ProviderSession) {
	pss.Lock()
	defer pss.Unlock()
	pvId := providerSession.GetPvId()
	if -1 == pvId {
		return
	}
	old, ok := pss.providerSessions.Load(pvId)
	if !ok {
		LOGGER.Errorf("Not bound providee:%d", pvId)
	} else if old == providerSession {
		pss.providerSessions.Delete(pvId)
		provider := pss.node.(*Provider)
		if providerSession.IsAUSession() {
			provider.RemoveAUPvId(pvId)
		} else if providerSession.IsPhantomSession() {
			provider.RemovePhantomPvId(pvId)
		} else if providerSession.IsGameServerSession() {
			provider.RemovePhantomGS(providerSession.GetServerId())
		}
		LOGGER.Infof("unbind providee:%v", providerSession)
		if !provider.ProvideeBroken(pvId) {

			pss.providerSessions.Range(func(k, v interface{}) bool {
				providerSession := v.(*ProviderSession)
				provideeBroken := msg.NewProvideeBroken()
				provideeBroken.TypedPB().PvId = pvId
				provideeBroken.TypedPB().Provider = &pb.ProviderInfo{
					Ip:   provider.providerIp,
					Port: provider.port,
				}
				providerSession.Send(provideeBroken)
				return true
			})
		}
	}
}
