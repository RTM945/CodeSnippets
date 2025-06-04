package msg

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
	"google.golang.org/protobuf/proto"
)

type Pong struct {
	pb      *pb.Pong
	typeId  uint32
	pvId    uint32
	session io.Session
	ctx     any
}

func NewPong() *Pong {
	return &Pong{
		pb:     &pb.Pong{},
		typeId: 8,
	}
}

func (msg *Pong) SetSession(session io.Session) {
	msg.session = session
}

func (msg *Pong) GetSession() io.Session {
	return msg.session
}

func (msg *Pong) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Pong) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Pong) GetType() uint32 { return msg.typeId }

func (msg *Pong) GetPvId() uint32 { return msg.pvId }

func (msg *Pong) SetPvId(pvId uint32) {
	msg.pvId = pvId
}

func (msg *Pong) GetContext() any { return msg.ctx }

func (msg *Pong) SetContext(ctx any) { msg.ctx = ctx }

func (msg *Pong) GetPB() proto.Message {
	return msg.pb
}

func (msg *Pong) TypedPB() *pb.Pong {
	return msg.pb
}

func (msg *Pong) Dispatch() {
	msg.session.Process(msg)
}

func (msg *Pong) Process() error {
	panic("unimplemented")
}
