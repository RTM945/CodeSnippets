package processor

import (
	"ares/switcher"
	"ares/switcher/msg"
)

func ProvideeKick(kick *msg.ProvideeKick) error {
	typedPB := kick.TypedPB()
	linkerSession, ok := switcher.GetLinker().Sessions().GetSession(typedPB.GetClientSid()).(*switcher.LinkerSession)
	if ok && linkerSession != nil {
		_ = switcher.GetLinker().OnSessionError(linkerSession, uint32(typedPB.Reason))
		providerSession := kick.GetSession()
		switcher.LOGGER.Infof("Providee kick: %v reason: %v providerSession: %v", typedPB.Reason, typedPB.Reason, providerSession)
	}
	return nil
}
