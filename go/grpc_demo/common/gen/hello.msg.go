package hellomsg

import (
	"context"
	"google.golang.org/protobuf/proto"
	hellopb "grpc_demo/proto/gen/hello/v1"
)

type HelloRequest struct {
	pb     *hellopb.HelloRequest
	typeId uint32
	pvId   uint32
	ctx    context.Context
}

func NewHelloRequest() *HelloRequest {
	return &HelloRequest{
		typeId: 1,
		pb:     &hellopb.HelloRequest{},
	}
}

func (msg *HelloRequest) Marshal() ([]byte, error) {
	return msg.pb.MarshalVT()
}

func (msg *HelloRequest) Unmarshal(bytes []byte) error {
	return msg.pb.UnmarshalVT(bytes)
}

func (msg *HelloRequest) GetType() uint32 { return msg.typeId }

func (msg *HelloRequest) GetPvId() uint32 { return msg.pvId }

func (msg *HelloRequest) GetContext() context.Context { return msg.ctx }

func (msg *HelloRequest) GetPB() proto.Message {
	return msg.pb
}

func (msg *HelloRequest) TypedPB() *hellopb.HelloRequest {
	return msg.pb
}

func (msg *HelloRequest) Process() error {
	// ProcessRegister[msg.typeId].Process(msg)
	return nil
}
