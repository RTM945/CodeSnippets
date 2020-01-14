package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	_, _ = nc.Subscribe("rtm", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	_ = nc.Publish("rtm", []byte("Hello World!"))
	select {}
}
