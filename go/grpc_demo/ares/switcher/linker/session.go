package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

// Session client<->linker
type Session struct {
	*ares.Session
	linker *Linker
}

func NewLinkerSession(stream pb.Linker_ServeServer, linker *Linker) *Session {
	return &Session{
		Session: ares.NewSession(stream),
		linker:  linker,
	}
}

func (s *Session) HandleEnvelope(envelope *pb.Envelope) {
	creator, ok := MsgCreator[envelope.TypeId]
	if ok {
		// 自己处理
		m, err := creator(s, envelope)
		if err != nil {
			LOGGER.Errorf("session[%v] unmarshal err: %v", s, err)
			return
		}
		m.Dispatch()
	} else {
		if envelope.PvId != 0 {
			// 通过PvId获取对应的服务器进行转发
			// 网关把客户端请求转发给具体的业务服处理
			toSession := s.linker.provider.GetSessions().GetSession(envelope.PvId)
			if toSession == nil {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", envelope.PvId, s, envelope.TypeId)
				serverError := msg.NewServerError()
				serverError.TypedPB().PvId = envelope.PvId
				serverError.TypedPB().Code = pb.ServerError_SERVER_NOT_ACCESSIBLE
				_ = s.Send0(serverError)
				return
			}

		}

	}
}

func (s *Session) OnClose() {

}
