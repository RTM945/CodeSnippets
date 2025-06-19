package provider

import (
	ares "ares/pkg/io"
	"sync"
)

type Sessions struct {
	sync.Mutex
	sessions map[uint32]*Session
}

func (s *Sessions) GetSession(pvId uint32) ares.ISession {
	s.Lock()
	defer s.Unlock()
	return s.sessions[pvId]
}

func (s *Sessions) Sessions() []ares.ISession {
	res := make([]ares.ISession, 0, len(s.sessions))
	s.Lock()
	defer s.Unlock()
	for _, v := range s.sessions {
		res = append(res, v)
	}
	return res
}

func (s *Sessions) OnAddSession(session ares.ISession) {

}

func (s *Sessions) OnRemoveSession(session ares.ISession) {

}
