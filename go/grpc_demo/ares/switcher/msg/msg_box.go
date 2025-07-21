package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type MsgBox struct {
	*ares.Msg
	pb *pb.MsgBox
}

func NewMsgBox() *MsgBox {
	return &MsgBox{
		Msg: ares.NewMsg(3),
		pb:  &pb.MsgBox{},
	}
}

func (msg *MsgBox) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *MsgBox) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *MsgBox) TypedPB() *pb.MsgBox {
	return msg.pb
}

func (msg *MsgBox) String() string {
	return fmt.Sprintf("MsgBox[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
