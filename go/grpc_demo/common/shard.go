package common

import (
	hellomsg "grpc_demo/common/gen"
)

var MsgCreator = map[uint32]func() Msg{
	1: func() Msg { return hellomsg.NewHelloRequest() },
}
