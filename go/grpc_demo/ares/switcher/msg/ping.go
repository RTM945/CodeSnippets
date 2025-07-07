package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type Ping struct {
	*ares.Msg
	pb *pb.Ping
}

func PingCreator(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error) {
	res := NewPing()
	res.SetSession(session)
	res.SetContext(envelope)
	err := res.Unmarshal(envelope.GetPayload())
	return res, err
}

func NewPing() *Ping {
	return &Ping{
		Msg: ares.NewMsg(4),
		pb:  &pb.Ping{},
	}
}

func (msg *Ping) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Ping) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Ping) TypedPB() *pb.Ping {
	return msg.pb
}

func (msg *Ping) Process() error {
	resp := NewPong()
	resp.TypedPB().Serial = msg.TypedPB().Serial
	return msg.GetSession().Send(resp)
}
