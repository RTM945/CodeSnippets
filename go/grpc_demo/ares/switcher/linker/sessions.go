package linker

import (
	"sync"
	"sync/atomic"
)

var sessionCNT uint32

type Sessions struct {
	sync.Mutex
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
	s.Lock()
	defer s.Unlock()
	return s.sessions[sid]
}

func (s *Sessions) RemoveSession(sid uint32) {
	s.Lock()
	defer s.Unlock()
	atomic.AddUint32(&sessionCNT, -1)
	delete(s.sessions, sid)
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
