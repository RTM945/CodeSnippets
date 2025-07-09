package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type LinkerSessions struct {
	*ares.Sessions
	node ares.INode
}

func NewLinkerSessions(node ares.INode) *LinkerSessions {
	return &LinkerSessions{
		Sessions: ares.NewSessions(),
		node:     node,
	}
}

func (ls *LinkerSessions) OnAddSession(session ares.ISession) {
	if ls.Size() >= uint32(maxSession) {
		linkerSession := session.(*LinkerSession)
		linkerSession.CloseBySessionError(pb.SessionError_OVER_MAX_SESSIONS)
		return
	}
	ls.Sessions.OnAddSession(session)
}

func (ls *LinkerSessions) OnRemoveSession(session ares.ISession) {
	linkerSession := session.(*LinkerSession)
	linkerSession.OnClose()
}
