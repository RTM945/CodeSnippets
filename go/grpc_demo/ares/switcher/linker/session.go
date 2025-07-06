package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
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
			toSession, ok := s.linker.provider.GetSessions().GetSession(envelope.PvId).(*provider.Session)
			if toSession == nil || !ok {
				LOGGER.Errorf("Client to Providee, No Providee exist, pvid: %d, session: %v, typeId: %d", envelope.PvId, s, envelope.TypeId)
				serverError := msg.NewServerError()
				serverError.TypedPB().PvId = envelope.PvId
				serverError.TypedPB().Code = pb.ServerError_SERVER_NOT_ACCESSIBLE
				_ = s.Send0(serverError)
				return
			}
			// 后端服务还没准备好
			if toSession.CheckToProvide() && !s.CheckState(int(pb.ClientState_TOPROVIDEE)) {
				LOGGER.Errorf("Client to Providee, state error: %d, session: %v", s.GetState(), s)
				s.linker.CloseSession(s, pb.SessionError_CANT_DISPATCH)
				return
			}
			// 白名单
			if s.WhiteFilterByProvider(toSession) {
				LOGGER.Errorf("providee writeip kick, pvid: %d, session: %v, typeId: %d", envelope.PvId, s, envelope.TypeId)
				s.linker.CloseSession(s, pb.SessionError_OPEN_WHITE_IP)
				return
			}
			// 黑名单
			if s.BlackFilterByProvider(toSession) {
				LOGGER.Errorf("providee blackip kick, pvid: %d, session: %v, typeId: %d", envelope.PvId, s, envelope.TypeId)
				s.linker.CloseSession(s, pb.SessionError_OPEN_BLACK_IP)
				return
			}
			s.receiveUnknown(envelope.TypeId)
			// TODO Statistics
			err := s.SendToProvidee(envelope, toSession)
			if err != nil {
				LOGGER.Errorf("session[%v] sendToProvidee msg=%v err: %v", s, envelope, err)
			}
		}
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

func (s *Session) SendToProvidee(envelope *pb.Envelope, toSession *provider.Session) error {
	// TODO Statistics
	dispatch := msg.NewDispatch()
	dispatch.TypedPB().PvId = envelope.PvId
	dispatch.TypedPB().TypeId = envelope.TypeId
	dispatch.TypedPB().Payload = envelope.Payload
	return toSession.Send(dispatch)
}
