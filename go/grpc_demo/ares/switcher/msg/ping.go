package msg

import (
	"ares/pkg/io"
	linkerpb "ares/proto/switcher"
	"google.golang.org/protobuf/proto"
)

type Ping struct {
	pb      *linkerpb.Ping
	url     string
	pvId    uint32
	session io.Session
	ctx     any
}

func NewPing() *Ping {
	return &Ping{
		pb:   &linkerpb.Ping{},
		url:  "type.googleapis.com/switcher.Ping",
		pvId: 1,
	}
}

func (msg *Ping) SetSession(session io.Session) {
	msg.session = session
}

func (msg *Ping) GetSession() io.Session {
	return msg.session
}

func (msg *Ping) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Ping) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Ping) GetType() string { return msg.url }

func (msg *Ping) GetPvId() uint32 { return msg.pvId }

func (msg *Ping) GetContext() any { return msg.ctx }

func (msg *Ping) GetPB() proto.Message {
	return msg.pb
}

func (msg *Ping) TypedPB() *linkerpb.Ping {
	return msg.pb
}

func (msg *Ping) Dispatch() {
	msg.session.Process(msg)
}

func (msg *Ping) Process() error {
	resp := NewPong()
	resp.TypedPB().Serial = msg.TypedPB().Serial
	return msg.session.Send(resp)
}
