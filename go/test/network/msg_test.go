package network

import (
	"github.com/golang/protobuf/proto"
	"gotest/network/protobuf"
	"testing"
)

func TestMsg(t *testing.T) {
	echo := protobuf.Echo{Msg: "test"}
	data, _ := proto.Marshal(&echo)
	msg, err := CreateMsg(1, data)
	if err != nil {
		t.Fatal(err)
	}
	if echo, ok := msg.(*Echo); ok {
		t.Log("Echo:", echo)
	}
}
