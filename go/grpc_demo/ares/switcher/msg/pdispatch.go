package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type PDispatch struct {
	*ares.Msg
	pb *pb.PDispatch
}

func NewPDispatch() *PDispatch {
	return &PDispatch{
		Msg: ares.NewMsg(77),
		pb:  &pb.PDispatch{},
	}
}

func (msg *PDispatch) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *PDispatch) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *PDispatch) TypedPB() *pb.PDispatch {
	return msg.pb
}
