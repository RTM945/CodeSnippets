package main

import (
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"math"
	shard "reares/cmd/go-netty"
	"reares/logic"
)

type LoggerHandler struct{}

func main() {
	logic.Init()
	var childInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.LengthFieldCodec(binary.BigEndian, math.MaxInt, 0, 4, 0, 4)).
			AddLast(shard.MsgEncoder{}).
			AddLast(shard.MsgDecoder{}).
			AddLast(shard.LogicHandler{})
	}
	err := netty.NewBootstrap(netty.WithChildInitializer(childInitializer)).
		Listen(":9527").Sync()
	if err != nil {
		return
	}
}
