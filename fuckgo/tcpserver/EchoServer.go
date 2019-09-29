package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":10001")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}
		go Echo(conn)
	}
}

func Echo(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		req, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println(err)
			break
		}
		text := string(req)
		fmt.Printf("---> %s: %s\n", conn.RemoteAddr().String(), text)
		if text == "END\n" {
			conn.Write([]byte("bye~\n"))
			break
		}
		conn.Write(req)
	}
}
