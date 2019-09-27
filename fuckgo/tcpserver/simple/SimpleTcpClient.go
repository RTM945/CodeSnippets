package simple

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type TcpClient struct {
	conn net.Conn
}

func (client *TcpClient) start(port string, c chan int) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client.conn = conn
	defer client.Close()
	fmt.Println("start client")
	reader := bufio.NewReader(conn)
	c <- 1
	for {
		response, err := reader.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Printf("<--- %s: %s\n", conn.RemoteAddr().String(), string(response))
		case io.EOF:
			fmt.Println("remote closed", err)
			break
		default:
			fmt.Println("ERROR", err)
			break
		}
		if string(response) == "END" {
			break
		}
	}
}

func (client *TcpClient) Send(msg string) {
	if client.conn == nil {
		fmt.Errorf("connection closed")
	} else {
		client.conn.Write([]byte(msg))
		fmt.Printf("send %s \n", msg)
	}
}

func (client *TcpClient) Close() {
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}
}
