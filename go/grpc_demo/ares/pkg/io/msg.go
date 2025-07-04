package io

type IMsg interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetType() uint32
	SetType(uint32)
	GetPvId() uint32
	SetPvId(pvId uint32)
	SetContext(ctx any)
	Process() error
	Dispatch()
	GetSession() ISession
	SetSession(session ISession)
}

type Msg struct {
	pvId    uint32
	typeId  uint32
	ctx     any
	session ISession
}

func NewMsg(typeId uint32) *Msg {
	return &Msg{
		typeId: typeId,
	}
}

func (msg *Msg) Marshal() ([]byte, error) {
	panic("implement me")
}

func (msg *Msg) Unmarshal(data []byte) error {
	panic("implement me")
}

func (msg *Msg) GetType() uint32 { return msg.typeId }

func (msg *Msg) SetType(typeId uint32) {
	msg.typeId = typeId
}

func (msg *Msg) GetPvId() uint32 { return msg.pvId }

func (msg *Msg) SetPvId(pvId uint32) {
	msg.pvId = pvId
}

func (msg *Msg) GetContext() any { return msg.ctx }

func (msg *Msg) SetContext(ctx any) { msg.ctx = ctx }

func (msg *Msg) Dispatch() {
	msg.session.Process(msg)
}

func (msg *Msg) Process() error {
	panic("implement me")
}

func (msg *Msg) SetSession(session ISession) {
	msg.session = session
}

func (msg *Msg) GetSession() ISession { return msg.session }

type MsgCreatorFunc func(session ISession, pvId, typeId uint32, payload []byte) (IMsg, error)

type IMsgCreator interface {
	Create(session ISession, pvId, typeId uint32, payload []byte) (IMsg, error)
	Register(id uint32, f MsgCreatorFunc)
}
