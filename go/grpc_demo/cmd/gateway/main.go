package main

import (
	"grpc_demo/service/0ares/gateway"
)

func main() {
	go func() {
		gateway.StartServer()
	}()
	select {}
}
