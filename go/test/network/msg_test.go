package network

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"gotest/network/protobuf"
	"testing"
)

func TestMsg(t *testing.T) {
	echo := protobuf.Echo{Msg: "test"}
	data, _ := proto.Marshal(&echo)
	header := &MsgHeader{
		TypeId: 1,
		PvId:   0,
	}
	msg, err := CreateMsg(header, nil, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	if echo, ok := msg.(*Echo); ok {
		t.Log("Echo:", echo)
		echo.Process()
	}
}
