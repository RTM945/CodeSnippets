package simple

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type EchoServer struct {
}

func (server EchoServer) start(port string, c chan int) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("start server")
	c <- 1
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("handle %s \n", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		response, err := reader.ReadBytes(byte('\n'))
		if err == nil {
			fmt.Printf("---> %s: %s\n", conn.RemoteAddr().String(), string(response))
		}
		if string(response) == "END" {
			conn.Write([]byte("bye"))
			break
		} else {
			conn.Write([]byte(string(string(response))))
		}
	}
}
