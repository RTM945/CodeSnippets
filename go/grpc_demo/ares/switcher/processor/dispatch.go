package processor

import (
	pb "ares/proto/gen"
	"ares/switcher"
	"ares/switcher/msg"
)

type DispatchProcessor struct {
}

func NewDispatchProcessor() *DispatchProcessor {
	return &DispatchProcessor{}
}

func (p *DispatchProcessor) Process(dispatch *msg.Dispatch) error {
	typedPB := dispatch.TypedPB()
	create, err := dispatch.GetSession().Node().MsgCreator().Create(
		dispatch.GetSession(), typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetPayload(),
	)
	if err != nil {
		provideeKick := msg.NewProvideeKick()
		provideeKick.TypedPB().ClientSid = typedPB.GetClientSid()
		provideeKick.TypedPB().Reason = pb.ProvideeKick_EXCEPTION
		_ = dispatch.GetSession().Send(provideeKick)
		switcher.LOGGER.Errorf("Process pvId=%d, typeId=%d, clientSid=%d", typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetClientSid())
	} else {
		// msgDebug.OnReceive(create, session)
		create.SetContext(dispatch)
		create.Dispatch()
	}

	return nil
}
