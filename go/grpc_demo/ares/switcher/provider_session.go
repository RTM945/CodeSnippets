package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"fmt"
	"time"
)

type ProviderSession struct {
	*ares.Session
	checkToProvidee  bool
	brokenSessionIds map[uint32]int64
	aliveTime        int64
	provideeInfo     *pb.ProvideeInfo
}

func NewProviderSession(stream pb.Provider_ServeServer, node ares.INode) *ProviderSession {
	return &ProviderSession{
		Session:          ares.NewSession(stream, node),
		brokenSessionIds: make(map[uint32]int64),
	}
}

func (ps *ProviderSession) CheckToProvide() bool {
	return ps.checkToProvidee
}

func (ps *ProviderSession) WhiteFilter() bool {
	return ps.CheckState(int(pb.ProvideeState_WHITEIP))
}

func (ps *ProviderSession) BlackFilter() bool {
	return ps.CheckState(int(pb.ProvideeState_BLACKIP))
}

func (ps *ProviderSession) SessionBroken(brokenSessionId uint32) {
	if _, ok := ps.brokenSessionIds[brokenSessionId]; !ok {
		ps.brokenSessionIds[brokenSessionId] = time.Now().Unix()
		LOGGER.Infof("Add a broeken session, sessionId=%d", brokenSessionId)
		clientBroken := msg.NewClientBroken()
		clientBroken.TypedPB().ClientSid = brokenSessionId
		_ = ps.Send(clientBroken)
	}
}

func (ps *ProviderSession) Alive() bool {
	provider := ps.Session.Node().(*Provider)
	return time.Now().Unix()-ps.aliveTime < provider.sessionTimeout
}

func (ps *ProviderSession) Check() {
	now := time.Now().Unix()
	provider := ps.Session.Node().(*Provider)
	for k, v := range ps.brokenSessionIds {
		if now-v > provider.sessionTimeout {
			LOGGER.Infof("Removed a broken session, sessionId=%d", k)
			delete(ps.brokenSessionIds, k)
		}
	}
	if len(ps.brokenSessionIds) > 0 {
		LOGGER.Infof("Now broken clientsids.size=%d", len(ps.brokenSessionIds))
	}
}

func (ps *ProviderSession) GetPvId() int32 {
	if ps.provideeInfo == nil {
		return -1
	}
	return int32(ps.provideeInfo.PvId)
}

func (ps *ProviderSession) Send(msg ares.IMsg) error {
	pvId := ps.GetPvId()
	if 0 == msg.GetPvId() {
		if 0 == pvId {
			return fmt.Errorf("not Bind Providee: %v, msg: %v", ps, msg)
		}
		msg.SetPvId(uint32(pvId))
	}
	return ps.Session.Send(msg)
}

func (ps *ProviderSession) String() string {
	return fmt.Sprintf("ProviderSession: %s, session: %s", ps.provideeInfo, ps.Session)
}

func (ps *ProviderSession) IsAUSession() bool {
	return ps.provideeInfo.ServerType == uint32(pb.ServerType_AU)
}

func (ps *ProviderSession) IsPhantomSession() bool {
	return ps.provideeInfo.ServerType == uint32(pb.ServerType_PHANTOM)
}

func (ps *ProviderSession) IsGameServerSession() bool {
	return ps.provideeInfo.ServerType == uint32(pb.ServerType_LOGIC)
}

func (ps *ProviderSession) GetServerId() uint32 {
	return ps.provideeInfo.ServerId
}
