package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"time"
)

type LinkerSessions struct {
	*ares.Sessions
}

func NewLinkerSessions() *LinkerSessions {
	return &LinkerSessions{
		Sessions: ares.NewSessions(),
	}
}

func (ls *LinkerSessions) OnAddSession(session ares.ISession) {
	linker := linker.(*Linker)
	if ls.Size() >= linker.maxSession {
		linkerSession := session.(*LinkerSession)
		linkerSession.CloseBySessionError(pb.SessionError_OVER_MAX_SESSIONS)
		return
	}
	ls.Sessions.OnAddSession(session)
}

func (ls *LinkerSessions) OnRemoveSession(session ares.ISession) {
	linkerSession := session.(*LinkerSession)
	linkerSession.OnClose()
	ls.Sessions.OnRemoveSession(session)
}

func (ls *LinkerSessions) StartCheck() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			ls.RLock()
			toClose := make([]ares.ISession, 0, ls.Size())
			for _, s := range ls.Sessions.Sessions() {
				linkerSession := s.(*LinkerSession)
				if !linkerSession.Alive() {
					toClose = append(toClose, linkerSession)
				}
			}
			ls.RUnlock()
			for _, session := range toClose {
				session.Close()
			}
		}
	}()
}
