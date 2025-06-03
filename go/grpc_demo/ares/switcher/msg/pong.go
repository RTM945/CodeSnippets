package msg

import (
	"ares/pkg/io"
	linkerpb "ares/proto/switcher"
	"google.golang.org/protobuf/proto"
)

type Pong struct {
	pb      *linkerpb.Pong
	url     string
	pvId    uint32
	session io.Session
	ctx     any
}

func NewPong() *Pong {
	return &Pong{
		pb:   &linkerpb.Pong{},
		url:  "type.googleapis.com/switcher.Pong",
		pvId: 1,
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

func (msg *Pong) GetType() string { return msg.url }

func (msg *Pong) GetPvId() uint32 { return msg.pvId }

func (msg *Pong) GetContext() any { return msg.ctx }

func (msg *Pong) GetPB() proto.Message {
	return msg.pb
}

func (msg *Pong) TypedPB() *linkerpb.Pong {
	return msg.pb
}

func (msg *Pong) Dispatch() {
	msg.session.Process(msg)
}

func (msg *Pong) Process() error {
	panic("unimplemented")
}
