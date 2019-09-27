package simple

import (
	"testing"
)

func TestServer(t *testing.T) {
	port := "0.0.0.0:10001"
	server := &EchoServer{}
	client := &TcpClient{}
	cs := make(chan int)
	cc := make(chan int)
	go server.Start(port, cs)
	defer server.Close()
	for {
		select {
		case <-cs:
			go client.Start(port, cc)
		case <-cc:
			commands := []string{"w", "t", "f", "END"}
			for _, command := range commands {
				client.Send(command)
			}
			return
		}
	}
}
