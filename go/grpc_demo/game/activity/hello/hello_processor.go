package hello

import hellomsg "grpc_demo/common/gen"

func init() {
	hellomsg.ProcessHelloRequest = ProcessHelloRequest
}

func ProcessHelloRequest(msg *hellomsg.HelloRequest) error {

	return nil
}
