package msg

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
	"google.golang.org/protobuf/proto"
)

type Ping struct {
	pb      *pb.Ping
	typeId  uint32
	pvId    uint32
	session io.Session
	ctx     any
}

func PingCreator(session io.Session, envelope *pb.Envelope) (io.Msg, error) {
	res := NewPing()
	res.session = session
	res.ctx = envelope
	err := res.Unmarshal(envelope.Payload)
	return res, err
}

func NewPing() *Ping {
	return &Ping{
		pb:     &pb.Ping{},
		typeId: 4,
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

func (msg *Ping) GetType() uint32 { return msg.typeId }

func (msg *Ping) GetPvId() uint32 { return msg.pvId }

func (msg *Ping) SetPvId(pvId uint32) {
	msg.pvId = pvId
}

func (msg *Ping) GetContext() any { return msg.ctx }

func (msg *Ping) SetContext(ctx any) { msg.ctx = ctx }

func (msg *Ping) GetPB() proto.Message {
	return msg.pb
}

func (msg *Ping) TypedPB() *pb.Ping {
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
