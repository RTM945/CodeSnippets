package provider

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

// Session linker<->provider
type Session struct {
	*ares.Session
}

func NewProviderSession(stream pb.Provider_ServeServer) *Session {
	return &Session{
		Session: ares.NewSession(stream),
	}
}
