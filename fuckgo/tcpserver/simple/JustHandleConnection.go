package simple

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type JustHandleConnectionServer struct {
	stop bool
}

func (server *JustHandleConnectionServer) start(port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for !server.stop {
		conn, err := l.Accept()
		if err != nil {
			//log.Fatal("Error accepting: ", err.Error())
			return
		}
		log.Println("conn Accepted.")
		go func(conn net.Conn, server *JustHandleConnectionServer) {
			defer conn.Close()
			buf, err := ioutil.ReadAll(conn)
			if err != nil {
				return
			}
			msg := string(buf[:])
			log.Println(msg)
			if msg == "END" {
				server.stop = true
				l.Close()
			}
		}(conn, server)
	}
}
