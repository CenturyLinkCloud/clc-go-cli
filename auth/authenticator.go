package auth

import (
	"errors"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"io/ioutil"
	"log"
	"os"
)

type credentials struct {
	user, password, err string
}

func AuthenticateCommand(opt *options.Options, conf *config.Config) (cn base.Connection, err error) {
	creds := make([]credentials, 0)
	creds = append(creds, credentials{opt.User, opt.Password, "Both --user and --password options should be specified."})
	creds = append(creds, credentials{os.Getenv("CLC_USER"), os.Getenv("CLC_PASSWORD"), "Both CLC_USER and CLC_PASSWORD environment variables should be specified."})
	profile := opt.Profile
	if profile == "" {
		profile = os.Getenv("CLC_PROFILE")
	}
	if profile != "" {
		if p, ok := conf.Profiles[profile]; ok {
			creds = append(creds, credentials{p.User, p.Password, fmt.Sprintf("Incorrect profile '%s'. Both User and Password should be specified.", profile)})
		} else {
			return nil, fmt.Errorf("Profile '%s' doesn't exist.", profile)
		}
	}
	creds = append(creds, credentials{conf.User, conf.Password, "Incorrect config. Both User and Password should be specified."})

	loggerDestination := ioutil.Discard
	trace := opt.Trace
	if !trace {
		trace = os.Getenv("CLC_TRACE") == "true"
	}
	if trace {
		loggerDestination = os.Stdout
	}
	logger := log.New(loggerDestination, "", log.LstdFlags)

	for _, c := range creds {
		if c.user != "" || c.password != "" {
			if c.user == "" || c.password == "" {
				return nil, errors.New(c.err)
			}
			return connection.NewConnection(c.user, c.password, opt.AccountAlias, logger)
		}
	}

	return nil, fmt.Errorf("No credentials provided. Use 'clc login --help' to view list of all authentication options.")
}
