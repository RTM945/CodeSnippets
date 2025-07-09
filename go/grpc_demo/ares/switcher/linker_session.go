package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"golang.org/x/time/rate"
)

// LinkerSession client<->linker
type LinkerSession struct {
	*ares.Session
	node           ares.INode
	minRateLimiter *rate.Limiter
	maxRateLimiter *rate.Limiter
}

func NewLinkerSession(stream pb.Linker_ServeServer, node ares.INode) *LinkerSession {
	session := &LinkerSession{
		Session: ares.NewSession(stream),
		node:    node,
	}
	if linkerRateMin > 0 {
		session.minRateLimiter = rate.NewLimiter(rate.Limit(linkerRateMin), 1)
	}
	if linkerRateMin > 0 {
		session.maxRateLimiter = rate.NewLimiter(rate.Limit(linkerRateMax), 1)
	}
	return session
}

func (s *LinkerSession) HandleEnvelope(envelope *pb.Envelope) {
	m, err := s.node.MsgCreator().Create(s, envelope)
	if err != nil {
		LOGGER.Errorf("session[%v] create err: %v", s, err)
		return
	}
	if m != nil {
		m.Dispatch()
	}
}

func (s *LinkerSession) receiveUnknown(typeId uint32) {
	if s.maxRateLimiter != nil && !s.maxRateLimiter.Allow() {
		s.CloseBySessionError(pb.SessionError_RATE_LIMIT)
		return
	}
	if s.minRateLimiter != nil && !s.minRateLimiter.Allow() {
		LOGGER.Warnf("min rate limiter, type: %d, session: %v", typeId, s)
	}
}

func (s *LinkerSession) CloseBySessionError(code pb.SessionError_Code) {
	sessionError := msg.NewSessionError()
	sessionError.TypedPB().Code = code
	_ = s.Send0(sessionError)
	s.Close()
}
