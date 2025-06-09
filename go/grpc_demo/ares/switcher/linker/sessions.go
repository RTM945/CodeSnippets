package linker

import (
	"sync"
	"sync/atomic"
)

var sessionCNT uint32

type Sessions struct {
	sync.RWMutex
	sessions map[uint32]*Session
}

func NewSessions() *Sessions {
	return &Sessions{
		sessions: make(map[uint32]*Session),
	}
}

func (s *Sessions) AddSession(session *Session) {
	s.Lock()
	defer s.Unlock()
	atomic.AddUint32(&sessionCNT, 1)
	s.sessions[session.GetSid()] = session
}

func (s *Sessions) GetSession(sid uint32) *Session {
	s.RLock()
	defer s.RUnlock()
	return s.sessions[sid]
}

func (s *Sessions) RemoveSession(session *Session) {
	s.Lock()
	defer s.Unlock()
	atomic.AddUint32(&sessionCNT, -1)
	delete(s.sessions, session.GetSid())
}

func (s *Sessions) Sessions() []*Session {
	res := make([]*Session, 0, len(s.sessions))
	s.RLock()
	defer s.RUnlock()
	for _, v := range s.sessions {
		res = append(res, v)
	}
	return res
}

func (s *Sessions) Stop() {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.sessions {
		v.Close()
	}
	clear(s.sessions)
}
