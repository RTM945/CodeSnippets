package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type ProvideeBroken struct {
	*ares.Msg
	pb *pb.ProvideeBroken
}

func ProvideeBrokenCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewProvideeBroken()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
}

func NewProvideeBroken() *ProvideeBroken {
	return &ProvideeBroken{
		Msg: ares.NewMsg(63),
		pb:  &pb.ProvideeBroken{},
	}
}

func (msg *ProvideeBroken) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *ProvideeBroken) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *ProvideeBroken) TypedPB() *pb.ProvideeBroken {
	return msg.pb
}

func (msg *ProvideeBroken) String() string {
	return fmt.Sprintf("ProvideeBroken[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
