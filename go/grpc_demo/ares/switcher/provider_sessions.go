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
	providerSessions map[uint32]*ProviderSession
}

func NewProviderSessions(node ares.INode) *ProviderSessions {
	return &ProviderSessions{
		node:             node,
		providerSessions: make(map[uint32]*ProviderSession),
	}
}

func (pss *ProviderSessions) CheckAlive() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			pss.Lock()
			toClose := make([]*ProviderSession, 0)
			for _, providerSession := range pss.providerSessions {
				if !providerSession.Alive() {
					toClose = append(toClose, providerSession)
				} else {
					providerSession.Check()
				}
			}
			pss.Unlock()
			for _, session := range toClose {
				session.Close()
			}
		}
	}()
}

func (pss *ProviderSessions) OnAddSession(session ares.ISession) {

}

func (pss *ProviderSessions) OnRemoveSession(session ares.ISession) {
	providerSession := session.(*ProviderSession)
	pss.unbindSession(providerSession)
}

func (pss *ProviderSessions) GetSession(pvId uint32) ares.ISession {
	pss.Lock()
	defer pss.Unlock()
	return pss.providerSessions[pvId]
}

func (pss *ProviderSessions) AllSessions() []ares.ISession {
	pss.Lock()
	defer pss.Unlock()
	result := make([]ares.ISession, 0)
	for _, providerSession := range pss.providerSessions {
		result = append(result, providerSession)
	}
	return result
}

func (pss *ProviderSessions) unbindSession(providerSession *ProviderSession) {
	pss.Lock()
	defer pss.Unlock()
	pvId := providerSession.GetPvId()
	if -1 == pvId {
		return
	}
	old, ok := pss.providerSessions[uint32(pvId)]
	if !ok {
		LOGGER.Errorf("Not bound providee:%d", pvId)
	} else if old == providerSession {
		delete(pss.providerSessions, uint32(pvId))
		provider := pss.node.(*Provider)
		if providerSession.IsAUSession() {
			provider.RemoveAUPvId(uint32(pvId))
		} else if providerSession.IsPhantomSession() {
			provider.RemovePhantomPvId(uint32(pvId))
		} else if providerSession.IsGameServerSession() {
			provider.RemovePhantomGS(providerSession.GetServerId())
		}
		LOGGER.Infof("unbind providee:%v", providerSession)

		if provider.phantomPvIds.Size() == 0 {
			provideeBroken := msg.NewProvideeBroken()
			provideeBroken.TypedPB().PvId = uint32(pvId)
			provideeBroken.TypedPB().Provider = &pb.ProviderInfo{
				Ip:   provider.providerIp,
				Port: provider.port,
			}
			for _, s := range pss.providerSessions {
				s.Send(provideeBroken)
			}
		} else {
			for _, v := range provider.phantomPvIds.Snapshot() {
				provideeBroken := msg.NewProvideeBroken()
				provideeBroken.TypedPB().PvId = uint32(pvId)
				provideeBroken.TypedPB().Provider = &pb.ProviderInfo{
					Ip:   provider.providerIp,
					Port: provider.port,
				}
				provider.SendToProvidee(v, provideeBroken)
			}
		}
	}
}
