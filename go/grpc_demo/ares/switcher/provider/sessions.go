package provider

import (
	"sync"
)

type Sessions struct {
	sync.Mutex
	sessions map[uint32]*Session
}

func (s *Sessions) GetSession(pvId uint32) *Session {
	s.Lock()
	defer s.Unlock()
	return s.sessions[pvId]
}

func (s *Sessions) Sessions() []*Session {
	res := make([]*Session, 0, len(s.sessions))
	s.Lock()
	defer s.Unlock()
	for _, v := range s.sessions {
		res = append(res, v)
	}
	return res
}
