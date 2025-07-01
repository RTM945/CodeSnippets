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
	dispatch := msg.TypedPB()
	create, err := msg.GetSession().Node().MsgCreator().Create(
		msg.GetSession(), dispatch.GetPvId(), dispatch.GetTypeId(), dispatch.GetPayload(),
	)
	if err != nil {
		provideeKick := NewProvideeKick()
		provideeKick.TypedPB().ClientSid = msg.TypedPB().GetClientSid()
		provideeKick.TypedPB().Reason = pb.ProvideeKick_EXCEPTION
		msg.GetSession().Send(provideeKick)
		ares.LOGGER.Errorf("Dispatch pvId=%d, typeId=%d, clientSid=%d", dispatch.GetPvId(), dispatch.GetTypeId(), dispatch.GetClientSid())
	} else {
		// msgDebug.OnReceive(create, session)
		create.SetContext(msg)
		create.Dispatch()
	}

	return nil
}
