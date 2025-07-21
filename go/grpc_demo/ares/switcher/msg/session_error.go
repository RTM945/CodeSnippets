package msg

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
	"fmt"
)

type SessionError struct {
	*ares.Msg
	pb *pb.SessionError
}

func NewSessionError() *SessionError {
	return &SessionError{
		Msg: ares.NewMsg(6),
		pb:  &pb.SessionError{},
	}
}

func (msg *SessionError) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *SessionError) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *SessionError) TypedPB() *pb.SessionError {
	return msg.pb
}

func (msg *SessionError) String() string {
	return fmt.Sprintf("SessionError[type=%d, pvId=%d, pb=%v]", msg.GetType(), msg.GetPvId(), msg.pb.String())
}
