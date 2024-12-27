package client_switch

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	shard "reares/cmd/go-netty"
	"reares/protobuf"
)

type KeyExchange struct {
	*shard.MsgBase
	Processor
	Key []byte
}

func InitKeyExchange(msgProcessor Processor) *KeyExchange {
	header := &shard.MsgHeader{}
	header.TypeId = 2
	return &KeyExchange{
		MsgBase:   shard.NewMsgBase(header),
		Processor: msgProcessor,
	}
}

func NewKeyExchange() *KeyExchange {
	header := &shard.MsgHeader{}
	header.TypeId = 2
	return &KeyExchange{
		MsgBase: shard.NewMsgBase(header),
	}
}

func (msg *KeyExchange) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.KeyExchange{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	msg.Key = tmp.Key
	return nil
}

func (msg *KeyExchange) Encode(dst *bytes.Buffer) error {
	tmp := &protobuf.RSAKeyExchange{
		Key: msg.Key,
	}
	data, err := proto.Marshal(tmp)
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (msg *KeyExchange) Dispatch() {
	msg.Process()
}

func (msg *KeyExchange) Process() error {
	return msg.ProcessKeyExchange(msg)
}
