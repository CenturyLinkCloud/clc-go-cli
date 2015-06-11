package base

import (
	"fmt"
)

func LoadCommand(resource, command string) (Command, error) {
	for _, cmd := range AllCommands {
		if cmd.Resource() == resource && cmd.Command() == command {
			return cmd, nil
		}
	}
	return nil, fmt.Errorf("Command %s %s doesn't exist.", resource, command)
}
