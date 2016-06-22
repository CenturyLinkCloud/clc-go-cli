package command_loader

import (
	"fmt"
	"sort"

	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/help"
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

func GetResources() []string {
	resources := []string{}
	m := map[string]bool{}
	for _, cmd := range cli.AllCommands {
		m[cmd.Resource()] = true
	}
	for k, _ := range m {
		resources = append(resources, k)
	}
	return resources
}

func GetCommands(resource string) []string {
	commands := []string{}
	m := resourceCommandsInfo(resource)
	if m == nil {
		return []string{""}
	}
	for k := range m {
		commands = append(commands, k)
	}
	return commands
}

func GetCommandsWithDescriptions(resource string) string {
	commands := []help.Argument{}
	m := resourceCommandsInfo(resource)
	if m == nil {
		return ""
	}
	sortedKeys := []string{}
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		commands = append(commands, help.Argument{
			Name:        k,
			Description: m[k].ShowBrief(),
		})
	}
	return help.ForResource(help.Resource{
		Name:     resource,
		Commands: commands,
	})
}

func resourceCommandsInfo(resource string) map[string]base.Command {
	m := map[string]base.Command{}
	for _, cmd := range cli.AllCommands {
		if cmd.Resource() == resource {
			if cmd.Command() == "" {
				return nil
			}
			m[cmd.Command()] = cmd
		}
	}
	return m
}
