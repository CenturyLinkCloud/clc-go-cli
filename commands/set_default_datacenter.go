package commands

import (
	"fmt"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

type SetDefaultDC struct {
	CommandBase
}

func NewSetDefaultDC(info CommandExcInfo) *SetDefaultDC {
	s := SetDefaultDC{}
	s.ExcInfo = info
	s.Input = &datacenter.SetDefault{}
	return &s
}

func (s *SetDefaultDC) Execute(cn base.Connection) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	m, ok := s.Input.(*datacenter.SetDefault)
	if !ok {
		panic("Input model must be of type *datacenter.SetDefault.")
	}

	datacenters := []datacenter.ListRes{}
	URL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", base.URL)
	err = cn.ExecuteRequest("GET", URL, nil, &datacenters)
	if err != nil {
		return err
	}

	valid := false
	for _, d := range datacenters {
		if strings.ToLower(d.Id) == strings.ToLower(m.DataCenter) {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("%s: there is no data center with such code", m.DataCenter)
	}

	conf.DefaultDataCenter = m.DataCenter
	config.Save(conf)
	success := fmt.Sprintf("%s is now the default data center.", m.DataCenter)
	s.Output = &success
	return nil
}
