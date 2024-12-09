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

func NewAcceptor(host, port string) *Acceptor {
	return &Acceptor{host: host, port: port}
}

func (acceptor *Acceptor) Start() error {
	listener, err := net.Listen("tcp", net.JoinHostPort(acceptor.host, acceptor.port))
	if err != nil {
		return err
	}
	acceptor.listener = listener

	go acceptor.accept()
	log.Println("Listening on " + acceptor.listener.Addr().String())
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
		tcpConn := conn.(*net.TCPConn)
		onNewTcpConnection(tcpConn)
		session := NewSession(tcpConn)
		reader := bufio.NewReader(tcpConn)
		go func() {
			onRead(session, reader)
			RemoveSession(session)
			_ = tcpConn.Close()
		}()
	}
}

type Connector struct {
	*Session
}

func NewConnector() *Connector {
	return &Connector{}
}

func (connector *Connector) Connect(target string) {
	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatal(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	onNewTcpConnection(tcpConn)
	session := NewSession(tcpConn)
	reader := bufio.NewReader(tcpConn)
	connector.Session = session
	go func() {
		onRead(session, reader)
		RemoveSession(session)
		_ = tcpConn.Close()
	}()
}

func (connector *Connector) Send(msg Msg) error {
	return connector.Session.Send(msg)
}

func onNewTcpConnection(tcpConn *net.TCPConn) {
	err := tcpConn.SetReadBuffer(655360)
	if err != nil {
		log.Printf("SetReadBuffer error: %v\n", err)
	}
	err = tcpConn.SetWriteBuffer(655360)
	if err != nil {
		log.Printf("SetWriteBuffer error: %v\n", err)
	}
	err = tcpConn.SetNoDelay(true)
	if err != nil {
		log.Printf("SetNoDelay error: %v\n", err)
	}
	err = tcpConn.SetKeepAlive(true)
	if err != nil {
		log.Printf("SetKeepAlive error: %v\n", err)
	}
}

func onRead(session *Session, reader *bufio.Reader) {
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

	}
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
	log.Println("Recv msg:", msg)
	return msg, err
}

func EncodeMsg(buffer *bytes.Buffer, msg Msg) error {
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
	log.Println("Send msg:", msg)
	return nil
}
