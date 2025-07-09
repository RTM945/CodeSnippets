package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"ares/switcher/provider"
)

type MsgCreator struct {
	node     ares.INode
	registry map[uint32]ares.MsgCreatorFunc
}

func NewMsgCreator(node ares.INode) *MsgCreator {
	msgCreator := &MsgCreator{
		node:     node,
		registry: make(map[uint32]ares.MsgCreatorFunc),
	}
	initMsg(msgCreator)
	return msgCreator
}

func (c *MsgCreator) Register(id uint32, f ares.MsgCreatorFunc) {
	c.registry[id] = f
}

func (c *MsgCreator) Create(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error) {
	creator, ok := c.registry[envelope.GetTypeId()]
	if ok {
		return creator(session, envelope)
	} else {
		if envelope.GetPvId() != 0 {
			// 通过PvId获取对应的服务器进行转发
			// 网关把客户端请求转发给具体的业务服处理
			linkerSession := session.(*Session)
			toSession, ok := provider.GetInstance().Sessions().GetSession(envelope.GetPvId()).(*provider.Session)
			if toSession == nil || !ok {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", envelope.GetPvId(), linkerSession, envelope.GetTypeId())
				serverError := msg.NewServerError()
				serverError.TypedPB().PvId = envelope.GetPvId()
				serverError.TypedPB().Code = pb.ServerError_SERVER_NOT_ACCESSIBLE
				_ = linkerSession.Send0(serverError)
				return nil, nil
			}
			linker := c.node.(*Linker)
			// 后端服务还没准备好
			if toSession.CheckToProvide() && !linkerSession.CheckState(int(pb.ClientState_TOPROVIDEE)) {
				LOGGER.Errorf("Client to Providee, state error: %d, session: %v", linkerSession.GetState(), linkerSession)
				linker.CloseSession(linkerSession, pb.SessionError_CANT_DISPATCH)
				return nil, nil
			}
			// 白名单
			if linkerSession.WhiteFilterByProvider(toSession) {
				LOGGER.Errorf("providee writeip kick, pvid: %d, session: %v, typeId: %d", envelope.GetPvId(), linkerSession, envelope.GetTypeId())
				linker.CloseSession(linkerSession, pb.SessionError_OPEN_WHITE_IP)
				return nil, nil
			}
			// 黑名单
			if linkerSession.BlackFilterByProvider(toSession) {
				LOGGER.Errorf("providee blackip kick, pvid: %d, session: %v, typeId: %d", envelope.GetPvId(), linkerSession, envelope.GetTypeId())
				linker.CloseSession(linkerSession, pb.SessionError_OPEN_BLACK_IP)
				return nil, nil
			}
			linkerSession.receiveUnknown(envelope.GetTypeId())
			// TODO Statistics
			dispatch := msg.NewDispatch()
			dispatch.TypedPB().PvId = envelope.GetPvId()
			dispatch.TypedPB().TypeId = envelope.GetTypeId()
			dispatch.TypedPB().Payload = envelope.GetPayload()
			err := toSession.Send(dispatch)
			if err != nil {
				LOGGER.Errorf("session[%v] sendToProvidee msg=%v err: %v", linkerSession, envelope, err)
			}
		}
		return nil, nil
	}
}
