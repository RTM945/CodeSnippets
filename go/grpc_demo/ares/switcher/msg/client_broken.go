package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type ClientBroken struct {
	*ares.Msg
	pb *pb.ClientBroken
}

func NewClientBroken() *ClientBroken {
	return &ClientBroken{
		Msg: ares.NewMsg(55),
		pb:  &pb.ClientBroken{},
	}
}

func (msg *ClientBroken) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *ClientBroken) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *ClientBroken) TypedPB() *pb.ClientBroken {
	return msg.pb
}

func (msg *ClientBroken) String() string {
	return fmt.Sprintf("ClientBroken[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
