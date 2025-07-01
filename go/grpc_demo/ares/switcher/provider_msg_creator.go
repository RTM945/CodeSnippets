package switcher

import (
	ares "ares/pkg/io"
)

type ProviderMsgCreator struct {
	registry map[uint32]ares.MsgCreatorFunc
}

func NewProviderMsgCreator() *ProviderMsgCreator {
	msgCreator := &ProviderMsgCreator{
		registry: providerMsgRegistry,
	}
	return msgCreator
}

func (c *ProviderMsgCreator) Register(id uint32, f ares.MsgCreatorFunc) {
	c.registry[id] = f
}

func (c *ProviderMsgCreator) Create(session ares.ISession, pvId, typeId uint32, payload []byte) (ares.IMsg, error) {
	creator, ok := c.registry[typeId]
	if ok {
		return creator(session, pvId, typeId, payload)
	} else {
		providerSession := session.(*ProviderSession)
		toSession, ok := provider.Sessions().GetSession(pvId).(*ProviderSession)
		if toSession == nil || !ok {
			LOGGER.Errorf("Providee to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", pvId, providerSession, typeId)
			return nil, nil
		}

	}
	return nil, nil
}
