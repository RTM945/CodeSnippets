package provider

import (
	"ares/logger"
	ares "ares/pkg/io"
)

var LOGGER = logger.GetLogger("provider")

var provider = &Provider{
	sessions:   nil,
	msgCreator: nil,
}

func GetInstance() *Provider {
	return provider
}

type Provider struct {
	sessions   *Sessions
	msgCreator ares.IMsgCreator
}

func (p *Provider) Sessions() ares.ISessions {
	return p.sessions
}

func (p *Provider) MsgCreator() ares.IMsgCreator {
	return p.msgCreator
}
