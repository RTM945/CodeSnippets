package switcher

import (
	ares "ares/pkg/io"
	"ares/switcher/msg"
)

type ProviderMsgCreator struct {
	registry map[int32]ares.MsgCreatorFunc
}

func NewProviderMsgCreator() *ProviderMsgCreator {
	msgCreator := &ProviderMsgCreator{
		registry: make(map[int32]ares.MsgCreatorFunc),
	}
	return msgCreator
}

func (c *ProviderMsgCreator) Register(typeId int32, f ares.MsgCreatorFunc) {
	c.registry[typeId] = f
}

func (c *ProviderMsgCreator) Create(session ares.ISession, pvId, typeId int32, payload []byte) (ares.IMsg, error) {
	creator := c.registry[typeId]
	if creator != nil {
		return creator(session, pvId, typeId, payload)
	} else {
		providerSession := session.(*ProviderSession)
		toSession, ok := provider.Sessions().GetSession(pvId).(*ProviderSession)
		if toSession == nil || !ok {
			LOGGER.Errorf("Providee to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", pvId, providerSession, typeId)
			return nil, nil
		}
		pDispatch := msg.NewPDispatch()
		pDispatch.TypedPB().PvId = providerSession.GetPvId()
	}
	return nil, nil
}
