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
	if !metaModel.IsValid() {
		return nil
	}
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
	switch field.Interface().(type) {
	case int64:
		var argInt int64
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Int {
			argInt = arg.(int64)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.Float64 {
			if valid.IsWhole(arg.(float64)) {
				argInt = int64(arg.(float64))
				mismatch = false
			}
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if valid.IsInt(arg.(string)) {
				argInt, _ = valid.ToInt(arg.(string))
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be integer.", key)
		} else {
			field.SetInt(argInt)
			return nil
		}
	case float64:
		var argFloat64 float64
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Float64 {
			argFloat64 = arg.(float64)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if valid.IsFloat(arg.(string)) {
				argFloat64, _ = valid.ToFloat(arg.(string))
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be float.", key)
		} else {
			field.SetFloat(argFloat64)
			return nil
		}
	case time.Time:
		var argTime time.Time
		var err error
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.String {
			if argTime, err = time.Parse(timeFormat, arg.(string)); err == nil {
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be datetime in `YYYY-MM-DD hh:mm:ss` format.", key)
		} else {
			field.Set(reflect.ValueOf(argTime))
			return nil
		}
	case bool:
		var argBool bool
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Bool {
			argBool = arg.(bool)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if arg == "true" {
				argBool = true
				mismatch = false
			} else if arg == "false" {
				argBool = false
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be either true or false.", key)
		} else {
			field.SetBool(argBool)
		}
		return nil
	case string:
		if reflect.ValueOf(arg).Kind() != reflect.String {
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
