package command_loader

import (
	"fmt"
	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"strings"
)

func LoadCommand(resource, command string) (base.Command, error) {
	resourceFound := false
	for _, cmd := range cli.AllCommands {
		if cmd.Resource() == resource {
			resourceFound = true
		}
		if cmd.Resource() == resource && (cmd.Command() == "" || cmd.Command() == command) {
			return cmd, nil
		}
	}

	if !resourceFound {
		return nil, fmt.Errorf("Command not found: '%s'.", resource)
	}

	if command == "" {
		return nil, fmt.Errorf("Command should be specified. Use 'clc %s --help' to list all avaliable commands.", resource)
	}
	return nil, fmt.Errorf("Command %s %s not found. Use 'clc %s --help' to list all avaliable commands.", resource, command)
}

func GetResources() string {
	resources := []string{}
	m := map[string]bool{}
	for _, cmd := range cli.AllCommands {
		m[cmd.Resource()] = true
	}
	for k, _ := range m {
		resources = append(resources, k)
	}
	return strings.Join(resources, "\n")
}
