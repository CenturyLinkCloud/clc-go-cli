package cli

import (
	"github.com/altoros/century-link-cli/base"
)

func ApplyDefaultBehaviour(model interface{}) error {
	if m, ok := model.(base.AdjustableModel); ok {
		return m.ApplyDefaultBehaviour()
	}
	return nil
}
