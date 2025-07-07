package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/provider"
	"golang.org/x/time/rate"
	"net"
)

// Session client<->linker
type Session struct {
	*ares.Session
	linker         *Linker
	minRateLimiter *rate.Limiter
	maxRateLimiter *rate.Limiter
}

func NewLinkerSession(stream pb.Linker_ServeServer, linker *Linker) *Session {
	session := &Session{
		Session: ares.NewSession(stream),
		linker:  linker,
	}
	if linker.rateMin > 0 {
		session.minRateLimiter = rate.NewLimiter(rate.Limit(linker.rateMin), 1)
	}
	if linker.rateMax > 0 {
		session.maxRateLimiter = rate.NewLimiter(rate.Limit(linker.rateMax), 1)
	}
	return session
}

func (s *Session) HandleEnvelope(envelope *pb.Envelope) {
	msg, err := s.linker.msgCreator.Create(s, envelope)
	if err != nil {
		LOGGER.Errorf("session[%v] create err: %v", s, err)
		return
	}
	if msg != nil {
		msg.Dispatch()
	}
}

func (s *Session) OnClose() {

}

func (s *Session) WhiteFilterByProvider(ps *provider.Session) bool {
	if !ps.WhiteFilter() {
		return false
	}
	host, _, err := net.SplitHostPort(ps.RemoteAddr().String())
	if err == nil {
		for ip := range s.linker.GetWhiteIps() {
			if host == ip {
				return false
			}
		}
	} else {
		LOGGER.Errorf("client ip: %v, err: %v", s.RemoteAddr(), err)
	}
	return true
}

func (s *Session) BlackFilterByProvider(ps *provider.Session) bool {
	if !ps.BlackFilter() {
		return false
	}
	host, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err == nil {
		for ip := range s.linker.GetBlackIps() {
			if host == ip {
				return false
			}
		}
	} else {
		LOGGER.Errorf("client ip: %v, err: %v", s.RemoteAddr(), err)
	}
	return true
}

func (s *Session) receiveUnknown(typeId uint32) {
	if s.maxRateLimiter != nil && !s.maxRateLimiter.Allow() {
		s.linker.CloseSession(s, pb.SessionError_RATE_LIMIT)
		return
	}
	if s.minRateLimiter != nil && !s.minRateLimiter.Allow() {
		LOGGER.Warnf("min rate limiter, type: %d, session: %v", typeId, s)
	}
}
