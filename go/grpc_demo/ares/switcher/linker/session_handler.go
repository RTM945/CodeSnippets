package linker

import (
	ares "ares/pkg/io"
)

type SessionHandler struct {
	linker *Linker
}

func NewSessionHandler(linker *Linker) *SessionHandler {
	return &SessionHandler{
		linker: linker,
	}
}

func (sh *SessionHandler) OnAddSession(session ares.Session) {
	linkerSession := session.(*Session)

	sh.linker.sessions.AddSession(linkerSession)
}

func (sh *SessionHandler) OnRemoveSession(session ares.Session) {
	linkerSession := session.(*Session)
	linkerSession.OnClose()
	sh.linker.sessions.RemoveSession(linkerSession)

}

func (sh *SessionHandler) Stop() {
	sh.linker.sessions.Stop()
}
