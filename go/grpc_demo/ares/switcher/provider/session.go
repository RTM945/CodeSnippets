package provider

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

// Session linker<->provider
type Session struct {
	*ares.Session
	checkToProvidee bool
}

func NewProviderSession(stream pb.Provider_ServeServer) *Session {
	return &Session{
		Session: ares.NewSession(stream),
	}
}

func (s *Session) CheckToProvide() bool {
	return s.checkToProvidee
}

func (s *Session) WhiteFilter() bool {
	return s.CheckState(int(pb.ProvideeState_WHITEIP))
}

func (s *Session) BlackFilter() bool {
	return s.CheckState(int(pb.ProvideeState_BLACKIP))
}
