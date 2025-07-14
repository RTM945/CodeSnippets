package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"golang.org/x/time/rate"
	"net"
	"time"
)

// LinkerSession client<->linker
type LinkerSession struct {
	*ares.Session
	minRateLimiter *rate.Limiter
	maxRateLimiter *rate.Limiter
	aliveTime      int64
	bindProvidees  []uint32
}

func NewLinkerSession(stream pb.Linker_ServeServer, node ares.INode) *LinkerSession {
	session := &LinkerSession{
		Session:       ares.NewSession(stream, node),
		aliveTime:     time.Now().Unix(),
		bindProvidees: make([]uint32, 0),
	}
	linker := session.Session.Node().(*Linker)
	if linker.rateMin > 0 {
		session.minRateLimiter = rate.NewLimiter(rate.Limit(linker.rateMin), 1)
	}
	if linker.rateMax > 0 {
		session.maxRateLimiter = rate.NewLimiter(rate.Limit(linker.rateMax), 1)
	}
	return session
}

func (s *LinkerSession) HandleEnvelope(envelope *pb.Envelope) {
	m, err := s.Session.Node().MsgCreator().Create(s, envelope.GetPvId(), envelope.GetTypeId(), envelope.GetPayload())
	if err != nil {
		LOGGER.Errorf("session[%v] create err: %v", s, err)
		return
	}
	if m != nil {
		m.Dispatch()
	}
}

func (s *LinkerSession) receiveUnknown(typeId uint32) {
	s.ResetAlive()
	if s.maxRateLimiter != nil && !s.maxRateLimiter.Allow() {
		s.CloseBySessionError(uint32(pb.SessionError_RATE_LIMIT))
		return
	}
	if s.minRateLimiter != nil && !s.minRateLimiter.Allow() {
		LOGGER.Warnf("min rate limiter, type: %d, session: %v", typeId, s)
	}
}

func (s *LinkerSession) WhiteFilterByProvider(ps *ProviderSession) bool {
	if !ps.WhiteFilter() {
		return false
	}
	host, _, err := net.SplitHostPort(ps.RemoteAddr().String())
	if err == nil {
		linker := s.Session.Node().(*Linker)
		for ip := range linker.GetWhiteIps() {
			if host == ip {
				return false
			}
		}
	} else {
		LOGGER.Errorf("client ip: %v, err: %v", s.RemoteAddr(), err)
	}
	return true
}

func (s *LinkerSession) BlackFilterByProvider(ps *ProviderSession) bool {
	if !ps.BlackFilter() {
		return false
	}
	host, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err == nil {
		linker := s.Session.Node().(*Linker)
		for ip := range linker.GetBlackIps() {
			if host == ip {
				return false
			}
		}
	} else {
		LOGGER.Errorf("client ip: %v, err: %v", s.RemoteAddr(), err)
	}
	return true
}

func (s *LinkerSession) Alive() bool {
	linker := s.Session.Node().(*Linker)
	return time.Now().Unix()-s.aliveTime < linker.sessionTimeout
}

func (s *LinkerSession) ResetAlive() {
	s.aliveTime = time.Now().Unix()
}

func (s *LinkerSession) BindProvidee(pvId uint32) {
	s.bindProvidees = append(s.bindProvidees, pvId)
}

func (s *LinkerSession) GetBindProvidees() []uint32 {
	return s.bindProvidees
}

func (s *LinkerSession) CloseBySessionError(code uint32) {
	sessionError := msg.NewSessionError()
	sessionError.TypedPB().Code = code
	_ = s.Send0(sessionError)
	s.Close()
}
