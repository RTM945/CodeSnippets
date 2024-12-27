package echo

import (
	"fmt"
	shard "reares/cmd/go-netty"
	"reares/cmd/go-netty/proto/echo"
)

type Logic struct {
	session shard.Session
}

func GetEchoLogic(session shard.Session) *Logic {
	return &Logic{
		session: session,
	}
}

func (logic *Logic) Echo(msg string) error {
	secho := echo.NewSEcho()
	secho.Msg = msg
	return logic.session.Send(secho)
}

func (logic *Logic) TestSEcho(msg string) error {
	fmt.Println(msg)
	return nil
}
