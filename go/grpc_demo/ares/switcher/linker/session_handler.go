package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type SessionHandler struct {
	linker *Linker
}

func NewSessionHandler(linker *Linker) *SessionHandler {
	return &SessionHandler{
		linker: linker,
	}
}

func (sh *SessionHandler) OnAddSession(session ares.ISession) {
	linkerSession := session.(*Session)
	if !sh.linker.CanAddSession() {
		sh.linker.CloseSession(linkerSession, pb.SessionError_OVER_MAX_SESSIONS)
		return
	}
	sh.linker.sessions.AddSession(linkerSession)
}

func (sh *SessionHandler) OnRemoveSession(session ares.ISession) {
	linkerSession := session.(*Session)
	linkerSession.OnClose()
	sh.linker.sessions.RemoveSession(linkerSession)

}

func (sh *SessionHandler) Stop() {
	sh.linker.sessions.Stop()
}
