package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher"
)

type SendToClient struct {
	*ares.Msg
	pb *pb.SendToClient
}

func SendToClientCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewSendToClient()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
}

func NewSendToClient() *SendToClient {
	return &SendToClient{
		Msg: ares.NewMsg(73),
		pb:  &pb.SendToClient{},
	}
}

func (msg *SendToClient) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *SendToClient) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *SendToClient) TypedPB() *pb.SendToClient {
	return msg.pb
}

func (msg *SendToClient) Process() error {
	typedPB := msg.TypedPB()
	linkerSession, ok := switcher.GetLinker().Sessions().GetSession(typedPB.GetClientSid()).(*switcher.LinkerSession)
	if ok && linkerSession != nil {
		msgBox := NewMsgBox()
		msgBox.TypedPB().TypeId = typedPB.TypeId
		msgBox.TypedPB().Payload = typedPB.Payload
		return linkerSession.Send(msgBox)
	} else {
		// provide client broken
	}
	return nil
}
