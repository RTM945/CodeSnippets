package logic

import (
	"reares/internal/io"
	"reares/logic/echo"
	proto "reares/proto/echo"
)

func Init() {
	echoProcessor := echo.NewMsgProcessor()
	io.MsgCreator[1] = func() io.Msg { return proto.NewCEchoWithProcessor(echoProcessor) }
	io.MsgCreator[2] = func() io.Msg { return proto.NewSEchoWithProcessor(echoProcessor) }
}
