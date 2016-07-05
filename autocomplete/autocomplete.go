package autocomplete

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/auth"
	"github.com/centurylinkcloud/clc-go-cli/autocomplete/cache"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
	"strings"
	"time"
)

const (
	SEP = "\n"
)

func Run(args []string) string {
	if len(args) == 0 {
		return strings.Join(command_loader.GetResources(), SEP)
	}
	resource, err := command_loader.LoadResource(args[0])
	if err != nil {
		if len(args) == 1 {
			return strings.Join(command_loader.GetResources(), SEP)
		}
		return ""
	}
	if len(args) == 1 {
		return strings.Join(command_loader.GetCommands(resource), SEP)
	}

	cmdArg := args[1]
	cmd, err := command_loader.LoadCommand(resource, cmdArg)
	if err != nil {
		if len(args) == 2 {
			return strings.Join(command_loader.GetCommands(resource), SEP)
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
		return strings.Join(optionsAndArguments(cmd), SEP)
	}

	parsed, err := parser.ParseArguments(arguments)
	if err != nil {
		return ""
	}
	yes, _, err := options.AreToBeTakenFromFile(parsed)
	if yes || err != nil {
		return ""
	}

	conf, err := config.LoadConfig()
	if err != nil {
		conf = &config.Config{}
	}

	last := args[len(args)-1]
	opts, err := options.ExtractFrom(parsed)
	if err != nil {
		if last == "--output" {
			return strings.Join([]string{"json", "table", "text", "csv"}, SEP)
		} else if last == "--profile" {
			profiles := []string{}
			for k := range conf.Profiles {
				profiles = append(profiles, k)
			}
			return strings.Join(profiles, SEP)
		}
		return ""
	}
	if strings.HasPrefix(last, "--") {
		key := parser.NormalizePropertyName(last)
		if hasArg(cmd.InputModel(), key) {
			// Looking for enums.
			enum, exist := model_validator.FieldOptions(cmd.InputModel(), key)
			if exist {
				return strings.Join(enum, SEP)
			}

			// Resolving API-related property names.
			if inferable, ok := cmd.InputModel().(base.IDInferable); ok {
				cn, err := auth.AuthenticateCommand(opts, conf)
				if err != nil {
					return ""
				}

				datacenter.ApplyDefault(inferable, conf)

				// Due to the fact API requests may take a long time we cache
				// the results for some short amount of time.
				names, inCache := cache.Get(cacheKey(resource, cmd.Command(), key))
				if !inCache {
					stop := make(chan bool)
					// The following routine repeatdly sends dots to stdout
					// what may serve as a waiting indicator in shells that
					// support such kind of interaction. The output may be
					// supressed in those that do not.
					go wait(stop)
					names, err = inferable.GetNames(cn, key)
					stop <- true
					if err != nil {
						return ""
					}
					cache.Put(cacheKey(resource, cmd.Command(), key), names)
				}
				if names != nil {
					return strings.Join(names, SEP)
				}
			}
			return ""
		}
	}
	return strings.Join(optionsAndArguments(cmd), SEP)
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

func cacheKey(r, c, k string) string {
	return fmt.Sprintf("%s-%s-%s", r, c, k)
}

func wait(stop <-chan bool) {
	printed := 0
	for {
		select {
		case <-stop:
			for printed > 0 {
				fmt.Print("\b \b")
				printed -= 1
			}
			return
		default:
			if printed < 3 {
				fmt.Print(".")
				printed += 1
			} else {
				fmt.Print("\b\b\b   \b\b\b")
				printed = 0
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}
