package main

import (
	"encoding/base64"
	"github.com/go-netty/go-netty"
	shard "reares/cmd/go-netty"
	client_switch "reares/cmd/go-netty/proto/client-switch"
	"reares/pkg/rc4"
	"reares/pkg/rsa"
)

type NodeFactory struct {
}

func (NodeFactory) CreateSession(channel netty.Channel) shard.Session {
	session := Session{
		StateSession: shard.NewStateSession(channel),
	}
	return &session
}

func (NodeFactory) OnAddSession(session shard.Session) {}

func (factory NodeFactory) OnRemoveSession(session shard.Session) {

}

type Session struct {
	*shard.StateSession
	rsa             *rsa.Key
	serverPublicKey []byte
}

func (s *Session) SetServerPublicKey(publicKey []byte) {
	s.serverPublicKey = publicKey
}

func (s *Session) SendRSAKeyExchange(serverPublicKey []byte) error {
	s.SetServerPublicKey(serverPublicKey)
	s.rsa = rsa.GetInstance()
	encoded, err := s.rsa.GetPublicKeyEncoded()
	if err != nil {
		return err
	}
	rsaKeyExchange := client_switch.NewRSAKeyExchange()
	rsaKeyExchange.Key = encoded
	return s.Send(rsaKeyExchange)
}

func (s *Session) SetServerKey(serverKey []byte) error {
	serverKey, err := rsa.Decrypt(s.rsa.GetPrivateKey(), serverKey)
	if err != nil {
		return err
	}
	securityEncoder := shard.SecurityEncoder{
		RC4: rc4.NewRC4(serverKey),
	}
	s.GetChannel().Pipeline().AddFirst(securityEncoder)
	return s.SendKeyExchange()
}

func (s *Session) SendKeyExchange() error {
	key := shard.RandomKey(32)
	encodedKey := make([]byte, base64.StdEncoding.EncodedLen(len(key)))
	base64.StdEncoding.Encode(encodedKey, key)
	securityDecoder := shard.SecurityDecoder{
		RC4: rc4.NewRC4(encodedKey),
	}
	s.GetChannel().Pipeline().AddFirst(securityDecoder)
	encrypt, err := rsa.Encrypt(s.serverPublicKey, encodedKey)
	if err != nil {
		return err
	}
	exchange := client_switch.NewKeyExchange()
	exchange.Key = encrypt
	return s.Send(exchange)
}
