package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

func DispatchCreator(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error) {
	res := NewDispatch()
	res.SetSession(session)
	res.SetContext(envelope)
	err := res.Unmarshal(envelope.Payload)
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
	return nil
}
