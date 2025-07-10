package switcher

import ares "ares/pkg/io"

type ProviderSessions struct {
	*ares.Sessions
	node ares.INode
}

func NewProviderSessions(node ares.INode) *ProviderSessions {
	return &ProviderSessions{
		Sessions: ares.NewSessions(),
		node:     node,
	}
}
