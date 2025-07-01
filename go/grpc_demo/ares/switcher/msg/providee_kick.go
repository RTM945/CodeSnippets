package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type ProvideeKick struct {
	*ares.Msg
	pb *pb.ProvideeKick
}

func NewProvideeKick() *ProvideeKick {
	return &ProvideeKick{
		Msg: ares.NewMsg(53),
		pb:  &pb.ProvideeKick{},
	}
}

func (msg *ProvideeKick) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *ProvideeKick) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *ProvideeKick) TypedPB() *pb.ProvideeKick {
	return msg.pb
}

func (msg *ProvideeKick) Process() error {
	msg.GetSession().Node().Sessions()
	sessionErr := NewSessionError()
	sessionErr.TypedPB().Code = uint32(msg.TypedPB().Reason)
	//_ = session.Send0(sessionErr)
	//session.Close()
	return nil
}
