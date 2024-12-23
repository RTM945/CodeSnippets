package echo

import (
	"fmt"
	"reares/internal/io"
	"reares/proto/echo"
)

type Logic struct {
	session io.Session
}

func GetEchoLogic(session io.Session) *Logic {
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
