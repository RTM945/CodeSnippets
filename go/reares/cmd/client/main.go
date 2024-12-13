package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"net"
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

	// Generate client's ECDH key pair
	clientPriv, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating client key:", err)
		return
	}

	// Receive and deserialize server's public key
	serverPubBytes := make([]byte, 65) // 65 bytes for uncompressed P256 public key
	if _, err := conn.Read(serverPubBytes); err != nil {
		fmt.Println("Error receiving server public key:", err)
		return
	}
	pub, err := x509.ParsePKIXPublicKey(serverPubBytes)
	if err != nil {
		fmt.Println("Error deserializing server public key:", err)
		return
	}
	serverPub, ok := pub.(*ecdh.PublicKey)
	if !ok {
		fmt.Println("key is not an ECDH public key")
		return
	}

	// Serialize and send client's public key
	key, err := x509.MarshalPKIXPublicKey(clientPriv.PublicKey())
	if err != nil {
		fmt.Println("Error marshalling client key:", err)
		return
	}
	if _, err := conn.Write(key); err != nil {
		fmt.Println("Error sending client public key:", err)
		return
	}

	// Derive shared key
	sharedSecret, err := clientPriv.ECDH(serverPub)
	if err != nil {
		fmt.Println("Error deriving shared key:", err)
		return
	}
	sharedKey := sha256.Sum256(sharedSecret)
	fmt.Println("Shared key derived on client:", sharedKey)

	// Read acknowledgment
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Println("Message from server:", string(buf[:n]))
}
