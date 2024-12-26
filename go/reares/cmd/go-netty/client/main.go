package main

import (
	"github.com/go-netty/go-netty"
	shard "reares/cmd/go-netty"
	"reares/cmd/go-netty/logic"
	"time"
)

const Service = "Client"

func main() {
	Init()
	logic.Init()
	var childInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(shard.MsgEncoder{}).
			AddLast(shard.LengthFieldBasedFrameDecoder).
			AddLast(shard.MsgDecoder{}).
			AddLast(shard.LogicHandler{
				NodeFactory: NodeFactory{},
			})
	}
	_, err := netty.NewBootstrap(
		netty.WithClientInitializer(childInitializer),
	).Connect("127.0.0.1:9527")
	if err != nil {
		return
	}
	//echo := echo.NewCEcho()
	//echo.Msg = "test"
	//client.Write(echo)
	time.Sleep(time.Second * 100)
}
