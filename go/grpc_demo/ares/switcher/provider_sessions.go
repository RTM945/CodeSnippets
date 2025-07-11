package switcher

import (
	ares "ares/pkg/io"
	"sync"
	"time"
)

type ProviderSessions struct {
	*ares.Sessions
	node ares.INode
	sync.Mutex
	providerSessions map[uint32]*ProviderSession
}

func NewProviderSessions(node ares.INode) *ProviderSessions {
	return &ProviderSessions{
		Sessions:         ares.NewSessions(),
		node:             node,
		providerSessions: make(map[uint32]*ProviderSession),
	}
}

func (ps *ProviderSessions) CheckAlive() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			ps.Lock()
			toClose := make([]*ProviderSession, 0, ps.Size())
			for _, providerSession := range ps.providerSessions {
				if !providerSession.Alive() {
					toClose = append(toClose, providerSession)
				} else {
					providerSession.Check()
				}
			}
			ps.Unlock()
			for _, session := range toClose {
				session.Close()
			}
		}
	}()
}

func (ps *ProviderSessions) OnAddSession(session ares.ISession) {

}

func (ps *ProviderSessions) OnRemoveSession(session ares.ISession) {
	providerSession := session.(*ProviderSession)
	ps.unbindSession(providerSession)
}

func (ps *ProviderSessions) unbindSession(providerSession *ProviderSession) {
	ps.Lock()
	defer ps.Unlock()
	pvId := providerSession.GetPvId()
	if -1 == pvId {
		return
	}
	old, ok := ps.providerSessions[uint32(pvId)]
	if !ok {
		LOGGER.Errorf("Not bound providee:%d", pvId)
	} else if old == providerSession {
		delete(ps.providerSessions, uint32(pvId))
		provider := ps.node.(*Provider)

	}
}
