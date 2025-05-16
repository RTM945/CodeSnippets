package common

import (
	"grpc_demo/common/gen"
)

var MsgCreator = map[uint32]func() Msg{
	1: func() Msg { return hellomsg.NewHelloRequest() },
}

type TestProcessor struct {
}

func (test *TestProcessor) Process(msg Msg) error {
	return nil
}

var MsgProcessor = map[uint32]Processor{
	1: &TestProcessor{},
}
