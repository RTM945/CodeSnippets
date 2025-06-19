package provider

import (
	"ares/logger"
	ares "ares/pkg/io"
)

var LOGGER = logger.GetLogger("provider")

type Provider struct {
	sessions *Sessions
}

func (p *Provider) GetSessions() ares.ISessions {
	return p.sessions
}
