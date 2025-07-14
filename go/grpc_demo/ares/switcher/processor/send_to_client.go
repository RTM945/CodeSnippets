package processor

import (
	"ares/switcher"
	"ares/switcher/msg"
)

type SendToClientProcessor struct {
}

func NewSendToClientProcessor() *SendToClientProcessor {
	return &SendToClientProcessor{}
}

func (p *SendToClientProcessor) Process(sendToClient *msg.SendToClient) error {
	typedPB := sendToClient.TypedPB()
	linkerSession, ok := switcher.GetLinker().Sessions().GetSession(typedPB.GetClientSid()).(*switcher.LinkerSession)
	if ok && linkerSession != nil {
		msgBox := msg.NewMsgBox()
		msgBox.TypedPB().TypeId = typedPB.TypeId
		msgBox.TypedPB().Payload = typedPB.Payload
		return linkerSession.Send(msgBox)
	} else {
		switcher.GetProvider().ClientBroken(typedPB.GetClientSid(), linkerSession)
	}
	return nil
}
