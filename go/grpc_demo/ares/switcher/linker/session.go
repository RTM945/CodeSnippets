package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"net"
	"sync/atomic"
)

// Session client<->linker
type Session struct {
	ares.State
	stream      pb.Linker_ServeServer
	sid         uint32
	remoteAddr  net.Addr
	ctx         context.Context
	cancel      context.CancelFunc
	sendChan    chan *pb.Envelope
	processChan chan ares.Msg
}

var genSessionId uint32

var chanSize = 64

func NewLinkerSession(stream pb.Linker_ServeServer) *Session {
	session := &Session{
		sendChan:    make(chan *pb.Envelope, chanSize),
		processChan: make(chan ares.Msg, chanSize),
		stream:      stream,
		sid:         atomic.AddUint32(&genSessionId, 1),
	}
	if p, ok := peer.FromContext(stream.Context()); ok {
		session.remoteAddr = p.Addr
	}
	session.ctx, session.cancel = context.WithCancel(stream.Context())
	return session
}

func (s *Session) String() string {
	return fmt.Sprintf("[sid = %d, remoteAddr = %s]", s.sid, s.remoteAddr)
}

func (s *Session) Context() context.Context {
	return s.ctx
}

func (s *Session) GetSid() uint32 {
	return s.sid
}

func (s *Session) Process(msg ares.Msg) {
	s.processChan <- msg
}

func (s *Session) Send(msg ares.Msg) error {
	payload, err := msg.Marshal()
	if err != nil {
		return err
	}

	envelope := &pb.Envelope{
		TypeId:  msg.GetType(),
		PvId:    msg.GetPvId(),
		Payload: payload,
	}

	s.sendChan <- envelope
	return nil
}

func (s *Session) startProcess() {
	defer LOGGER.Infof("session[%v] process goroutine stopped", s)
	defer func() {
		if err := recover(); err != nil {
			LOGGER.Errorf("session[%v] process goroutine panic: %v", s, err)
		}
	}()
	for {
		if err := s.ctx.Err(); err != nil {
			for m := range s.processChan {
				if err := m.Process(); err != nil {
					LOGGER.Errorf("session[%v] process remaining msg %v err: %v", s, m, err)
				}
			}
			return
		}

		select {
		case m := <-s.processChan:
			err := m.Process()
			if err != nil {
				LOGGER.Errorf("session[%v] process %v err: %v", s, m, err)
			}
		}
	}
}

func (s *Session) startSend() {
	defer LOGGER.Infof("session[%v] send goroutine stopped", s)
	defer func() {
		if err := recover(); err != nil {
			LOGGER.Errorf("session[%v] send goroutine panic: %v", s, err)
		}
	}()
	for {
		if err := s.ctx.Err(); err != nil {
			// session close 后就不发了
			return
		}
		select {
		case envelope := <-s.sendChan:
			if err := s.stream.Send(envelope); err != nil {
				LOGGER.Errorf("session[%v] send err: %v", s, err)
				return
			}
		}
	}
}

func (s *Session) HandleEnvelope(envelope *pb.Envelope) {
	creator, ok := MsgCreator[envelope.TypeId]
	if ok {
		// 自己处理
		m, err := creator(s, envelope)
		if err != nil {
			LOGGER.Errorf("session[%v] unmarshal err: %v", s, err)
			return
		}
		m.Dispatch()
	} else {
		if envelope.PvId != 0 {
			// 通过PvId获取对应的服务器进行转发
			// 网关把客户端请求转发给具体的业务服处理
		}

	}
}

func (s *Session) Close() {
	s.cancel()
}

func (s *Session) OnClose() {

}
