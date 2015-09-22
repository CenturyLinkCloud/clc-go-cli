package parser

import (
	"encoding/json"
	"errors"
	"fmt"
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
			key = NormalizePropertyName(args[i])
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
			parsedArg := map[string]interface{}{}
			err = json.Unmarshal([]byte(args[i]), &parsedArg)
			if err != nil {
				parsedArg, err = ParseObject(args[i])
				if err != nil {
					return nil, fmt.Errorf("%s is neither a valid JSON object nor a valid object in a=b,c=d.. format.", args[i])
				}
			} else {
				NormalizeKeys(parsedArg)
			}
			for k, v := range parsedArg {
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
			n := NormalizePropertyName(k)
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

func NormalizePropertyName(prName string) string {
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

func DenormalizePropertyName(name string) string {
	denormalized := []rune{'-'}
	for i, char := range name {
		var prevIsUpper, nextIsUpper bool
		r := rune(char)

		if i < len(name)-1 && unicode.IsUpper(rune(name[i+1])) {
			nextIsUpper = true
		} else if i > 0 && unicode.IsUpper(rune(name[i-1])) {
			prevIsUpper = true
		}
		if unicode.IsUpper(r) && !prevIsUpper && !nextIsUpper {
			denormalized = append(denormalized, '-')
			denormalized = append(denormalized, unicode.ToLower(r))
		} else {
			denormalized = append(denormalized, r)
		}
	}
	return string(denormalized)
}

type state func(r rune) error

var curState state
var curQuote rune
var curItem []rune
var items []string

var isFilter bool
var nextRune rune
var prevRune rune
var conditions []string

func ParseObject(obj string) (map[string]interface{}, error) {
	curState = startParseKey
	curQuote = '\000'
	curItem = []rune{}
	items = []string{}
	isFilter = false
	for _, c := range obj {
		err := curState(c)
		if err != nil {
			return nil, err
		}
	}
	curState('\000')
	if len(items) <= 1 {
		return nil, errors.New("Object is not in a=b,c=d,.. notation.")
	}
	res := make(map[string]interface{}, 0)
	for i := 0; i < len(items); i += 2 {
		key := NormalizePropertyName(items[i])
		if i == len(items)-1 {
			res[key] = nil
		} else {
			res[key] = items[i+1]
		}
	}
	return res, nil
}

func ParseFilterObject(obj string) (map[string]Filter, error) {
	curState = startParseKey
	curQuote = '\000'
	curItem = []rune{}
	items = []string{}
	isFilter = true
	conditions = []string{}
	for i, c := range obj {
		if i < len(obj)-1 {
			nextRune = rune(obj[i+1])
		}
		if i > 0 {
			prevRune = rune(obj[i-1])
		}
		err := curState(c)
		if err != nil {
			return nil, err
		}
	}
	curState('\000')
	if len(items) <= 1 {
		return nil, errors.New("Invalid filter format.")
	}
	res := make(map[string]Filter, 0)
	for i := 0; i < len(items); i += 2 {
		key := items[i]
		if len(conditions) <= i/2 {
			return nil, fmt.Errorf("Failed to parse the filter: enclose it in quotes")
		}
		if i == len(items)-1 {
			res[key] = Filter{conditions[i/2], ""}
		} else {
			res[key] = Filter{conditions[i/2], items[i+1]}
		}
	}
	return res, nil
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
		if !isFilter {
			curState = parseSimpleKey
		} else {
			curState = parseFilterKey
		}
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

func parseFilterKey(r rune) error {
	if strings.Contains("^$~<>", string(r)) && nextRune == '=' {
		conditions = append(conditions, fmt.Sprintf("%s=", string(r)))
		return nil
	}
	switch r {
	case '<', '>', '=':
		if r != '=' || !strings.Contains("^$~<>", string(prevRune)) {
			conditions = append(conditions, string(r))
		}
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
