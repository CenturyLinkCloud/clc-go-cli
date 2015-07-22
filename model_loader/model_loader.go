package model_loader

import (
	"encoding/json"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func LoadModel(parsedArgs map[string]interface{}, inputModel interface{}) error {
	metaModel := reflect.ValueOf(inputModel)
	if metaModel.Kind() != reflect.Ptr {
		return fmt.Errorf("Input model must be passed by pointer.")
	}
	for k, v := range parsedArgs {
		field, err := getFieldByName(metaModel, k)
		if err != nil {
			return err
		}
		err = loadValue(k, v, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadValue(key string, arg interface{}, field reflect.Value) error {
	if arg == nil {
		return nil
	}
	isArgSlice := reflect.ValueOf(arg).Kind() == reflect.Slice
	switch field.Interface().(type) {
	case int64:
		if isArgSlice || !valid.IsInt(arg.(string)) {
			return fmt.Errorf("Type mismatch: %s value must be integer.", key)
		} else {
			argInt, _ := valid.ToInt(arg.(string))
			field.SetInt(argInt)
			return nil
		}
	case float64:
		if isArgSlice || !valid.IsFloat(arg.(string)) {
			return fmt.Errorf("Type mismatch: %s value must be float.", key)
		} else {
			argFloat, _ := valid.ToFloat(arg.(string))
			field.SetFloat(argFloat)
			return nil
		}
	case time.Time:
		mismatch := isArgSlice
		var argTime time.Time
		var err error
		if !mismatch {
			if argTime, err = time.Parse(timeFormat, arg.(string)); err != nil {
				mismatch = true
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be datetime in `YYYY-MM-DD hh:mm:ss` format.", key)
		} else {
			field.Set(reflect.ValueOf(argTime))
			return nil
		}
	case bool:
		if arg == "true" {
			field.SetBool(true)
		} else if arg == "false" {
			field.SetBool(false)
		} else {
			return fmt.Errorf("Type mismatch: %s value must be either true or false.", key)
		}
		return nil
	case string:
		if isArgSlice {
			return fmt.Errorf("Type mismatch: %s value must be string.", key)
		}
		field.SetString(arg.(string))
		return nil
	}
	if isStruct(field) {
		argStruct, err := parseStruct(arg)
		if err != nil {
			return err
		}
		for k, v := range argStruct {
			nestedField, err := getFieldByName(field.Addr(), k)
			if err != nil {
				return err
			}
			err = loadValue(k, v, nestedField)
			if err != nil {
				return err
			}
		}
		return nil
	} else if isSlice(field) {
		argSlice, err := parseSlice(arg)
		if err != nil {
			return err
		}
		for _, v := range argSlice {
			elementPtr := getEmplySliceType(field)
			err = loadValue(key, v, elementPtr.Elem())
			if err != nil {
				return err
			}
			field.Set(reflect.Append(field, elementPtr.Elem()))
		}
		return nil
	}
	return fmt.Errorf("Unsupported field type %s", field.Kind())
}

func getFieldByName(model reflect.Value, name string) (reflect.Value, error) {
	field := model.Elem().FieldByName(name)
	if !field.IsValid() {
		return reflect.ValueOf(nil), fmt.Errorf("Field `%s` does not exist.", name)
	}
	return field, nil
}

func isStruct(model reflect.Value) bool {
	return model.Kind() == reflect.Struct
}

func isSlice(model reflect.Value) bool {
	return model.Kind() == reflect.Slice
}

// Parses an object of type map[string]interface{} either from JSON or from a=b,c=d,.. notation.
// Also, calls NormalizeKeys with the parsed object.
// If arg is already of type map[string]interface{} returns it as is.
func parseStruct(arg interface{}) (map[string]interface{}, error) {
	if argMap, isMap := arg.(map[string]interface{}); isMap {
		return argMap, nil
	}
	parsed := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(arg.(string)), &parsed); err == nil {
		parser.NormalizeKeys(parsed)
		return parsed, nil
	}
	if parsed, err := parser.ParseObject(arg.(string)); err == nil {
		return parsed, nil
	}
	return nil, fmt.Errorf("`%s` is neither in JSON nor in key=value,.. format.", arg.(string))
}

// Parses an object of type []interface{} either from JSON.
// Also, calls NormalizeKeys with the parsed object.
// If arg is already of type []interface{} returns it as is.
func parseSlice(arg interface{}) ([]interface{}, error) {
	if argSlice, isSlice := arg.([]interface{}); isSlice {
		return argSlice, nil
	}
	parsed := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(arg.(string)), &parsed); err == nil {
		parser.NormalizeKeys(parsed)
		return parsed, nil
	}
	return nil, fmt.Errorf("`%s` is neither in JSON nor in key=value,.. format.", arg.(string))
}

func getEmplySliceType(slice reflect.Value) reflect.Value {
	return reflect.New(slice.Type().Elem())
}
