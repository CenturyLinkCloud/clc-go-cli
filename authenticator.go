package cli

import (
	"github.com/altoros/century-link-cli/base"
	"io/ioutil"
	"log"
)

func AuthenticateCommand(opt *Options) (cn base.Connection, err error) {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	return newConnection(opt.User, opt.Password, logger)
}
