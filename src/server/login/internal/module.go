package internal

import (
	"server/thrift/thrift-go-gen/leader"
	"github.com/luojinbo008/thrift_manage/module"
	"fmt"
	"context"
)

type Module struct {
	*module.Thrift
}

const (
	HOST = "localhost"
	PORT = 8080
)

type leaderThrift struct {}

func (fdi *leaderThrift) Register(ctx context.Context, data *leader.Leader) (x bool, err error){

	fmt.Println("11111111111111")
	return true, nil
}

func (m *Module) OnInit() {

	handler := &leaderThrift{}
	processor := leader.NewLeaderThriftProcessor(handler)

	m.Thrift = &module.Thrift{
		Host:	HOST,
		Port:	PORT,
		TProcessor:	processor,
	}
}

func (m *Module) OnDestroy() {

}

func (m *Module) SetModuleName() (string) {
	return "login"
}