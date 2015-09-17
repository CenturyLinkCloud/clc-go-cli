package commands

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/options"
)

type Login struct {
	CommandBase
}

type inputStub struct{}

func NewLogin(info CommandExcInfo) *Login {
	l := Login{}
	l.ExcInfo = info
	return &l
}

func (l *Login) InputModel() interface{} {
	return &inputStub{}
}

func (l *Login) Login(opts *options.Options, conf *config.Config) string {
	if opts.User == "" || opts.Password == "" {
		return "Both --user and --password options must be specified."
	}

	conf.User = opts.User
	conf.Password = opts.Password
	if err := config.Save(conf); err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Logged in as %s.", opts.User)
}
