package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

func DispatchCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewDispatch()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
}

type Dispatch struct {
	*ares.Msg
	pb *pb.Dispatch
}

func NewDispatch() *Dispatch {
	return &Dispatch{
		Msg: ares.NewMsg(51),
		pb:  &pb.Dispatch{},
	}
}

func (msg *Dispatch) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Dispatch) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Dispatch) TypedPB() *pb.Dispatch {
	return msg.pb
}

func (msg *Dispatch) Process() error {
	typedPB := msg.TypedPB()
	create, err := msg.GetSession().Node().MsgCreator().Create(
		msg.GetSession(), typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetPayload(),
	)
	if err != nil {
		provideeKick := NewProvideeKick()
		provideeKick.TypedPB().ClientSid = typedPB.GetClientSid()
		provideeKick.TypedPB().Reason = pb.ProvideeKick_EXCEPTION
		_ = msg.GetSession().Send(provideeKick)
		ares.LOGGER.Errorf("Process pvId=%d, typeId=%d, clientSid=%d", typedPB.GetPvId(), typedPB.GetTypeId(), typedPB.GetClientSid())
	} else {
		// msgDebug.OnReceive(create, session)
		create.SetContext(msg)
		create.Dispatch()
	}

	return nil
}
