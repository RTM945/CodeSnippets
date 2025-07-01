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

func (c *MsgCreator) Create(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	creator, ok := c.registry[typeId]
	if ok {
		return creator(session, pvId, typeId, payload)
	} else {
		if pvId != 0 {
			// 通过PvId获取对应的服务器进行转发
			// 网关把客户端请求转发给具体的业务服处理
			linkerSession := session.(*Session)
			toSession, ok := provider.GetInstance().Sessions().GetSession(pvId).(*provider.Session)
			if toSession == nil || !ok {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				serverError := msg.NewServerError()
				serverError.TypedPB().PvId = pvId
				serverError.TypedPB().Code = pb.ServerError_SERVER_NOT_ACCESSIBLE
				_ = linkerSession.Send0(serverError)
				return nil, nil
			}
			linker := c.node.(*Linker)
			// 后端服务还没准备好
			if toSession.CheckToProvide() && !linkerSession.CheckState(int(pb.ClientState_TOPROVIDEE)) {
				LOGGER.Errorf("Client to Providee, state error: %d, session: %v", linkerSession.GetState(), linkerSession)
				linker.CloseSession(linkerSession, uint32(pb.SessionError_CANT_DISPATCH))
				return nil, nil
			}
			// 白名单
			if linkerSession.WhiteFilterByProvider(toSession) {
				LOGGER.Errorf("providee writeip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				linker.CloseSession(linkerSession, uint32(pb.SessionError_OPEN_WHITE_IP))
				return nil, nil
			}
			// 黑名单
			if linkerSession.BlackFilterByProvider(toSession) {
				LOGGER.Errorf("providee blackip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				linker.CloseSession(linkerSession, uint32(pb.SessionError_OPEN_BLACK_IP))
				return nil, nil
			}
			linkerSession.receiveUnknown(typeId)
			// TODO Statistics
			dispatch := msg.NewDispatch()
			dispatch.TypedPB().PvId = pvId
			dispatch.TypedPB().TypeId = typeId
			dispatch.TypedPB().Payload = payload
			err := toSession.Send(dispatch)
			if err != nil {
				LOGGER.Errorf("session[%v] sendToProvidee msg pvid: %d typeId: %d err: %v", linkerSession, pvId, typeId, err)
			}
		}
		return nil, nil
	}
}
