package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
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
	Brief    string
	Help     string
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

func (c *CommandBase) ShowBrief() string {
	return c.ExcInfo.Brief
}

func (c *CommandBase) ShowHelp() string {
	return c.ExcInfo.Brief + c.ExcInfo.Help
}

func (c *CommandBase) InputModel() interface{} {
	return c.Input
}

func (c *CommandBase) OutputModel() interface{} {
	return c.Output
}
