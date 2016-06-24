package model_validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
	clc_errors "github.com/centurylinkcloud/clc-go-cli/errors"
	"github.com/centurylinkcloud/clc-go-cli/parser"
)

func ValidateModel(model interface{}) error {
	err := validateStruct(model)
	if err != nil {
		return err
	}

	if m, ok := model.(base.ValidatableModel); ok {
		err = m.Validate()
	}

	return err
}

func validateStruct(s interface{}) error {
	if s == nil {
		return nil
	}

	var err error
	var errMsgs []string

	err = validateWithGovalidator(s)
	if err != nil {
		errMsgs = append(errMsgs, err.Error())
	}

	err = validateEnums(s)
	if err != nil {
		errMsgs = append(errMsgs, err.Error())
	}

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			err = validateStruct(field.Interface())
			if err != nil {
				err = fmt.Errorf("The %s field has following errors:\n%s", strings.ToLower(v.Type().Field(i).Name), err)
				errMsgs = append(errMsgs, err.Error())
			}
		}
	}

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, "\n"))
	}
	return nil
}

func validateWithGovalidator(s interface{}) error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		parts := strings.Split(err.Error(), ";")
		errors := []string{}
		for _, p := range parts[:len(parts)-1] {
			if len(p) == 0 {
				continue
			}
			errors = append(errors, pretifyError(p))
		}
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
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

	var errs []string

	numFields := typ.NumField()
OverFields:
	for i := 0; i < numFields; i++ {
		name := typ.FieldByIndex([]int{i}).Name
		opts, exist := FieldOptions(model, name)
		// fmt.Println(name, "exist=", exist, "opts=", opts)
		if exist {
			field := v.FieldByIndex([]int{i})
			fieldValue := field.String()
			if fieldValue == "" && opts[len(opts)-1] == "optional" {
				// fmt.Println("field is optional")
				continue
			}
			for _, o := range opts {
				if strings.ToLower(o) == strings.ToLower(fieldValue) {
					continue OverFields
				}
			}
			errMsg := fmt.Sprintf("%s value must be one of %s", strings.ToLower(name), strings.Join(opts, ", "))
			errs = append(errs, errMsg)
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
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
		return clc_errors.EmptyField(strings.TrimPrefix(parser.DenormalizePropertyName(field), "--")).Error()
	}
	return err
}
