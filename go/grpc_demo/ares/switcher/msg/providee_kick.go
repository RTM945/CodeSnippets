package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher"
)

type ProvideeKick struct {
	*ares.Msg
	pb *pb.ProvideeKick
}

func ProvideeKickCreator(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	res := NewProvideeKick()
	res.SetSession(session)
	res.SetPvId(pvId)
	res.SetType(typeId)
	err := res.Unmarshal(payload)
	return res, err
}

var ProvideeKickProcessor = func(msg *ProvideeKick) error { panic("implement me") }

func NewProvideeKick() *ProvideeKick {
	return &ProvideeKick{
		Msg: ares.NewMsg(53),
		pb:  &pb.ProvideeKick{},
	}
}

func (msg *ProvideeKick) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *ProvideeKick) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *ProvideeKick) TypedPB() *pb.ProvideeKick {
	return msg.pb
}

func (msg *ProvideeKick) Process() error {
	typedPB := msg.TypedPB()
	linkerSession, ok := switcher.GetLinker().Sessions().GetSession(typedPB.GetClientSid()).(*switcher.LinkerSession)
	if ok && linkerSession != nil {
		_ = switcher.GetLinker().OnSessionError(linkerSession, uint32(typedPB.Reason))
		providerSession := msg.GetSession()
		switcher.LOGGER.Infof("Providee kick: %v reason: %v providerSession: %v", typedPB.Reason, typedPB.Reason, providerSession)
	}
	return nil
}
