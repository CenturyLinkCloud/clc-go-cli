package model_adjuster

import (
	"reflect"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
)

func ApplyDefaultBehaviour(model interface{}) error {
	adjustEnums(model)
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

// adjustEnums allows for case-insensitivity.
func adjustEnums(model interface{}) {
	meta := reflect.ValueOf(model)
	if meta.Kind() != reflect.Ptr {
		return
	}
	meta = meta.Elem()
	if meta.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < meta.NumField(); i++ {
		t := meta.Type().FieldByIndex([]int{i})
		opts, exist := model_validator.FieldOptions(model, t.Name)
		if !exist {
			continue
		}
		f := meta.FieldByIndex([]int{i})
		if f.String() == "" {
			continue
		}
		for _, o := range opts {
			if strings.ToLower(o) == strings.ToLower(f.String()) {
				f.SetString(o)
				break
			}
		}
	}
}
