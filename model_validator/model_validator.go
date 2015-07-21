package model_validator

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"strings"
)

func ValidateModel(model interface{}) error {
	_, err := govalidator.ValidateStruct(model)
	if err != nil {
		var msg string
		parts := strings.Split(err.Error(), ";")
		errors := parts[:len(parts)-1]
		for _, fieldErr := range errors {
			msg += fmt.Sprintf("%s\n", fieldErr)
		}
		return fmt.Errorf(msg)
	}
	if m, ok := model.(base.ValidatableModel); ok {
		err = m.Validate()
	}
	return err
}
