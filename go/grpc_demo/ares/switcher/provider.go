package switcher

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"ares/switcher/msg"
	"sync"
)

type Provider struct {
	sessions             ares.ISessions
	msgCreator           ares.IMsgCreator
	msgProcessor         ares.IMsgProcessor
	sessionTimeout       int64
	brokenSessionTimeout int64
	auPvIds              sync.Map
	phantomPvIds         sync.Map

	//云服务器不能直接使用外网ip监听
	providerIp string
	port       int32

	pb.UnimplementedProviderServer
}

func (p *Provider) Sessions() ares.ISessions {
	return p.sessions
}

func (p *Provider) MsgCreator() ares.IMsgCreator {
	return p.msgCreator
}

func (p *Provider) MsgProcessor() ares.IMsgProcessor {
	return p.msgProcessor
}

func (p *Provider) ClientBroken(clientSid int32, linkerSession *LinkerSession) {
	if linkerSession == nil {
		for _, v := range provider.Sessions().AllSessions() {
			providerSession := v.(*ProviderSession)
			providerSession.SessionBroken(clientSid)
		}
	} else {
		for _, pvId := range linkerSession.GetBindProvidees() {
			providerSession := provider.Sessions().GetSession(pvId).(*ProviderSession)
			if providerSession != nil {
				clientBroken := msg.NewClientBroken()
				clientBroken.TypedPB().ClientSid = clientSid
				_ = providerSession.Send(clientBroken)
			}
		}
	}
}
func (p *Provider) ProvideeBroken(pvId int32) bool {
	sendByPhantom := false
	p.phantomPvIds.Range(func(key, value interface{}) bool {
		phantomPvId := key.(int32)
		provideeBroken := msg.NewProvideeBroken()
		provideeBroken.TypedPB().PvId = pvId
		provideeBroken.TypedPB().Provider = &pb.ProviderInfo{
			Ip:   provider.providerIp,
			Port: provider.port,
		}
		p.SendToProvidee(phantomPvId, provideeBroken)
		sendByPhantom = true
		return true
	})
	return sendByPhantom
}

func (p *Provider) AddAUPvId(auPvId int32) {
	p.auPvIds.Store(auPvId, struct{}{})
	LOGGER.Infof("Bind AU pvId %d", auPvId)
}

func (p *Provider) RemoveAUPvId(auPvId int32) {
	p.auPvIds.Delete(auPvId)
	LOGGER.Infof("UnBind AU pvId %d", auPvId)
}

func (p *Provider) AddPhantomPvId(phantomPvId int32) {
	p.phantomPvIds.Store(phantomPvId, struct{}{})
	LOGGER.Infof("Bind Phantom PvId %d", phantomPvId)
}

func (p *Provider) RemovePhantomPvId(phantomPvId int32) {
	p.phantomPvIds.Delete(phantomPvId)
	LOGGER.Infof("UnBind Phantom PvId %d", phantomPvId)
}

func (p *Provider) RemovePhantomGS(serverId int32) {

}

func (p *Provider) SendToProvidee(pvId int32, msg ares.IMsg) {
	ps := p.sessions.GetSession(pvId).(*ProviderSession)
	if ps == nil {
		LOGGER.Errorf("Switcher To Providee, No Providee, pvid: %d", pvId)
		return
	}
	ps.SendAsync(msg)
}
