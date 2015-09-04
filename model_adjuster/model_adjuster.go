package model_adjuster

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

func ApplyDefaultBehaviour(model interface{}) error {
	if m, ok := model.(base.AdjustableModel); ok {
		return m.ApplyDefaultBehaviour()
	}
	return nil
}

func InferID(model interface{}, cn base.Connection) error {
	if named, ok := model.(base.IDInferable); ok {
		if err := named.InferID(cn); err != nil {
			return err
		}
	}
	return nil
}
