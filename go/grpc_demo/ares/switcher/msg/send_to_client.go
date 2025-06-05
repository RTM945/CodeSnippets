package msg

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
	"google.golang.org/protobuf/proto"
)

type SendToClient struct {
	pb      *pb.SendToClient
	typeId  uint32
	pvId    uint32
	session io.Session
	ctx     any
}

func SendToClientCreator(session io.Session, envelope *pb.Envelope) (io.Msg, error) {
	res := NewSendToClient()
	res.session = session
	res.ctx = envelope
	err := res.Unmarshal(envelope.Payload)
	return res, err
}

func NewSendToClient() *SendToClient {
	return &SendToClient{
		pb:     &pb.SendToClient{},
		typeId: 73,
	}
}

func (msg *SendToClient) SetSession(session io.Session) {
	msg.session = session
}

func (msg *SendToClient) GetSession() io.Session {
	return msg.session
}

func (msg *SendToClient) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *SendToClient) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *SendToClient) GetType() uint32 { return msg.typeId }

func (msg *SendToClient) GetPvId() uint32 { return msg.pvId }

func (msg *SendToClient) SetPvId(pvId uint32) {
	msg.pvId = pvId
}

func (msg *SendToClient) GetContext() any { return msg.ctx }

func (msg *SendToClient) SetContext(ctx any) { msg.ctx = ctx }

func (msg *SendToClient) GetPB() proto.Message {
	return msg.pb
}

func (msg *SendToClient) TypedPB() *pb.SendToClient {
	return msg.pb
}

func (msg *SendToClient) Dispatch() {
	msg.session.Process(msg)
}

func (msg *SendToClient) Process() error {
	// linkerSession := switcher.GetLinkerSessions().GetSession(msg.pb.ClientSid)
	// linkerSession.Send(NewMsgBox(msg.pb.TypeId, msg.pb.PvId, msg.pb.Payload))
	return nil
}
