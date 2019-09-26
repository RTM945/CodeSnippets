package simple

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	commands := []string{"w", "t", "f", "END"}
	for _, command := range commands {
		go func(command string) {
			conn, err := net.Dial("tcp", ":10001")
			if err != nil {
				t.Error("could not connect to server: ", err)
			}
			defer conn.Close()
			if command == "END" {
				time.Sleep(time.Second)
			}
			log.Printf("Send %s \n", command)
			conn.Write([]byte(command))
		}(command)
	}
	server := &JustHandleConnectionServer{}
	server.start(10001)
}
