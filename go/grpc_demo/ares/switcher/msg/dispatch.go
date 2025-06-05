package msg

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
	"google.golang.org/protobuf/proto"
)

type Dispatch struct {
	pb      *pb.Dispatch
	typeId  uint32
	pvId    uint32
	session io.Session
	ctx     any
}

func NewDispatch() *Dispatch {
	return &Dispatch{
		pb:     &pb.Dispatch{},
		typeId: 51,
	}
}

func (msg *Dispatch) SetSession(session io.Session) {
	msg.session = session
}

func (msg *Dispatch) GetSession() io.Session {
	return msg.session
}

func (msg *Dispatch) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *Dispatch) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *Dispatch) GetType() uint32 { return msg.typeId }

func (msg *Dispatch) GetPvId() uint32 { return msg.pvId }

func (msg *Dispatch) SetPvId(pvId uint32) {
	msg.pvId = pvId
}

func (msg *Dispatch) GetContext() any { return msg.ctx }

func (msg *Dispatch) SetContext(ctx any) { msg.ctx = ctx }

func (msg *Dispatch) GetPB() proto.Message {
	return msg.pb
}

func (msg *Dispatch) TypedPB() *pb.Dispatch {
	return msg.pb
}

func (msg *Dispatch) Dispatch() {
	msg.session.Process(msg)
}

func (msg *Dispatch) Process() error {
	return nil
}
