package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

func ParseArguments(args []string) (res map[string]interface{}, err error) {
	if len(args) == 0 {
		return res, nil
	}

	err = json.Unmarshal([]byte(args[0]), res)
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		return nil, fmt.Errorf("Invalid json parameter '%s'", e.Value)
	}
	i := 0
	if err == nil {
		i++
	}

	for ; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			return nil, fmt.Errorf("Invalid option format, option %s should start with --", args[i])
		}
		if i+1 == len(args) {
			res[normalizePropertyName(args[i])] = nil
		} else if i+2 == len(args) || strings.HasPrefix(args[i+2], "--") {
			obj, err := parseObject(args[i+1])
			if err != nil {
				return nil, err
			}
			res[normalizePropertyName(args[i])] = obj
			i++
		} else {
			array := make([]interface{}, 0)
			j := i + 1
			for ; j < len(args) && !strings.HasPrefix(args[j], "--"); j++ {
				obj, err := parseObject(args[i+1])
				if err != nil {
					return nil, err
				}
				array = append(array, obj)
			}
			res[normalizePropertyName(args[i])] = array
			i = j
		}
	}
	return res, nil
}

func normalizePropertyName(prName string) string {
	prName = strings.TrimLeft(prName, "--")
	array := strings.Split(prName, "-")
	res := make([]rune, 0)
	for _, item := range array {
		chars := []rune(item)
		chars[0] = unicode.ToUpper(chars[0])
		res = append(res, chars...)
	}
	return string(res)
}

func parseObject(obj string) (interface{}, error) {
	array := strings.Split(obj, ",")
	if len(array) == 1 {
		return parsePropertyValue(array[0])
	}
	res := make([]interface{}, 0)
	for _, item := range array {
		obj, err := parsePropertyValue(item)
		if err != nil {
			return nil, err
		}
		res = append(res, obj)
	}
	return res, nil
}

func parsePropertyValue(prop string) (interface{}, error) {
	array := strings.Split(prop, "=")
	switch len(array) {
	case 1:
		return prop, nil
	case 2:
		return map[string]string{array[0]: array[1]}, nil
	default:
		return nil, fmt.Errorf("Incorrect property value %s", prop)
	}
}
