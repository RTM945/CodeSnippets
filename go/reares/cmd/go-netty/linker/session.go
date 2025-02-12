package main

import (
	"encoding/base64"
	"github.com/go-netty/go-netty"
	shard "reares/cmd/go-netty"
	"reares/cmd/go-netty/proto/client-switch"
	"reares/pkg/rc4"
	"reares/pkg/rsa"
	"time"
)

type NodeFactory struct {
}

func (factory NodeFactory) CreateSession(channel netty.Channel) shard.Session {
	session := &Session{
		StateSession: shard.NewStateSession(channel),
		aliveTime:    time.Now().UnixMilli(),
	}
	sessions.Store(session.GetSid(), session)
	return session
}

func (factory NodeFactory) OnAddSession(session shard.Session) {
	linkerSession := session.(*Session)
	linkerSession.rsa = rsa.GetInstance()
	encoded, err := linkerSession.rsa.GetPublicKeyEncoded()
	if err != nil {
		panic(err)
	}
	rsaKeyExchange := client_switch.NewRSAKeyExchange()
	rsaKeyExchange.Key = encoded
	err = linkerSession.Send(rsaKeyExchange)
	if err != nil {
		panic(err)
	}
}

func (factory NodeFactory) OnRemoveSession(session shard.Session) {
	sessions.Delete(session.GetSid())
	// notify other severs
}

type Session struct {
	*shard.StateSession
	rsa       *rsa.Key
	aliveTime int64
}

func (s *Session) SendKeyExchange(clientPublicKey []byte) error {
	key := shard.RandomKey(32)
	encodedKey := make([]byte, base64.StdEncoding.EncodedLen(len(key)))
	base64.StdEncoding.Encode(encodedKey, key)
	securityDecoder := shard.SecurityDecoder{
		RC4: rc4.NewRC4(encodedKey),
	}
	s.GetChannel().Pipeline().AddFirst(securityDecoder)
	encrypt, err := rsa.Encrypt(clientPublicKey, encodedKey)
	if err != nil {
		return err
	}
	exchange := client_switch.NewKeyExchange()
	exchange.Key = encrypt
	return s.Send(exchange)
}

func (s *Session) ResetAlive() {
	s.aliveTime = time.Now().UnixMilli()
}

func (s *Session) SetClientKey(clientKey []byte) error {
	serverKey, err := rsa.Decrypt(s.rsa.GetPrivateKey(), clientKey)
	if err != nil {
		return err
	}
	securityEncoder := shard.SecurityEncoder{
		RC4: rc4.NewRC4(serverKey),
	}
	s.GetChannel().Pipeline().AddFirst(securityEncoder)
	// should send zone endpoints
	return nil
}

func (s *Session) alive() bool {
	return time.Now().UnixMilli()-s.aliveTime < 120000 // 120s
}
