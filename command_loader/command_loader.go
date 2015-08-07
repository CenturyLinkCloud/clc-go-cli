package command_loader

import (
	"fmt"
	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"strings"
)

func LoadResource(resource string) (string, error) {
	for _, cmd := range cli.AllCommands {
		if cmd.Resource() == resource {
			return resource, nil
		}
	}
	return "", fmt.Errorf("Resource not found: '%s'. Use 'clc --help' to list all available resources.", resource)
}

func LoadCommand(resource, command string) (base.Command, error) {
	for _, cmd := range cli.AllCommands {
		if cmd.Resource() == resource && (cmd.Command() == "" || cmd.Command() == command) {
			return cmd, nil
		}
	}
	if command == "" {
		return nil, fmt.Errorf("Command should be specified. Use 'clc %s --help' to list all avaliable commands.", resource)
	}
	return nil, fmt.Errorf("Command %s %s not found. Use 'clc %s --help' to list all avaliable commands.", resource, command, resource)
}

func GetResources() string {
	resources := []string{}
	m := map[string]bool{}
	for _, cmd := range cli.AllCommands {
		m[cmd.Resource()] = true
	}
	for k, _ := range m {
		resources = append(resources, fmt.Sprintf("  %s", k))
	}
	return strings.Join(resources, "\n")
}

func GetCommands(resource string) string {
	commands := []string{}
	m := map[string]bool{}
	for _, cmd := range cli.AllCommands {
		if cmd.Resource() == resource {
			if cmd.Command() == "" {
				return ""
			}
			m[cmd.Command()] = true
		}
	}
	for k, _ := range m {
		commands = append(commands, k)
	}
	return strings.Join(commands, "\n")
}
