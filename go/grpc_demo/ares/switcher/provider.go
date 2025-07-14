package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

type Provider struct {
	sessions             ares.ISessions
	msgCreator           ares.IMsgCreator
	msgProcessor         ares.IMsgProcessor
	sessionTimeout       int64
	brokenSessionTimeout int64

	pb.UnimplementedProviderServer
}

func (p *Provider) Sessions() ares.ISessions {
	return p.sessions
}

func (p *Provider) MsgCreator() ares.IMsgCreator {
	return p.msgCreator
}

func (p *Provider) MsgProcessor() ares.IMsgProcessor {
	return p.msgProcessor
}

func (p *Provider) ClientBroken(clientSid uint32, linkerSession *LinkerSession) {
	if linkerSession == nil {
		for _, v := range provider.Sessions().AllSessions() {
			providerSession := v.(*ProviderSession)
			providerSession.SessionBroken(clientSid)
		}
	} else {
		for _, pvId := range linkerSession.GetBindProvidees() {
			providerSession := provider.Sessions().GetSession(pvId)
			if providerSession != nil {
				clientBroken := msg.NewClientBroken()
				clientBroken.TypedPB().ClientSid = clientSid
				_ = providerSession.Send(clientBroken)
			}
		}
	}
}
