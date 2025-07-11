package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"time"
)

type ProviderSession struct {
	*ares.Session
	checkToProvidee  bool
	brokenSessionIds map[uint32]int64
	aliveTime        int64
	provideeInfo     *ProvideeInfo
}

type ProvideeInfo struct {
	pvId       uint32
	serverType uint32
	serverId   uint32
	ip         int
	topics     map[string]struct{}
}

func NewProviderSession(stream pb.Provider_ServeServer, node ares.INode) *ProviderSession {
	return &ProviderSession{
		Session:          ares.NewSession(stream, node),
		brokenSessionIds: make(map[uint32]int64),
	}
}

func (s *ProviderSession) CheckToProvide() bool {
	return s.checkToProvidee
}

func (s *ProviderSession) WhiteFilter() bool {
	return s.CheckState(int(pb.ProvideeState_WHITEIP))
}

func (s *ProviderSession) BlackFilter() bool {
	return s.CheckState(int(pb.ProvideeState_BLACKIP))
}

func (s *ProviderSession) SessionBroken(brokenSessionId uint32) bool {
	if _, ok := s.brokenSessionIds[brokenSessionId]; !ok {
		s.brokenSessionIds[brokenSessionId] = time.Now().Unix()
		LOGGER.Infof("Add a broeken session, sessionId=%d", brokenSessionId)
		return true
	}
	return false
}

func (s *ProviderSession) Alive() bool {
	provider := s.Session.Node().(*Provider)
	return time.Now().Unix()-s.aliveTime < provider.sessionTimeout
}

func (s *ProviderSession) Check() {
	now := time.Now().Unix()
	provider := s.Session.Node().(*Provider)
	for k, v := range s.brokenSessionIds {
		if now-v > provider.sessionTimeout {
			LOGGER.Infof("Removed a broken session, sessionId=%d", k)
			delete(s.brokenSessionIds, k)
		}
	}
	if len(s.brokenSessionIds) > 0 {
		LOGGER.Infof("Now broken clientsids.size=%d", len(s.brokenSessionIds))
	}
}

func (ps *ProviderSession) GetPvId() int32 {
	if ps.provideeInfo == nil {
		return 0
	}
	return int32(ps.provideeInfo.pvId)
}
