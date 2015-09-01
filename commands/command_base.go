package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/help"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
)

type CommandBase struct {
	Input   interface{}
	Output  interface{}
	ExcInfo CommandExcInfo
}

type CommandExcInfo struct {
	Verb     string
	Url      string
	Resource string
	Command  string
	Help     help.Command
}

func (c *CommandBase) Execute(cn base.Connection) error {
	return cn.ExecuteRequest(c.ExcInfo.Verb, c.ExcInfo.Url, c.Input, c.Output)
}

func (c *CommandBase) Resource() string {
	return c.ExcInfo.Resource
}

func (c *CommandBase) Command() string {
	return c.ExcInfo.Command
}

func (c *CommandBase) Arguments() []string {
	if c.Input == nil {
		return []string{}
	}

	metaPtr := reflect.TypeOf(c.Input)
	if metaPtr.Kind() != reflect.Ptr {
		return []string{}
	}
	meta := metaPtr.Elem()
	if meta.Kind() != reflect.Struct {
		return []string{}
	}

	args := []string{}
	n := meta.NumField()
	for i := 0; i < n; i++ {
		f := meta.FieldByIndex([]int{i})
		args = append(args, parser.DenormalizePropertyName(f.Name))
	}
	return args
}

func (c *CommandBase) ShowBrief() []string {
	return c.ExcInfo.Help.Brief
}

func (c *CommandBase) ShowHelp() string {
	return help.ForCommand(c.ExcInfo.Help)
}

func (c *CommandBase) InputModel() interface{} {
	return c.Input
}

func (c *CommandBase) OutputModel() interface{} {
	return c.Output
}

func (c *CommandBase) IsOffline() bool {
	return false
}

func (c *CommandBase) ExecuteOffline() (string, error) {
	return "", nil
}
