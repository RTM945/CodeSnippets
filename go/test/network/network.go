package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	maxFrameLength = 1<<31 - 1
	lengthFieldLen = 4
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
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
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

	reader := bufio.NewReader(conn)

	for {
		buffer, err := decodeFrame(reader)
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed connection")
				return
			} else {
				log.Println("Error decoding frame:", err)
				return
			}
		}
		msg, err := decodeMsg(session, buffer)
		if err != nil {
			log.Println("Error decoding message:", err)
			break
		}

		log.Println(msg)

	}
}

func decodeFrame(reader *bufio.Reader) ([]byte, error) {
	lengthField := make([]byte, lengthFieldLen)
	if _, err := io.ReadFull(reader, lengthField); err != nil {
		return nil, err
	}

	messageLength := int32(binary.BigEndian.Uint32(lengthField))

	if messageLength > maxFrameLength {
		return nil, fmt.Errorf("frame too large: %d", messageLength)
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("invalid frame length: %d", messageLength)
	}

	message := make([]byte, messageLength)
	if _, err := io.ReadFull(reader, message); err != nil {
		return nil, err
	}

	return message, nil
}

func decodeMsg(session *Session, data []byte) (Msg, error) {
	buffer := bytes.NewBuffer(data)
	header := &MsgHeader{}
	err := header.Decode(buffer)
	if err != nil {
		return nil, err
	}
	return CreateMsg(header, session, buffer)
}
