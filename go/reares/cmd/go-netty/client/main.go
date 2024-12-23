package main

import (
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"math"
	shard "reares/cmd/go-netty"
	"reares/logic"
	"reares/proto/echo"
	"time"
)

func main() {
	logic.Init()
	var childInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.LengthFieldCodec(binary.BigEndian, math.MaxInt, 0, 4, 0, 4)).
			AddLast(shard.MsgEncoder{}).
			AddLast(shard.MsgDecoder{}).
			AddLast(shard.LogicHandler{})
	}
	client, err := netty.NewBootstrap(
		netty.WithClientInitializer(childInitializer),
	).Connect("127.0.0.1:9527")
	if err != nil {
		return
	}
	echo := echo.NewCEcho()
	echo.Msg = "test"
	client.Write(echo)
	time.Sleep(time.Second * 3)
}
