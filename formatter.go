package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/formatters"
)

type Formatter interface {
	FormatOutput(model interface{}) (string, error)
}

func GetOutputFormatter(options *Options) Formatter {
	return &formatters.JsonFormatter{}
}
