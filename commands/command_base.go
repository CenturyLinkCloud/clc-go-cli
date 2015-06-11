package commands

import (
	"github.com/altoros/century-link-cli/base"
)

type CommandBase struct {
	InputModel  interface{}
	OutputModel interface{}
	ExeInfo     CommandExcInfo
}

type CommandExcInfo struct {
	Method   string
	Url      string
	Resource string
	Command  string
}
