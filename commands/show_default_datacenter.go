package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/config"
)

type ShowDefaultDC struct {
	CommandBase
}

func NewShowDefaultDC(info CommandExcInfo) *ShowDefaultDC {
	u := ShowDefaultDC{}
	u.ExcInfo = info
	return &u
}

func (u *ShowDefaultDC) IsOffline() bool {
	return true
}

func (u *ShowDefaultDC) ExecuteOffline() (string, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	if conf.DefaultDataCenter == "" {
		return "No data center is currently set as default.", nil
	}

	return conf.DefaultDataCenter, nil
}
