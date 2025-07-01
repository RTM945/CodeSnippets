package switcher

import (
	ares "ares/pkg/io"
	"ares/switcher/msg"
)

var linkerMsgRegistry = map[uint32]ares.MsgCreatorFunc{
	4: msg.PingCreator,
}
