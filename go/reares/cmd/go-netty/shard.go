package shard

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/go-netty/go-netty/utils/pool/pbuffer"
	"reares/internal/io"
)

type MsgEncoder struct{}

func (MsgEncoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	msg := message.(io.Msg)
	buffer := pbuffer.Get(4096)
	buffer.Reset()
	//err := binary.Write(buffer, binary.BigEndian, uint32(0))
	//if err != nil {
	//	panic(err)
	//}
	var err error
	header := msg.GetHeader()
	err = header.Encode(buffer)
	if err != nil {
		panic(err)
	}
	err = msg.Encode(buffer)
	if err != nil {
		panic(err)
	}
	//totalLen := buffer.Len() - 4
	//data := buffer.Bytes()
	//binary.BigEndian.PutUint32(data[:4], uint32(totalLen))
	if msg.GetHeader().TypeId == 2 {
		fmt.Println()
	}
	ctx.HandleWrite(buffer)
	pbuffer.Put(buffer)
}

type MsgDecoder struct{}

func (MsgDecoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	session := ctx.Channel().Attachment().(io.Session)
	data := utils.MustToBytes(message)
	buffer := pbuffer.Get(len(data))
	buffer.Reset()
	buffer.Write(data)
	header := &io.MsgHeader{}
	err := header.Decode(buffer)
	if err != nil {
		panic(err)
	}
	msg, err := io.CreateMsg(header, session, buffer)
	if err != nil {
		panic(err)
	}
	ctx.HandleRead(msg)
}

type LogicHandler struct{}

func (LogicHandler) HandleActive(ctx netty.ActiveContext) {
	ctx.Channel().SetAttachment(session{channel: ctx.Channel()})
	ctx.HandleActive()
}

func (LogicHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	msg := message.(io.Msg)
	msg.Dispatch()
	ctx.HandleRead(message)
}

func (LogicHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	session := ctx.Channel().Attachment().(io.Session)
	session.OnClose()
	ctx.HandleInactive(ex)
}

func (LogicHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	fmt.Println("HandleException:", ex)
}

type session struct {
	channel netty.Channel
}

func (s session) Send(msg io.Msg) error {
	return s.channel.Write(msg)
}

func (s session) GetSid() int32 {
	return 0
}

func (s session) OnClose() {}
