package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type SendToClient struct {
	*ares.Msg
	pb *pb.SendToClient
}

func SendToClientCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewSendToClient()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
}

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
