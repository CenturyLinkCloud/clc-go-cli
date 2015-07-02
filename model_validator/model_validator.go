package model_validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

func ValidateModel(model interface{}) error {
	_, err := govalidator.ValidateStruct(model)
	if err != nil {
		return err
	}
	if m, ok := model.(base.ValidatableModel); ok {
		err = m.Validate()
	}
	return err
}
