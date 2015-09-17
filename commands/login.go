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
	if opts.User == "" && opts.Password == "" && opts.Profile == "" {
		return "Either a profile or a user and a password must be specified."
	}

	if (opts.User == "") != (opts.Password == "") {
		return "Both --user and --password options must be specified."
	}

	var user, password string
	if opts.User != "" {
		user, password = opts.User, opts.Password
	} else {
		var profile config.Profile
		ok := false
		if profile, ok = conf.Profiles[opts.Profile]; !ok {
			return fmt.Sprintf("Profile %s does not exist.", opts.Profile)
		}
		user, password = profile.User, profile.Password
	}

	conf.User = user
	conf.Password = password
	if err := config.Save(conf); err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Logged in as %s.", user)
}
