package main

import (
	"bufio"
	"fmt"
	"reares/internal/io"
	"reares/proto"
)

func main() {
	io.MsgCreator[1] = func() io.Msg { return proto.NewEcho() }
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
	decodeMsg, err := io.DecodeMsg(nil, data)
	fmt.Println(decodeMsg)
}
