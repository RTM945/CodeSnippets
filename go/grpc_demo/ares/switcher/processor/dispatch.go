package processor

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

func Dispatch(dispatch *msg.Dispatch) error {
	typedPB := dispatch.TypedPB()
	create, err := dispatch.GetSession().Node().MsgCreator().Create(
		dispatch.GetSession(), typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetPayload(),
	)
	if err != nil {
		provideeKick := msg.NewProvideeKick()
		provideeKick.TypedPB().ClientSid = typedPB.GetClientSid()
		provideeKick.TypedPB().Reason = pb.ProvideeKick_EXCEPTION
		_ = dispatch.GetSession().Send(provideeKick)
		ares.LOGGER.Errorf("Dispatch pvId=%d, typeId=%d, clientSid=%d", typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetClientSid())
	} else {
		// msgDebug.OnReceive(create, session)
		create.SetContext(dispatch)
		create.Dispatch()
	}

	return nil
}
