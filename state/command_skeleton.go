package state

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/centurylinkcloud/clc-go-cli/base"
)

func ArgumentsToJSON(args map[string]interface{}, model interface{}) (string, error) {
	if reflect.ValueOf(model).Kind() != reflect.Invalid {
		meta := reflect.ValueOf(model).Elem()
		numFields := meta.Type().NumField()
		if numFields != 0 {
			for i := 0; i < numFields; i++ {
				name, value := getField(meta, i)
				if _, isNil := value.(base.NilField); isNil {
					continue
				}
				if _, ok := args[name]; !ok {
					args[name] = value
				}
			}
		}
	}

	bytes, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ArgumentsFromJSON(filename string) (map[string]interface{}, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	args := map[string]interface{}{}
	err = json.Unmarshal(bytes, &args)
	if err != nil {
		return nil, err
	}
	return args, nil
}

func getField(m reflect.Value, i int) (name string, value interface{}) {
	name = m.Type().FieldByIndex([]int{i}).Name
	value = m.FieldByIndex([]int{i}).Interface()
	return
}
