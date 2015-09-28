package model_validator

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"reflect"
	"strings"
)

func ValidateModel(model interface{}) error {
	if model == nil {
		return nil
	}
	_, err := govalidator.ValidateStruct(model)
	if err != nil {
		parts := strings.Split(err.Error(), ";")
		errors := parts[:len(parts)-1]
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

func contains(strs []string, s string) bool {
	for _, el := range strs {
		if s == el {
			return true
		}
	}
	return false
}
