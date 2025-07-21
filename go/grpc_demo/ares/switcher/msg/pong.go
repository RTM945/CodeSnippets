package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type Pong struct {
	*ares.Msg
	pb *pb.Pong
}

func NewPong() *Pong {
	return &Pong{
		Msg: ares.NewMsg(8),
		pb:  &pb.Pong{},
	}
}

func (msg *Pong) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Pong) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Pong) TypedPB() *pb.Pong {
	return msg.pb
}

func (msg *Pong) String() string {
	return fmt.Sprintf("Pong[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
