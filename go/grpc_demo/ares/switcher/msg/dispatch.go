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
