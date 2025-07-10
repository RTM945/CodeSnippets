package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"errors"
)

type LinkerMsgCreator struct {
	*ares.MsgCreator
}

func NewLinkerMsgCreator() *LinkerMsgCreator {
	msgCreator := &LinkerMsgCreator{
		MsgCreator: ares.NewMsgCreator(),
	}

	return msgCreator
}

func (c *LinkerMsgCreator) Create(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	create, err := c.MsgCreator.Create(session, pvId, typeId, payload)
	if err == nil {
		return create, nil
	} else if errors.Is(err, ares.NoMsgCreatorErr) {
		if pvId != 0 {
			linkerSession := session.(*LinkerSession)
			toSession, ok := provider.Sessions().GetSession(pvId).(*ProviderSession)
			if toSession == nil || !ok {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				_ = linker.OnServerError(linkerSession, pvId, uint32(pb.ServerError_SERVER_NOT_ACCESSIBLE))
				return nil, nil
			}
			// 后端服务还没准备好
			if toSession.CheckToProvide() && !linkerSession.CheckState(int(pb.ClientState_TOPROVIDEE)) {
				LOGGER.Errorf("Client to Providee, state error: %d, session: %v", linkerSession.GetState(), linkerSession)
				_ = linker.OnSessionError(linkerSession, uint32(pb.SessionError_CANT_DISPATCH))
				return nil, nil
			}
			// 白名单
			if linkerSession.WhiteFilterByProvider(toSession) {
				LOGGER.Errorf("providee writeip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				_ = linker.OnSessionError(linkerSession, uint32(pb.SessionError_OPEN_WHITE_IP))
				return nil, nil
			}
			// 黑名单
			if linkerSession.BlackFilterByProvider(toSession) {
				LOGGER.Errorf("providee blackip kick, pvid: %d, session: %v, typeId: %d", pvId, linkerSession, typeId)
				_ = linker.OnSessionError(linkerSession, uint32(pb.SessionError_OPEN_BLACK_IP))
				return nil, nil
			}
			linkerSession.receiveUnknown(typeId)
			// TODO Statistics
			err := linker.OnDispatch(toSession, pvId, typeId, payload)
			if err != nil {
				LOGGER.Errorf("session[%v] sendToProvidee msg pvid: %d typeId: %d err: %v", linkerSession, pvId, typeId, err)
			}
		}
	}
	return nil, err
}
