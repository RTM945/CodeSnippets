package linker

import (
	ares "ares/pkg/io"
	pb "ares/proto/gen"
)

type MsgCreatorFunc func(session ares.ISession, envelope *pb.Envelope) (ares.IMsg, error)

var MsgCreator = map[uint32]MsgCreatorFunc{}
