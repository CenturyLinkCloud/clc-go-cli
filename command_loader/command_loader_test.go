package command_loader_test

import (
	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"reflect"
	"sort"
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

func (c *command) Arguments() []string {
	return []string{}
}

func (c *command) ShowBrief() []string {
	return []string{"A testing command"}
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

func (c *command) IsOffline() bool {
	return false
}

func (c *command) ExecuteOffline() (string, error) {
	return "", nil
}

var cmd1, cmd2, cmd3, cmd4 base.Command

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
	cmd3 = &command{
		resource: "resource3",
		command:  "",
	}
	cmd4 = &command{
		resource: "resource1",
		command:  "command2",
	}
	cli.AllCommands = append(cli.AllCommands, []base.Command{cmd1, cmd2, cmd3, cmd4}...)
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
	_, err := command_loader.LoadResource("resource4")
	if err == nil || err.Error() != "Resource not found: 'resource4'. Use 'clc --help' to list all available resources." {
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

func TestGetResources(t *testing.T) {
	got := command_loader.GetResources()
	expected := []string{"resource1", "resource2", "resource3"}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nInvalid result.\nExpected: %v\nGot: %v", expected, got)
	}
}

func TestGetCommands(t *testing.T) {
	got := command_loader.GetCommands("resource1")
	expected := []string{"command1", "command2"}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nInvalid result.\nExpected: %v\nGot: %v", expected, got)
	}

	got = command_loader.GetCommands("resource3")
	expected = []string{""}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nInvalid result.\nExpected: %v\nGot: %v", expected, got)
	}
}

func TestGetCommandsWithDescriptions(t *testing.T) {
	got := command_loader.GetCommandsWithDescriptions("resource1")
	expected := `Available resource1 commands:

	command1
		A testing command


	command2
		A testing command

`
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nInvalid result.\nExpected: %v\nGot: %v", expected, got)
	}
}
