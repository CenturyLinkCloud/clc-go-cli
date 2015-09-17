package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
)

type GroupList struct {
	CommandBase
}

func NewGroupList(info CommandExcInfo) *GroupList {
	g := GroupList{}
	g.ExcInfo = info
	g.Input = &group.List{}
	return &g
}

func (g *GroupList) Execute(cn base.Connection) error {
	var err error
	var code string
	input := g.Input.(*group.List)

	code = input.DataCenter
	if input.All.Set {
		code = "all"
	}

	g.Output, err = group.Load(cn, code)
	if err != nil {
		return err
	}
	return nil
}
