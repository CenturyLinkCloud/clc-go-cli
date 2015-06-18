package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

func ApplyDefaultBehaviour(model interface{}) error {
	if m, ok := model.(base.AdjustableModel); ok {
		return m.ApplyDefaultBehaviour()
	}
	return nil
}
