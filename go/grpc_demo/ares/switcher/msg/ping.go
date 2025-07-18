package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type Ping struct {
	*ares.Msg
	pb *pb.Ping
}

func PingCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewPing()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
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
