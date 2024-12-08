package echo

import (
	"fmt"
	"reares/proto"
)

func ProcessEcho(echo *proto.Echo) error {
	fmt.Println(echo.Msg)
	return nil
}
