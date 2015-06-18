package cli

import (
	"fmt"
)

func LoadCommand(resource, command string) (Command, error) {
	resourceFound := false
	for _, cmd := range AllCommands {
		if cmd.Resource() == resource {
			resourceFound = true
		}
		if cmd.Resource() == resource && cmd.Command() == command {
			return cmd, nil
		}
	}

	if !resourceFound {
		return nil, fmt.Errorf("Resource not found: %s.", resource)
	}

	if command == "" {
		return nil, fmt.Errorf("Command should be specified. User 'clc %s --help' to list all avaliable commands.")
	}
	return nil, fmt.Errorf("Command %s %s not found. User 'clc %s --help' to list all avaliable commands.", resource, command)
}
