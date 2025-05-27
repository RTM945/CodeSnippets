package hello

import (
	hellomsg "grpc_demo/common/msg/gen/hello"
)

func init() {
	hellomsg.ProcessCHello = ProcessCHello
}

func ProcessCHello(msg *hellomsg.CHello) error {

	return nil
}
