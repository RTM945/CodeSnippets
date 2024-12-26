package main

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/go-netty/go-netty"
	shard "reares/cmd/go-netty"
	"reares/cmd/go-netty/proto/client-switch"
	"reares/cmd/go-netty/proto/echo"
	"reares/pkg/rc4"
	"reares/pkg/rsa"
)

type NodeFactory struct {
}

func (factory NodeFactory) CreateSession(channel netty.Channel) shard.Session {
	session := Session{
		channel: channel,
	}
	return &session
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

}

type Session struct {
	rsa     *rsa.Key
	channel netty.Channel
}

func (s Session) Send(msg shard.Msg) error {
	return s.channel.Write(msg)
}

func (s Session) GetSid() int32 {
	return 0
}

func (s Session) OnClose() {}

func (s Session) SendKeyExchange(clientPublicKey []byte) error {
	key := randomKey(32)
	encodedKey := make([]byte, base64.StdEncoding.EncodedLen(len(key)))
	base64.StdEncoding.Encode(encodedKey, key)
	securityDecoder := shard.SecurityDecoder{
		RC4: rc4.NewRC4(encodedKey),
	}
	s.channel.Pipeline().AddFirst(securityDecoder)
	encrypt, err := rsa.Encrypt(clientPublicKey, encodedKey)
	if err != nil {
		return err
	}
	exchange := client_switch.NewKeyExchange()
	exchange.Key = encrypt
	return s.Send(exchange)
}

func randomKey(size int) []byte {
	res := make([]byte, size)
	rand.Read(res)
	return res
}

func (s Session) SetClientKey(clientKey []byte) error {
	serverKey, err := rsa.Decrypt(s.rsa.GetPrivateKey(), clientKey)
	if err != nil {
		return err
	}
	securityEncoder := shard.SecurityEncoder{
		RC4: rc4.NewRC4(serverKey),
	}
	s.channel.Pipeline().AddFirst(securityEncoder)
	secho := echo.NewSEcho()
	secho.Msg = "【公式】勝利の女神：NIKKE\n@NIKKE_japan\n·\n21h\n【NIKKEモーション紹介】\nラピ：レッドフード(CV：#石川由依)\n\n◆使用武器\nマシンガン「セブンスドワーフゼロ」\n\n◆バーストスキル\n部隊の構成次第でバーストⅠ、またはバーストⅢとして活用可能。\n風圧コードと電撃コードの敵に有利なコードが適用される。\n\n※戦闘動画はテスト環境で撮った内容であり、実際の内容はゲーム内をご参照ください。"
	return s.Send(secho)
}
