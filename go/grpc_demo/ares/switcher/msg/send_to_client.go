package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type SendToClient struct {
	*ares.Msg
	pb *pb.SendToClient
}

func SendToClientCreator(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error) {
	res := NewSendToClient()
	res.SetSession(session)
	res.SetContext(envelope)
	err := res.Unmarshal(envelope.GetPayload())
	return res, err
}

var SendToClientProcessor = func(msg *SendToClient) error { panic("implement me") }

func NewSendToClient() *SendToClient {
	return &SendToClient{
		Msg: ares.NewMsg(73),
		pb:  &pb.SendToClient{},
	}
}

func (msg *SendToClient) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *SendToClient) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *SendToClient) TypedPB() *pb.SendToClient {
	return msg.pb
}

func (msg *SendToClient) Process() error {
	return SendToClientProcessor(msg)
}
