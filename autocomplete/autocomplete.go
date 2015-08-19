package autocomplete

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
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
	cmd, err := command_loader.LoadCommand(resource, cmdArg)
	if err != nil {
		if len(args) == 2 {
			return strings.Join(command_loader.GetCommands(resource), " ")
		}
		return ""
	}

	var arguments []string
	if cmd.Command() == "" {
		arguments = args[1:]
	} else {
		arguments = args[2:]
	}
	if len(arguments) == 0 {
		return strings.Join(optionsAndArguments(cmd), " ")
	}
	parsed, err := parser.ParseArguments(arguments)
	if err != nil {
		return ""
	}
	yes, _, err := options.AreToBeTakenFromFile(parsed)
	if yes || err != nil {
		return ""
	}
	_, err = options.ExtractFrom(parsed)
	if err != nil {
		// Do not show anything if option needs a value.
		return ""
	}
	last := args[len(args)-1]
	if strings.HasPrefix(last, "--") {
		if hasArg(cmd.InputModel(), parser.NormalizePropertyName(last)) {
			// TODO autocomplete property values
			return ""
		}
	}
	return strings.Join(optionsAndArguments(cmd), " ")
}

func optionsAndArguments(command base.Command) []string {
	opts := options.Get()
	args := command.Arguments()
	return append(opts, args...)
}

func hasArg(m interface{}, f string) bool {
	meta := reflect.ValueOf(m).Elem()
	if meta.FieldByName(f).IsValid() {
		return true
	}
	return false
}