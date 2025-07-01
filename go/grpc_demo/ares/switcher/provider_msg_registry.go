package switcher

import (
	ares "ares/pkg/io"
	"ares/switcher/msg"
)

var providerMsgRegistry = map[uint32]ares.MsgCreatorFunc{
	51: msg.DispatchCreator,
}
