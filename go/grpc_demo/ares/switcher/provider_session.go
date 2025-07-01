package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type ProviderSession struct {
	*ares.Session
	checkToProvidee bool
}

func NewProviderSession(stream pb.Provider_ServeServer, node ares.INode) *ProviderSession {
	return &ProviderSession{
		Session: ares.NewSession(stream, node),
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
