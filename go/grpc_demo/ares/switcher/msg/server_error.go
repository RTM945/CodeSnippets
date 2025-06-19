package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type ServerError struct {
	*ares.Msg
	pb *pb.ServerError
}

func NewServerError() *ServerError {
	return &ServerError{
		Msg: ares.NewMsg(11),
		pb:  &pb.ServerError{},
	}
}

func (msg *ServerError) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *ServerError) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *ServerError) TypedPB() *pb.ServerError {
	return msg.pb
}
