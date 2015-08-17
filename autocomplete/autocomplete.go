package autocomplete

import (
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"strings"
)

func Run(args []string) string {
	if len(args) == 0 {
		return strings.Join(command_loader.GetResources(), " ")
	}
	resource, err := command_loader.LoadResource(args[0])
	if err != nil {
		if len(args) == 1 {
			return strings.Join(command_loader.GetResources(), " ")
		}
		return ""
	}
	if len(args) == 1 {
		return strings.Join(command_loader.GetCommands(resource), " ")
	}
	cmdArg := args[1]
	_, err = command_loader.LoadCommand(resource, cmdArg)
	if err != nil {
		if len(args) == 2 {
			return strings.Join(command_loader.GetCommands(resource), " ")
		}
		return ""
	}
	return ""
}
