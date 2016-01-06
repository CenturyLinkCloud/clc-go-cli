package model_validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/errors"
	"github.com/centurylinkcloud/clc-go-cli/parser"
)

func ValidateModel(model interface{}) error {
	if model == nil {
		return nil
	}
	_, err := govalidator.ValidateStruct(model)
	if err != nil {
		parts := strings.Split(err.Error(), ";")
		errors := []string{}
		for _, p := range parts[:len(parts)-1] {
			errors = append(errors, pretifyError(p))
		}
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	err = validateEnums(model)
	if err != nil {
		return err
	}
	if m, ok := model.(base.ValidatableModel); ok {
		err = m.Validate()
	}
	return err
}

func FieldOptions(model interface{}, name string) ([]string, bool) {
	meta := reflect.ValueOf(model)
	if meta.Kind() == reflect.Ptr {
		meta = meta.Elem()
	}
	v, typ := meta, meta.Type()

	ft, exists := typ.FieldByName(name)
	if !exists {
		return nil, false
	}
	fv := v.FieldByName(name)
	if fv.Kind() != reflect.String {
		return nil, false
	}

	tag := ft.Tag.Get("oneOf")
	if tag == "" {
		return nil, false
	}

	return strings.Split(tag, ","), true
}

func validateEnums(model interface{}) error {
	var v reflect.Value
	var typ reflect.Type

	meta := reflect.ValueOf(model)
	if meta.Kind() == reflect.Ptr {
		meta = meta.Elem()
	}
	if meta.Kind() != reflect.Struct {
		panic("A struct or a pointer to a struct is expected.")
	}
	v, typ = meta, meta.Type()

	numFields := typ.NumField()
OverFields:
	for i := 0; i < numFields; i++ {
		name := typ.FieldByIndex([]int{i}).Name
		opts, exist := FieldOptions(model, name)
		if exist {
			field := v.FieldByIndex([]int{i})
			if field.String() == "" {
				continue
			}
			for _, o := range opts {
				if strings.ToLower(o) == strings.ToLower(field.String()) {
					continue OverFields
				}
			}
			return fmt.Errorf("%s value must be one of %s.", name, strings.Join(opts, ", "))
		}
	}
	return nil
}

func pretifyError(err string) string {
	// FIXME Since error messages can't be customized in govalidator we parse
	// messages manually here.
	parts := strings.Split(err, ":")
	field := parts[0]
	msg := parts[1]
	if strings.Contains(msg, "non zero value required") {
		return errors.EmptyField(strings.TrimPrefix(parser.DenormalizePropertyName(field), "--")).Error()
	}
	return err
}
