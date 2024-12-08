package io

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

type Acceptor struct {
	host, port string
	listener   net.Listener
}

func (acceptor *Acceptor) Start() error {
	listener, err := net.Listen("tcp", net.JoinHostPort(acceptor.host, acceptor.port))
	if err != nil {
		return err
	}
	acceptor.listener = listener

	go acceptor.accept()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	return nil
}

func (acceptor *Acceptor) accept() {
	for {
		conn, err := acceptor.listener.Accept()
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

	reader := bufio.NewReader(conn)

	for {
		buffer, err := DecodeFrame(reader)
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed connection")
				RemoveSession(session)
				break
			} else {
				log.Println("Error decoding frame:", err)
				RemoveSession(session)
				break
			}
		}
		msg, err := DecodeMsg(session, buffer)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue
		}

		msg.Dispatch()

		log.Println(msg)

	}

	RemoveSession(session)
	_ = conn.Close()
}

func DecodeFrame(reader *bufio.Reader) ([]byte, error) {
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

func DecodeMsg(session *Session, data []byte) (Msg, error) {
	buffer := GetBuffer()
	buffer.Write(data)
	header := &MsgHeader{}
	err := header.Decode(buffer)
	if err != nil {
		return nil, err
	}
	msg, err := CreateMsg(header, session, buffer)
	PutBuffer(buffer)
	return msg, err
}

func EncodeMsg(session *Session, buffer *bytes.Buffer, msg Msg) error {
	err := binary.Write(buffer, binary.BigEndian, uint32(0)) // 占位
	if err != nil {
		return err
	}
	header := msg.GetHeader()
	err = header.Encode(buffer)
	if err != nil {
		return err
	}
	err = msg.Encode(buffer)
	if err != nil {
		return err
	}
	totalLen := buffer.Len() - lengthFieldLen
	data := buffer.Bytes()

	binary.BigEndian.PutUint32(data[:4], uint32(totalLen))

	// todo session log
	return nil
}
