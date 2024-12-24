package shard

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/utils"
	"github.com/go-netty/go-netty/utils/pool/pbuffer"
	"math"
	"reares/internal/io"
)

var LengthFieldBasedFrameDecoder = lengthFieldBasedFrameDecoder{
	frame.LengthFieldCodec(binary.BigEndian, math.MaxInt, 0, 4, 0, 4),
}

type lengthFieldBasedFrameDecoder struct {
	codec codec.Codec
}

func (fc lengthFieldBasedFrameDecoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fc.codec.HandleRead(ctx, message)
}

func (fc lengthFieldBasedFrameDecoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	// 覆盖LengthFieldCodec的实现
	// 这里什么都不做直接交给下一个处理器
	ctx.HandleWrite(message)
}

type MsgEncoder struct{}

func (MsgEncoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	msg := message.(io.Msg)
	fmt.Println("encode:", msg.GetHeader().TypeId)
	buffer := pbuffer.Get(4096)
	buffer.Reset()
	var err error
	err = binary.Write(buffer, binary.BigEndian, uint32(0))
	if err != nil {
		panic(err)
	}

	header := msg.GetHeader()
	err = header.Encode(buffer)
	if err != nil {
		panic(err)
	}
	err = msg.Encode(buffer)
	if err != nil {
		panic(err)
	}
	totalLen := buffer.Len() - 4
	data := buffer.Bytes()
	binary.BigEndian.PutUint32(data[:4], uint32(totalLen))
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
	fmt.Println("decode:", msg.GetHeader().TypeId)
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
