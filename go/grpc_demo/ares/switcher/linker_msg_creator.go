package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

type LinkerMsgCreator struct {
	registry map[uint32]ares.MsgCreatorFunc
}

func NewLinkerMsgCreator() *LinkerMsgCreator {
	msgCreator := &LinkerMsgCreator{
		registry: make(map[uint32]ares.MsgCreatorFunc),
	}

	return msgCreator
}

func (c *LinkerMsgCreator) Register(typeId uint32, f ares.MsgCreatorFunc) {
	c.registry[typeId] = f
}

func (c *LinkerMsgCreator) Create(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	creator := c.registry[typeId]
	if creator != nil {
		return creator(session, pvId, typeId, payload)
	} else {
		if pvId != 0 {
			linkerSession := session.(*LinkerSession)
			toSession, ok := provider.Sessions().GetSession(pvId).(*ProviderSession)
			if toSession == nil || !ok {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				serverError := msg.NewServerError()
				serverError.TypedPB().PvId = pvId
				serverError.TypedPB().Code = pb.ServerError_SERVER_NOT_ACCESSIBLE
				_ = linkerSession.Send0(serverError)
				return nil, nil
			}
			// 后端服务还没准备好
			if toSession.CheckToProvide() && !linkerSession.CheckState(int(pb.ClientState_TOPROVIDEE)) {
				LOGGER.Errorf("Client to Providee, state error: %d, session: %v", linkerSession.GetState(), linkerSession)
				linkerSession.CloseBySessionError(uint32(pb.SessionError_CANT_DISPATCH))
				return nil, nil
			}
			// 白名单
			if linkerSession.WhiteFilterByProvider(toSession) {
				LOGGER.Errorf("providee writeip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				linkerSession.CloseBySessionError(uint32(pb.SessionError_OPEN_WHITE_IP))
				return nil, nil
			}
			// 黑名单
			if linkerSession.BlackFilterByProvider(toSession) {
				LOGGER.Errorf("providee blackip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				linkerSession.CloseBySessionError(uint32(pb.SessionError_OPEN_BLACK_IP))
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
	}
	return nil, nil
}
