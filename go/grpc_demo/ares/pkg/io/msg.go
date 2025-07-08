package io

import "errors"

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
	Register(id uint32, f MsgCreatorFunc)
	Create(session ISession, pvId, typeId uint32, payload []byte) (IMsg, error)
}

type MsgCreator struct {
	register map[uint32]MsgCreatorFunc
}

func NewMsgCreator() *MsgCreator {
	return &MsgCreator{
		register: make(map[uint32]MsgCreatorFunc),
	}
}

func (mc *MsgCreator) Register(id uint32, f MsgCreatorFunc) {
	mc.register[id] = f
}

var NoMsgCreatorErr = errors.New("no msg creator")

func (mc *MsgCreator) Create(session ISession, pvId, typeId uint32, payload []byte) (IMsg, error) {
	creator, ok := mc.register[typeId]
	if !ok {
		return nil, NoMsgCreatorErr
	}
	return creator(session, pvId, typeId, payload)
}

type MsgProcessorFunc[T IMsg] func(T) error

type IMsgProcessor interface {
	Register(typeId uint32, processor func(IMsg) error)
	Process(m IMsg) error
}

type MsgProcessor struct {
	register map[uint32]func(msg IMsg) error
}

func NewMsgProcessor() *MsgProcessor {
	return &MsgProcessor{
		register: make(map[uint32]func(msg IMsg) error),
	}
}

func (mp *MsgProcessor) Register(typeId uint32, processor func(IMsg) error) {
	RegisterMsgProcessor(mp, typeId, processor)
}

var NoMsgProcessorErr = errors.New("no msg processor")

func (mp *MsgProcessor) Process(m IMsg) error {
	h, ok := mp.register[m.GetType()]
	if !ok {
		return NoMsgProcessorErr
	}
	return h(m)
}

func RegisterMsgProcessor[T IMsg](mp IMsgProcessor, typeId uint32, processor MsgProcessorFunc[T]) {
	mp.Register(typeId, func(m IMsg) error {
		typed, ok := m.(T)
		if !ok {
			return nil
		}
		return processor(typed)
	})
}
