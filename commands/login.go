package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Login struct {
	CommandBase
}

func NewLogin(info CommandExcInfo) *Login {
	l := Login{}
	l.ExcInfo = info
	return &l
}

func (l *Login) Execute(cn base.Connection) error {
	return nil
}
