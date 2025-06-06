package provider

import "ares/logger"

var LOGGER = logger.GetLogger("provider")

type Provider struct {
	sessions *Sessions
}

func (p *Provider) GetSessions() *Sessions {
	return p.sessions
}
