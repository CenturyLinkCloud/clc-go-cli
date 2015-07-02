package formatter_provider

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"github.com/centurylinkcloud/clc-go-cli/options"
)

func GetOutputFormatter(options *options.Options) (base.Formatter, error) {
	switch options.Output {
	case "", "json":
		return &formatters.JsonFormatter{}, nil
	case "text":
		return &formatters.TextFormatter{}, nil
	case "table":
		return &formatters.TableFormatter{}, nil
	default:
		return nil, fmt.Errorf("Unknown output '%s'. Must be one of the following: json, table, text.", options.Output)
	}
}
