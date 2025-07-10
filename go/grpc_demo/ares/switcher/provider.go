package switcher

import ares "ares/pkg/io"

type Provider struct {
	sessions   ares.ISessions
	msgCreator ares.IMsgCreator
}

func (p *Provider) Sessions() ares.ISessions {
	return p.sessions
}

func (p *Provider) MsgCreator() ares.IMsgCreator {
	return p.msgCreator
}
