package switcher

import ares "ares/pkg/io"

var linkerMsgRegistry = map[uint32]ares.MsgCreatorFunc{}
