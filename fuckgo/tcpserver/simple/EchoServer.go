package simple

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type EchoServer struct {
	l net.Listener
}

func (server *EchoServer) Start(port string, c chan int) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	server.l = l
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

func (server *EchoServer) Close() {
	server.l.Close()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("handle %s \n", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR", err)
			break
		} else {
			msg := string(line)
			fmt.Printf("---> %s: %s\n", conn.RemoteAddr().String(), msg)
			fmt.Fprintln(conn, msg)
			if msg == "END" {
				fmt.Println("connection end")
				break
			}
		}
	}
}
