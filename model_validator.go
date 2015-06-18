package cli

import (
	"github.com/altoros/century-link-cli/base"
	"github.com/asaskevich/govalidator"
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
