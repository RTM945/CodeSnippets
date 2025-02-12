package main

import (
	"github.com/go-netty/go-netty"
	shard "reares/cmd/go-netty"
	"reares/cmd/go-netty/logic"
)

const Service = "Linker"

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
	err := netty.NewBootstrap(netty.WithChildInitializer(childInitializer)).
		Listen(":9527").Sync()
	if err != nil {
		return
	}
}
