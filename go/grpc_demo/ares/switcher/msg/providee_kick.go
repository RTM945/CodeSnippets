package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type ProvideeKick struct {
	*ares.Msg
	pb *pb.ProvideeKick
}

func ProvideeKickCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewProvideeKick()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
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

func (msg *ProvideeKick) String() string {
	return fmt.Sprintf("ProvideeKick[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
