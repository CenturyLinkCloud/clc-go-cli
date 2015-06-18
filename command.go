package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Command interface {
	Execute(cn base.Connection) error
	Resource() string
	Command() string
	InputModel() interface{}
	OutputModel() interface{}
}
