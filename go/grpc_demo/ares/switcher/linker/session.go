package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
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
	closeOnce   sync.Once
	wg          sync.WaitGroup
}

var genSessionId uint32

var chanSize = 64

func NewLinkerSession(stream grpc.ServerStream) *Session {
	session := &Session{
		sendChan:    make(chan *pb.Envelope, chanSize),
		processChan: make(chan ares.Msg, chanSize),
		stream:      stream.(pb.Linker_ServeServer),
		sid:         atomic.AddUint32(&genSessionId, 1),
	}
	return session
}

func (s *Session) SetRemoteAddr(addr net.Addr) {
	s.remoteAddr = addr
}

func (s *Session) SetCancel(cancel context.CancelFunc) {
	s.cancel = cancel
}

func (s *Session) String() string {
	return fmt.Sprintf("[sid = %d, remoteAddr = %s]", s.sid, s.remoteAddr)
}

func (s *Session) GetSid() uint32 {
	return s.sid
}

func (s *Session) Start(parentCtx context.Context) {
	s.wg.Add(3)

	go s.startSend()
	go s.startProcess()
	go s.startRecv()

	// 监听父 context 取消
	go func() {
		select {
		case <-parentCtx.Done():
			s.Close()
		case <-s.ctx.Done(): // 添加 session 的退出
		}
	}()

	// 等待所有 goroutine 结束
	s.wg.Wait()
	LOGGER.Infof("session[%d] all goroutines stopped", s.sid)
}

func (s *Session) Process(msg ares.Msg) {
	select {
	case s.processChan <- msg:
	case <-s.ctx.Done():
		LOGGER.Warnf("session[%d] is closed, drop message: %v", s.sid, msg)
	}
}

func (s *Session) startProcess() {
	defer s.wg.Done()
	defer LOGGER.Infof("session[%d] process goroutine stopped", s.sid)
	for {
		select {
		case m, ok := <-s.processChan:
			if !ok {
				return
			}
			err := m.Process()
			if err != nil {
				LOGGER.Errorf("session[%d] process %v err: %v", s.sid, m, err)
			}
		case <-s.ctx.Done():
			// 处理剩余消息
			for {
				select {
				case m, ok := <-s.processChan:
					if !ok {
						return
					}
					if err := m.Process(); err != nil {
						LOGGER.Errorf("session[%d] process remaining msg %v err: %v", s.sid, m, err)
					}
				default:
					return
				}
			}
		}
	}
}

func (s *Session) startRecv() {
	defer s.wg.Done()
	defer LOGGER.Infof("session[%d] recv goroutine stopped", s.sid)
	for {
		envelope, err := s.stream.Recv()
		if err != nil {
			if status.Code(err) == codes.Canceled {
				LOGGER.Infof("session[%d] context canceled", s.sid)
			}
			return
		}
		s.handleEnvelope(envelope)
	}
}

func (s *Session) handleEnvelope(envelope *pb.Envelope) {
	creator, ok := MsgCreator[envelope.TypeId]
	if ok {
		// 自己处理
		m, err := creator(s, envelope)
		if err != nil {
			LOGGER.Errorf("session[%d] unmarshal err: %v", s.sid, err)
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

	select {
	case s.sendChan <- envelope:
		return nil
	case <-s.ctx.Done():
		return fmt.Errorf("session[%d] is closed", s.sid)
	}
}

func (s *Session) startSend() {
	defer s.wg.Done()
	defer LOGGER.Infof("session[%d] send goroutine stopped", s.sid)

	for {
		select {
		case envelope := <-s.sendChan:
			if err := s.stream.Send(envelope); err != nil {
				LOGGER.Errorf("session[%d] send err: %v", s.sid, err)
				return
			}
		case <-s.ctx.Done():
			// 发送剩余消息
			for {
				select {
				case envelope := <-s.sendChan:
					if err := s.stream.Send(envelope); err != nil {
						LOGGER.Errorf("session[%d] send remaining msg err: %v", s.sid, err)
						return
					}
				default:
					return
				}
			}
		}
	}
}

func (s *Session) Close() {
	s.closeOnce.Do(func() {
		LOGGER.Infof("session[%d] closing", s.sid)
		s.cancel()
		close(s.processChan)
		close(s.sendChan)
	})
}

func (s *Session) OnClose() {

}
