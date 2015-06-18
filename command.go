package cli

import (
	"github.com/altoros/century-link-cli/base"
)

type Command interface {
	Execute(cn base.Connection) error
	Resource() string
	Command() string
	InputModel() interface{}
	OutputModel() interface{}
}
