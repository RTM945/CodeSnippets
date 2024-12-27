package logic

import (
	shard "reares/cmd/go-netty"
	echoLogic "reares/cmd/go-netty/logic/echo"
	echoProto "reares/cmd/go-netty/proto/echo"
)

func Init() {
	echoProcessor := echoLogic.NewMsgProcessor()
	shard.MsgCreator[100] = func() shard.Msg { return echoProto.InitCEcho(echoProcessor) }
	shard.MsgCreator[101] = func() shard.Msg { return echoProto.InitSEcho(echoProcessor) }
}