package linker

import (
	ares "ares/pkg/io"
	"ares/switcher/msg"
)

func initMsg(creator ares.IMsgCreator) {
	creator.Register(4, msg.PingCreator)
}
