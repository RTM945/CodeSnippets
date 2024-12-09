package main

import (
	"os"
	"os/signal"
	"reares/internal/io"
	"reares/logic"
	proto "reares/proto/echo"
	"syscall"
)

func main() {
	logic.Init()
	connector := io.NewConnector()
	connector.Connect("127.0.0.1:18290")
	echo := proto.NewCEcho()
	echo.Msg = "test"
	connector.Send(echo)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
