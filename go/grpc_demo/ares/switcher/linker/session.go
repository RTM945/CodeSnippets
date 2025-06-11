package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

// Session client<->linker
type Session struct {
	*ares.Session
}

func NewLinkerSession(stream pb.Linker_ServeServer) *Session {
	return &Session{
		Session: ares.NewSession(stream),
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
		}

	}
}

func (s *Session) OnClose() {

}
