package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type ServerList struct {
	CommandBase
}

func NewServerList(info CommandExcInfo) *ServerList {
	sl := ServerList{}
	sl.ExcInfo = info
	sl.Input = &server.List{}
	return &sl
}

func (sl *ServerList) Execute(cn base.Connection) error {
	var code string
	var err error
	input := sl.Input.(*server.List)

	code = input.DataCenter
	if input.All.Set {
		code = "all"
	}

	sl.Output, err = server.Load(cn, code)
	if err != nil {
		return err
	}
	return err
}
