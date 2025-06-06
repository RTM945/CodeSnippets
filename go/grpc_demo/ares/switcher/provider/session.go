package provider

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"io"
	"sync/atomic"
)

// Session linker<->provider
type Session struct {
	ares.State
	stream      pb.Provider_ServeServer
	sid         uint32
	sendChan    chan *pb.Envelope
	processChan chan ares.Msg
	closeChan   chan struct{}
}

var genSessionId uint32

var chanSize = 64

func NewProviderSession(stream pb.Provider_ServeServer) *Session {
	return &Session{
		sendChan:    make(chan *pb.Envelope, chanSize),
		processChan: make(chan ares.Msg, chanSize),
		closeChan:   make(chan struct{}),
		stream:      stream,
		sid:         atomic.AddUint32(&genSessionId, 1),
	}
}

func (s *Session) GetSid() uint32 {
	return s.sid
}

func (s *Session) Process(msg ares.Msg) {
	s.processChan <- msg
}

func (s *Session) StartProcess() {
	for m := range s.processChan {
		func() {
			defer func() {
				if r := recover(); r != nil {
					LOGGER.Errorf("session[%d] process %v panic: %v", s.sid, m, r)
				}
			}()
			err := m.Process()
			if err != nil {
				LOGGER.Errorf("session[%d] process %v err: %v", s.sid, m, err)
			}
		}()
	}
}

func (s *Session) Recv() error {
	for {
		envelope, err := s.stream.Recv()
		if err != nil {
			if err == io.EOF {
				LOGGER.Info("client closed stream")
				return nil
			}
		}
		creator, ok := MsgCreator[envelope.TypeId]
		if ok {
			// 自己处理
			m, err := creator(s, envelope)
			if err != nil {
				LOGGER.Errorf("session[%d] unmarshal err: %v", s.sid, err)
				continue
			}
			m.Dispatch()
		} else {
			if envelope.PvId != 0 {
				// 通过PvId获取对应的服务器进行转发
				// 这里是provider之间的转发 比如gs发给map
			}
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
	s.sendChan <- envelope
	return nil
}

func (s *Session) StartSend() {
	for envelope := range s.sendChan {
		if err := s.stream.Send(envelope); err != nil {
			LOGGER.Errorf("session[%d] send err: %v", s.sid, err)
			break
		}
	}
	LOGGER.Infof("session[%d] stop send", s.sid)
}

func (s *Session) Close() {
	close(s.closeChan)
}

func (s *Session) OnClose() {
	close(s.processChan)
	close(s.sendChan)
}
