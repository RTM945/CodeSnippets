package io

var genSessionId int32

type Session interface {
	Send(msg Msg) error
	GetSid() int32
	OnClose()
}

//type DefaultSession struct {
//	conn net.Conn
//	sid  int32
//}
//
//func NewDefaultSession(conn net.Conn) *DefaultSession {
//	return &DefaultSession{
//		conn: conn,
//		sid:  atomic.AddInt32(&genSessionId, 1),
//	}
//}
//
//func (session *DefaultSession) Send(msg Msg) error {
//	buffer := GetBuffer()
//	err := EncodeMsg(buffer, msg)
//	if err != nil {
//		PutBuffer(buffer)
//		return err
//	}
//	_, err = session.conn.Write(buffer.Bytes())
//	if err != nil {
//		PutBuffer(buffer)
//		return err
//	}
//	PutBuffer(buffer)
//	return nil
//}
//
//func (session *DefaultSession) GetSid() int32 {
//	return session.sid
//}
//
//func (session *DefaultSession) OnClose() {}
//
//type StateSession struct {
//	*DefaultSession
//	state int32
//}
//
//func NewStateSession(conn net.Conn) *StateSession {
//	return &StateSession{
//		DefaultSession: NewDefaultSession(conn),
//	}
//}
//
//func (session *StateSession) AddState(state int32) {
//	atomic.AddInt32(&session.state, state)
//}
//
//func (session *StateSession) GetState() int32 {
//	return atomic.LoadInt32(&session.state)
//}
//
//func (session *StateSession) CheckState(state int32) bool {
//	currentState := atomic.LoadInt32(&session.state)
//	return state == (currentState & state)
//}
//
//func (session *StateSession) RemoveState(state int32) {
//	for {
//		oldState := atomic.LoadInt32(&session.state)
//		newState := oldState &^ state
//		if atomic.CompareAndSwapInt32(&session.state, oldState, newState) {
//			break
//		}
//	}
//}
