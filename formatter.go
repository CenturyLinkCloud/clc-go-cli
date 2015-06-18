package cli

import (
	"github.com/altoros/century-link-cli/formatters"
)

type Formatter interface {
	FormatOutput(model interface{}) (string, error)
}

func GetOutputFormatter(options *Options) Formatter {
	return &formatters.JsonFormatter{}
}
