package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"sync"
	"sync/atomic"
)

var sessionCNT uint32

type Sessions struct {
	sync.RWMutex
	linker   *Linker
	sessions map[uint32]*Session
}

func NewSessions(linker *Linker) *Sessions {
	return &Sessions{
		linker:   linker,
		sessions: make(map[uint32]*Session),
	}
}

func (s *Sessions) GetSession(sid uint32) ares.ISession {
	s.RLock()
	defer s.RUnlock()
	return s.sessions[sid]
}

func (s *Sessions) Sessions() []ares.ISession {
	res := make([]ares.ISession, 0, len(s.sessions))
	s.RLock()
	defer s.RUnlock()
	for _, v := range s.sessions {
		res = append(res, v)
	}
	return res
}

func (s *Sessions) Size() uint32 {
	return sessionCNT
}

func (s *Sessions) OnAddSession(session ares.ISession) {
	linkerSession := session.(*Session)
	if !s.linker.CanAddSession() {
		s.linker.CloseSession(linkerSession, pb.SessionError_OVER_MAX_SESSIONS)
		return
	}
	s.Lock()
	defer s.Unlock()
	atomic.AddUint32(&sessionCNT, 1)
	s.sessions[session.GetSid()] = linkerSession
}

func (s *Sessions) OnRemoveSession(session ares.ISession) {
	linkerSession := session.(*Session)
	linkerSession.OnClose()
	s.Lock()
	defer s.Unlock()
	atomic.AddUint32(&sessionCNT, -1)
	delete(s.sessions, session.GetSid())
}

func (s *Sessions) Stop() {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.sessions {
		v.Close()
	}
	clear(s.sessions)
}
