package base

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func ParseArguments(inputModel interface{}, args []string) error {

	if reflect.TypeOf(inputModel).Kind != reflect.Ptr {
		panic("non pointer type")
	}
	elemType := reflect.TypeOf(inputModel).Elem()
	switch elemType.Kind {
	case reflect.Bool:
		val := inputModel.(*bool)
		*val, err = strconv.ParseBool(args[0])
		if err != nil {
			return err
		}
		args = args[1:]
	case reflect.Int:
		val := inputModel.(*int)
		*val, err = strconv.ParseInt(args[0], 0, 0)
		if err != nil {
			return err
		}
		args = args[1:]
	case reflect.String:
		val := inputModel.(*string)
		*val = args[0]
		if err != nil {
			return err
		}
		args = args[1:]
	case reflect.Slice:
		for !strings.HasPrefix(args[0], '-') {
			item := reflect.New(elemType.Elem())
			err = ParseArguments(item, args)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:

	}
}

func parseObject(inputModel interface{}, args []string) {

}
