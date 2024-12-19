package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reares/pkg/rc4"
	"reares/pkg/rsa"
)

func main() {
	//logic.Init()
	//connector := io.NewConnector()
	//connector.Connect("127.0.0.1:18290")
	//echo := echo.NewCEcho()
	//echo.Msg = "test"
	//connector.Send(echo)
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	//<-c
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	println("connected to server", conn.RemoteAddr().String())
	session := &Session{}

	byteBuf := bytes.NewBuffer(make([]byte, 0, 1024))

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				println(err)
			}
			break
		}
		if n == 0 {
			continue
		}
		buffer = buffer[:n]
		if session.securityDecoder != nil {
			// 流加密解密
			session.securityDecoder.DoUpdate(buffer)
			println("decode success:", conn.RemoteAddr().String())
		}
		byteBuf.Write(buffer[:n])

		if byteBuf.Len() < 4 {
			continue
		}
		frameLength := binary.BigEndian.Uint32(byteBuf.Bytes()[:4])
		if byteBuf.Len() < int(frameLength)+4 {
			continue
		}

		frame := byteBuf.Next(int(frameLength) + 8)
		typeId := binary.BigEndian.Uint32(frame[4:8])
		switch typeId {
		case 1:
			// RSA KEY EXCHANGE
			session.serverPublicKey = frame[8:]
			session.rsa = rsa.GetInstance()
			encoded, err := session.rsa.GetPublicKeyEncoded()
			if err != nil {
				println(err)
				return
			}
			println("recv rsa key from:", conn.RemoteAddr().String())
			send := make([]byte, 8)
			// length
			binary.BigEndian.PutUint32(send[:4], uint32(len(encoded)))
			// typeId
			binary.BigEndian.PutUint32(send[4:8], 1)
			send = append(send, encoded...)
			_, _ = conn.Write(send)
			println("send rsa key to:", conn.RemoteAddr().String())
		case 2:
			// KEY EXCHANGE
			encrypt := frame[8:]
			serverKey, err := rsa.Decrypt(session.rsa.GetPrivateKey(), encrypt)
			if err != nil {
				println(err)
				return
			}
			fmt.Println("serverKey:", serverKey)
			session.securityDecoder = rc4.NewRC4(serverKey)
			println("recv key from:", conn.RemoteAddr().String())
			key := randomKey(32)
			encodedKey := make([]byte, base64.StdEncoding.EncodedLen(len(key)))
			base64.StdEncoding.Encode(encodedKey, key)
			session.securityEncoder = rc4.NewRC4(encodedKey)
			fmt.Println("clientKey:", key)
			encrypt, err = rsa.Encrypt(session.serverPublicKey, encodedKey)
			if err != nil {
				println(err)
				return
			}

			send := make([]byte, 8)
			// length
			binary.BigEndian.PutUint32(send[:4], uint32(len(encrypt)))
			// typeId
			binary.BigEndian.PutUint32(send[4:8], 2)
			send = append(send, encrypt...)
			session.securityDecoder.DoUpdate(send)
			println("encode success:", conn.RemoteAddr().String())
			_, _ = conn.Write(send)
			println("send key to:", conn.RemoteAddr().String())
		case 3:
			println(string(frame[8:]))
		}
	}
}

type Session struct {
	rsa             *rsa.Key
	securityDecoder *rc4.RC4
	securityEncoder *rc4.RC4
	serverPublicKey []byte
}

func randomKey(size int) []byte {
	res := make([]byte, size)
	rand.Read(res)
	return res
}
