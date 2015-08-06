package command_loader_test

import (
	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"testing"
)

type command struct {
	resource string
	command  string
}

func (c *command) Execute(cn base.Connection) error {
	return nil
}

func (c *command) Resource() string {
	return c.resource
}

func (c *command) Command() string {
	return c.command
}

func (c *command) ShowHelp() string {
	return ""
}

func (c *command) InputModel() interface{} {
	return nil
}

func (c *command) OutputModel() interface{} {
	return nil
}

var cmd1, cmd2 base.Command

func init() {
	cli.AllCommands = make([]base.Command, 0)
	cmd1 = &command{
		resource: "resource1",
		command:  "command1",
	}
	cmd2 = &command{
		resource: "resource2",
		command:  "command2",
	}
	cli.AllCommands = append(cli.AllCommands, cmd1)
	cli.AllCommands = append(cli.AllCommands, cmd2)
}

func TestLoadExistingCommand(t *testing.T) {
	resource, err := command_loader.LoadResource("resource2")
	if err != nil {
		t.Error(err.Error())
	}
	cmd, err := command_loader.LoadCommand(resource, "command2")
	if err != nil {
		t.Error(err.Error())
	}
	if cmd != cmd2 {
		t.Error("cmd2 expected")
	}
}

func TestResourceNotFound(t *testing.T) {
	_, err := command_loader.LoadResource("resource3")
	if err == nil || err.Error() != "Resource not found: 'resource3'. Use 'clc --help' to list all available resources." {
		t.Errorf("Incorrect error %s", err)
	}
}

func TestCommandNotFound(t *testing.T) {
	resource, err := command_loader.LoadResource("resource2")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = command_loader.LoadCommand(resource, "")
	if err == nil || err.Error() != "Command should be specified. Use 'clc resource2 --help' to list all avaliable commands." {
		t.Errorf("Incorrect error %s", err)
	}
}
