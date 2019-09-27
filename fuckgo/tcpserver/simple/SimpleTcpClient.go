package simple

import (
	"bufio"
	"fmt"
	"net"
)

type TcpClient struct {
	conn net.Conn
}

func (client *TcpClient) Start(port string, c chan int) {
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
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR", err)
			break
		} else {
			msg := string(line)
			fmt.Printf("<--- %s: %s\n", conn.RemoteAddr().String(), msg)
			if msg == "END" {
				fmt.Println("client end")
				return
			}
		}
	}
}

func (client *TcpClient) Send(msg string) {
	fmt.Fprintln(client.conn, msg)
}

func (client *TcpClient) Close() {
	client.conn.Close()
}
