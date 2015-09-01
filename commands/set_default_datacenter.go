package commands

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

type SetDefaultDC struct {
	CommandBase
}

func NewSetDefaultDC(input interface{}, info CommandExcInfo) *SetDefaultDC {
	s := SetDefaultDC{}
	s.ExcInfo = info
	s.Input = input
	return &s
}

func (s *SetDefaultDC) IsOffline() bool {
	return true
}

func (s *SetDefaultDC) ExecuteOffline() (string, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	m, ok := s.Input.(*datacenter.SetDefault)
	if !ok {
		panic("Input model must be of type *datacenter.SetDefault.")
	}

	conf.DefaultDataCenter = m.DataCenter
	config.Save(conf)
	return fmt.Sprintf("%s is now the default data center.", m.DataCenter), nil
}
