package switcher

import (
	ares "ares/pkg/io"
	"ares/pkg/utils"
	pb "ares/proto/gen"
	"ares/switcher/msg"
)

type Provider struct {
	sessions             ares.ISessions
	msgCreator           ares.IMsgCreator
	msgProcessor         ares.IMsgProcessor
	sessionTimeout       int64
	brokenSessionTimeout int64

	auPvIds      utils.CopyOnWriteSet[uint32]
	phantomPvIds utils.CopyOnWriteSet[uint32]

	//云服务器不能直接使用外网ip监听
	providerIp string
	port       uint32

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

func (p *Provider) ClientBroken(clientSid uint32, linkerSession *LinkerSession) {
	if linkerSession == nil {
		for _, v := range provider.Sessions().AllSessions() {
			providerSession := v.(*ProviderSession)
			providerSession.SessionBroken(clientSid)
		}
	} else {
		for _, pvId := range linkerSession.GetBindProvidees() {
			providerSession := provider.Sessions().GetSession(pvId)
			if providerSession != nil {
				clientBroken := msg.NewClientBroken()
				clientBroken.TypedPB().ClientSid = clientSid
				_ = providerSession.Send(clientBroken)
			}
		}
	}
}

func (p *Provider) AddAUPvId(auPvId uint32) {
	p.auPvIds.Add(auPvId)
	LOGGER.Infof("Bind AU pvId %d", auPvId)
}

func (p *Provider) RemoveAUPvId(auPvId uint32) {
	p.auPvIds.Remove(auPvId)
	LOGGER.Infof("UnBind AU pvId %d", auPvId)
}

func (p *Provider) AddPhantomPvId(phantomPvId uint32) {
	p.phantomPvIds.Add(phantomPvId)
	LOGGER.Infof("Bind Phantom PvId %d", phantomPvId)
}

func (p *Provider) RemovePhantomPvId(phantomPvId uint32) {
	p.phantomPvIds.Remove(phantomPvId)
	LOGGER.Infof("UnBind Phantom PvId %d", phantomPvId)
}

func (p *Provider) RemovePhantomGS(serverId uint32) {

}

func (p *Provider) SendToProvidee(pvId uint32, msg ares.IMsg) {
	ps := p.sessions.GetSession(pvId)
	if ps == nil {
		LOGGER.Errorf("Switcher To Providee, No Providee, pvid: %d", pvId)
		return
	}
	ps.(*ProviderSession).Send(msg)
}
