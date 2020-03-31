package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":10001")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	respr := bufio.NewReader(conn)
	for {
		input, rerr := r.ReadBytes(byte('\n'))
		if rerr != nil {
			fmt.Println(rerr)
			break
		}
		conn.Write(input)

		resp, err := respr.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println(err)
			break
		}
		text := string(resp)
		fmt.Printf("---> %s: %s\n", conn.RemoteAddr().String(), text)
		if text == "bye~\n" {
			break
		}
	}
}
