package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type Provider struct {
	sessions             ares.ISessions
	msgCreator           ares.IMsgCreator
	sessionTimeout       int64
	brokenSessionTimeout int64

	OnClientBroken func(clientSid uint32, linkerSession *LinkerSession)

	pb.UnimplementedProviderServer
}

func (p *Provider) Sessions() ares.ISessions {
	return p.sessions
}

func (p *Provider) MsgCreator() ares.IMsgCreator {
	return p.msgCreator
}
