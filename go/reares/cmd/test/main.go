package main

import (
	"bufio"
	"fmt"
	"reares/internal/io"
	"reares/logic/echo"
	"reares/proto"
)

func main() {
	io.MsgCreator[1] = func() io.Msg { return proto.NewEcho() }
	io.MsgProcessor[1] = func(msg io.Msg) error { return echo.ProcessEcho(msg.(*proto.Echo)) }

	e := proto.NewEcho()
	e.Msg = "test"
	buffer := io.GetBuffer()
	var msg io.Msg = e
	err := io.EncodeMsg(nil, buffer, msg)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(buffer)
	data, err := io.DecodeFrame(reader)
	if err != nil {
		panic(err)
	}
	decodeMsg, err := io.DecodeMsg(nil, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(decodeMsg)
	//err = decodeMsg.Process()
	//if err != nil {
	//	panic(err)
	//}
	decodeMsg.Dispatch()
}
