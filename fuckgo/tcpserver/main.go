package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go func(conn net.Conn) {
			defer conn.Close()
			buf := make([]byte, 1024)
			_, err := conn.Read(buf)
			if err != nil {
				log.Println("Error reading:", err.Error())
			}
			conn.Write([]byte("Message received."))
		}(conn)
	}
}
