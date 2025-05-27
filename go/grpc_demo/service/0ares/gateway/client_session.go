package gateway

import (
	"grpc_demo/io"
	gatewaypb "grpc_demo/proto/gen/gateway/v1"
	"sync"
)

var clientSessions sync.Map

type Session struct {
	stream  gatewaypb.Gateway_RouteServer
	recvCh  chan *gatewaypb.Envelope
	sendCh  chan *gatewaypb.Envelope
	process chan io.Msg
}

func NewSession(stream gatewaypb.Gateway_RouteServer) *Session {
	return &Session{
		stream: stream,
	}
}

func (session *Session) Send(msg *gatewaypb.Envelope) {

}
