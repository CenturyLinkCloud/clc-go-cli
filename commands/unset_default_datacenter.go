package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/config"
)

type UnsetDefaultDC struct {
	CommandBase
}

func NewUnsetDefaultDC(info CommandExcInfo) *UnsetDefaultDC {
	u := UnsetDefaultDC{}
	u.ExcInfo = info
	return &u
}

func (u *UnsetDefaultDC) IsOffline() bool {
	return true
}

func (u *UnsetDefaultDC) ExecuteOffline() (string, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	conf.DefaultDataCenter = ""
	config.Save(conf)
	return "The default data center is unset.", nil
}
