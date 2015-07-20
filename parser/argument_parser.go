package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ParseArguments(args []string) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	if len(args) == 0 {
		return res, nil
	}

	var key string
	isSettingKey := false
	for i := 0; i < len(args); i++ {
		if args[i] == "--" {
			return nil, fmt.Errorf("-- is an invalid argument.")
		}
		if strings.HasPrefix(args[i], "--") {
			key = normalizePropertyName(args[i])
			if _, ok := res[key]; ok {
				return nil, fmt.Errorf("Option '%s' is specified twice.", key)
			}
			if i+1 == len(args) || strings.HasPrefix(args[i+1], "--") {
				res[key] = nil
			} else {
				isSettingKey = true
			}
			continue
		}
		if isSettingKey {
			if _, ok := res[key]; !ok {
				res[key] = args[i]
			} else if _, ok := res[key].(string); ok {
				array := make([]interface{}, 0)
				array = append(array, res[key], args[i])
				res[key] = array
			} else {
				res[key] = append(res[key].([]interface{}), args[i])
			}
			continue
		} else {
			jsonArg := map[string]interface{}{}
			err = json.Unmarshal([]byte(args[i]), &jsonArg)
			if err != nil {
				return nil, fmt.Errorf("Invalid JSON: %s.", args[i])
			}
			NormalizeKeys(jsonArg)
			for k, v := range jsonArg {
				if _, ok := res[k]; ok {
					return nil, fmt.Errorf("Option '%s' is specified twice.", k)
				}
				res[k] = v
			}
			continue
		}
	}
	return res, nil
}

func NormalizeKeys(arg interface{}) {
	if argObj, isObj := arg.(map[string]interface{}); isObj {
		for k, v := range argObj {
			n := normalizePropertyName(k)
			delete(argObj, k)
			(argObj)[n] = v
			NormalizeKeys(v)
		}
	} else if argArray, isArray := arg.([]interface{}); isArray {
		for _, v := range argArray {
			NormalizeKeys(v)
		}
	}
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

func normalizeValue(value string) interface{} {
	var obj interface{} = value
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		obj = val
	} else if val, err := strconv.ParseBool(value); err == nil {
		obj = val
	}
	return obj
}

type state func(r rune) error

var curState state
var curQuote rune
var curItem []rune
var items []string

func parseObject(obj string) interface{} {
	curState = startParseKey
	curQuote = '\000'
	curItem = []rune{}
	items = []string{}
	for _, c := range obj {
		err := curState(c)
		if err != nil {
			return normalizeValue(obj)
		}
	}
	curState('\000')
	if len(items) <= 1 {
		return normalizeValue(obj)
	}
	res := make(map[string]interface{}, 0)
	for i := 0; i < len(items); i += 2 {
		key := normalizePropertyName(items[i])
		if i == len(items)-1 {
			res[key] = nil
		} else {
			res[key] = normalizeValue(items[i+1])
		}
	}
	return res
}

func saveCurItem() {
	items = append(items, string(curItem))
	curItem = []rune{}
}

func startParseKey(r rune) error {
	switch r {
	case '\'', '"':
		curQuote = r
		curState = parseQuotedKey
	default:
		curItem = append(curItem, r)
		curState = parseSimpleKey
	}
	return nil
}

func parseSimpleKey(r rune) error {
	switch r {
	case '=':
		curState = startParseValue
		saveCurItem()
	case '\000':
		saveCurItem()
	default:
		curItem = append(curItem, r)
	}
	return nil
}

func parseQuotedKey(r rune) error {
	switch r {
	case curQuote:
		curState = keyParsed
		saveCurItem()
	case '\000':
		return errors.New("")
	default:
		curItem = append(curItem, r)
	}
	return nil
}

func keyParsed(r rune) error {
	switch r {
	case '=':
		curState = startParseValue
		return nil
	case '\000':
		return nil
	default:
		return errors.New("")

	}
}

func startParseValue(r rune) error {
	switch r {
	case '\'', '"':
		curQuote = r
		curState = parseQuotedValue
	default:
		curItem = append(curItem, r)
		curState = parseSimpleValue
	}
	return nil
}

func parseSimpleValue(r rune) error {
	switch r {
	case ',':
		curState = startParseKey
		saveCurItem()
	case '\000':
		saveCurItem()
	default:
		curItem = append(curItem, r)
	}
	return nil
}

func parseQuotedValue(r rune) error {
	switch r {
	case curQuote:
		curState = valueParsed
		saveCurItem()
	case '\000':
		return errors.New("")
	default:
		curItem = append(curItem, r)
	}
	return nil
}

func valueParsed(r rune) error {
	switch r {
	case ',':
		curState = startParseKey
		return nil
	case '\000':
		return nil
	default:
		return errors.New("")

	}
}
