package network

import (
	"errors"
	"log"
	"net"
)

type Handler func(conn net.Conn, packet *Packet)

// StartTCPServer 启动TCP服务器
func StartTCPServer(port string) {
	// 同时支持 TCP 和 TCP6
	listener, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TCP Server is listening at ", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept connect error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
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
			packet, err := Unpack(messageBuffer)
			if err != nil {
				if errors.Is(err, PacketNotCompleteErr) {
					break
				}
				log.Println("Unpack error:", err)
				return
			}

			// 处理消息
			//handleMessage(conn, packet)

			// 移除已处理的消息
			messageBuffer = messageBuffer[HeaderSize+packet.Length:]
		}
	}
}
