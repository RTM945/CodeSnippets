package network

import (
	"errors"
	"log"
	"net"
)

type Server struct {
	addr     string
	listener net.Listener
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.listener = listener

	go s.accept()
	return nil
}

func (s *Server) accept() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
			continue
		}

		go func() {
			handleConnection(conn)
		}()
	}
}

func handleConnection(conn net.Conn) {
	session := NewSession(conn)
	defer func() {
		RemoveSession(session)
		conn.Close()
	}()
	buffer := make([]byte, 1024)
	messageBuffer := make([]byte, 0)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read connection error:", err)
			return
		}

		messageBuffer = append(messageBuffer, buffer[:n]...)

		for len(messageBuffer) >= HeaderSize {
			length, protoId, protoData, err := Unpack(messageBuffer)
			if err != nil {
				if errors.Is(err, PacketNotCompleteErr) {
					break
				}
				log.Println("Unpack error:", err)
				messageBuffer = nil
				break
			}

			msg, err := CreateMsg(protoId, protoData)
			if err != nil {
				log.Printf("Create msg error protoId=%d err=%v:", protoId, err)
				messageBuffer = nil
				break
			}

			log.Println(msg)

			// 移除已处理的消息
			messageBuffer = messageBuffer[HeaderSize+length:]
		}
	}
}
