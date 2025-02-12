package client_switch

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"reares/internal/io"
	"reares/protobuf"
)

type RSAKeyExchange struct {
	*io.MsgBase
	Processor
	Key []byte
}

func InitRSAKeyExchange(msgProcessor Processor) *RSAKeyExchange {
	header := &io.MsgHeader{}
	header.TypeId = 1
	return &RSAKeyExchange{
		MsgBase:   io.NewMsgBase(header),
		Processor: msgProcessor,
	}
}

func NewRSAKeyExchange() *RSAKeyExchange {
	header := &io.MsgHeader{}
	header.TypeId = 1
	return &RSAKeyExchange{
		MsgBase: io.NewMsgBase(header),
	}
}

func (msg *RSAKeyExchange) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.RSAKeyExchange{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	msg.Key = tmp.Key
	return nil
}

func (msg *RSAKeyExchange) Encode(dst *bytes.Buffer) error {
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

func (msg *RSAKeyExchange) Dispatch() {
	msg.Process()
}

func (msg *RSAKeyExchange) Process() error {
	return msg.ProcessRSAKeyExchange(msg)
}
