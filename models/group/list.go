package group

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/errors"
)

type List struct {
	DataCenter string
	All        base.NilField
}

func (l *List) Validate() error {
	if !l.All.Set && l.DataCenter == "" {
		return errors.EmptyField("data-center")
	}
	return nil
}
