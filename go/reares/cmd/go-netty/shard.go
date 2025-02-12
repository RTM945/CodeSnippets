package shard

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/utils"
	"github.com/tidwall/transform"
	"log"
	"math"
	"reares/pkg/rc4"
	"sync/atomic"
)

type Session interface {
	Send(msg IMsg) error
	GetSid() int32
	OnClose()
	GetChannel() netty.Channel
}

var GenSessionId int32

type session struct {
	sid     int32
	channel netty.Channel
}

func newSession(channel netty.Channel) *session {
	return &session{
		channel: channel,
		sid:     atomic.AddInt32(&GenSessionId, 1),
	}
}

func (s *session) Send(msg IMsg) error {
	return s.channel.Write(msg)
}

func (s *session) GetSid() int32 {
	return s.sid
}

func (s *session) OnClose() {
	s.channel.Close(nil)
}

func (s *session) GetChannel() netty.Channel {
	return s.channel
}

type StateSession struct {
	*session
	state int32
}

func NewStateSession(channel netty.Channel) *StateSession {
	return &StateSession{
		session: newSession(channel),
	}
}

func (s *StateSession) AddState(state int32) {
	s.state |= state
}

func (s *StateSession) RemoveState(state int32) {
	s.state &= ^state
}

func (s *StateSession) GetState() int32 {
	return s.state
}

func (s *StateSession) CheckState(state int32) bool {
	return s.state&state == state
}

var LengthFieldBasedFrameDecoder = lengthFieldBasedFrameDecoder{
	codec: frame.LengthFieldCodec(binary.BigEndian, math.MaxInt, 0, 4, 0, 4),
}

type lengthFieldBasedFrameDecoder struct {
	codec codec.Codec
}

func (l lengthFieldBasedFrameDecoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	l.codec.HandleRead(ctx, message)
}

func (lengthFieldBasedFrameDecoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	ctx.HandleWrite(message)
}

type MsgEncoder struct {
}

func (e MsgEncoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	msg := message.(IMsg)
	buffer := new(bytes.Buffer)
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
	log.Println("send ", msg.GetHeader().TypeId)
	ctx.HandleWrite(buffer)
}

type MsgDecoder struct {
}

func (d MsgDecoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	session := ctx.Channel().Attachment().(Session)
	buffer := bytes.NewBuffer(utils.MustToBytes(message))
	header := &MsgHeader{}
	err := header.Decode(buffer)
	if err != nil {
		panic(err)
	}
	msg, err := CreateMsg(header, session, buffer)
	if err != nil {
		panic(err)
	}
	log.Println("recv:", msg.GetHeader().TypeId)
	ctx.HandleRead(msg)
}

type SessionFactory interface {
	CreateSession(channel netty.Channel) Session
}

type SessionManager interface {
	OnAddSession(session Session)
	OnRemoveSession(session Session)
}

type NodeFactory interface {
	SessionFactory
	SessionManager
}

type LogicHandler struct {
	NodeFactory NodeFactory
}

func (l LogicHandler) HandleActive(ctx netty.ActiveContext) {
	session := l.NodeFactory.CreateSession(ctx.Channel())
	ctx.Channel().SetAttachment(session)
	l.NodeFactory.OnAddSession(session)
	ctx.HandleActive()
}

func (LogicHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	msg := message.(IMsg)
	msg.Dispatch()
	ctx.HandleRead(message)
}

func (l LogicHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	session := ctx.Channel().Attachment().(Session)
	l.NodeFactory.OnRemoveSession(session)
	session.OnClose()
	ctx.HandleInactive(ex)
}

func (LogicHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	log.Println("HandleException:", ex)
}

type SecurityEncoder struct {
	RC4 *rc4.RC4
}

func (SecurityEncoder) CodecName() string {
	return "SecurityEncoder"
}

func (se SecurityEncoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	ctx.HandleRead(message)
}

func (se SecurityEncoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	buf := make([]byte, 1024)
	trans := transform.NewTransformer(func() ([]byte, error) {
		n, err := reader.Read(buf)
		if err != nil {
			return nil, err
		}
		se.RC4.DoUpdate(buf)
		return buf[:n], nil
	})

	ctx.HandleWrite(trans)
}

type SecurityDecoder struct {
	RC4 *rc4.RC4
}

func (SecurityDecoder) CodecName() string {
	return "SecurityDecoder"
}

func (sd SecurityDecoder) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	buf := make([]byte, 1024)
	trans := transform.NewTransformer(func() ([]byte, error) {
		n, err := reader.Read(buf)
		if err != nil {
			return nil, err
		}
		sd.RC4.DoUpdate(buf)
		return buf[:n], nil
	})
	ctx.HandleRead(trans)
}

func (sd SecurityDecoder) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	ctx.HandleWrite(message)
}

func RandomKey(size int) []byte {
	res := make([]byte, size)
	rand.Read(res)
	return res
}
