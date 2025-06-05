package linker

import (
	"ares/pkg/io"
	pb "ares/proto/gen"
)

type MsgCreatorFunc func(session io.Session, envelope *pb.Envelope) (io.Msg, error)

var MsgCreator = map[uint32]MsgCreatorFunc{}
