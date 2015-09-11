package group

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type List struct {
	DataCenter string
	All        base.NilField
}

func (l *List) Validate() error {
	if !l.All.Set && l.DataCenter == "" {
		return fmt.Errorf("DataCenter: non zero value required")
	}
	return nil
}
