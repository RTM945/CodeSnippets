package processor

import (
	"ares/switcher"
	"ares/switcher/msg"
)

type ProvideeKickProcessor struct {
}

func NewProvideeKickProcessor() *ProvideeKickProcessor {
	return &ProvideeKickProcessor{}
}

func (p *ProvideeKickProcessor) process(provideeKick *msg.ProvideeKick) error {
	typedPB := provideeKick.TypedPB()
	linkerSession, ok := switcher.GetLinker().Sessions().GetSession(typedPB.GetClientSid()).(*switcher.LinkerSession)
	if ok && linkerSession != nil {
		_ = switcher.GetLinker().OnSessionError(linkerSession, uint32(typedPB.Reason))
		providerSession := provideeKick.GetSession()
		switcher.LOGGER.Infof("Providee kick: %v reason: %v providerSession: %v", typedPB.Reason, typedPB.Reason, providerSession)
	}
	return nil
}
