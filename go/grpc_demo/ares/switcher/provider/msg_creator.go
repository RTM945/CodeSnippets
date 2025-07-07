package provider

import (
	ares "ares/pkg/io"
	"ares/switcher/msg"
)

var MsgCreator = map[uint32]ares.MsgCreatorFunc{
	73: msg.SendToClientCreator,
}
