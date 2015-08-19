package parser

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"reflect"
	"strings"
)

type Filter struct {
	Cond  string
	Value string
}

func (f *Filter) Match(value interface{}) (bool, error) {
	meta := reflect.ValueOf(value).Kind()
	if meta != reflect.String && meta != reflect.Bool && meta != reflect.Float64 {
		return false, fmt.Errorf("Only strings, numbers and booleans are supported in filters.")
	}
	switch f.Cond {
	case "^=", "$=", "~=":
		if str, ok := value.(string); !ok {
			return false, fmt.Errorf("Operations ~=, ^= and $= can only be used with strings.")
		} else {
			switch f.Cond {
			case "^=":
				return strings.HasPrefix(str, f.Value), nil
			case "$=":
				return strings.HasSuffix(str, f.Value), nil
			case "~=":
				return strings.Contains(str, f.Value), nil
			}
		}
	case "<", ">", "<=", ">=", "=":
		if meta == reflect.Bool && f.Cond != "=" {
			return false, fmt.Errorf("Operations <,>,<= and >= can not be used with booleans.")
		}
		switch meta {
		case reflect.String:
			v := value.(string)
			fV := f.Value
			switch f.Cond {
			case "<":
				return v < fV, nil
			case ">":
				return v > fV, nil
			case "<=":
				return v <= fV, nil
			case ">=":
				return v >= fV, nil
			case "=":
				return v == fV, nil
			}
		case reflect.Float64:
			if f.Value == "" {
				return false, fmt.Errorf("Non-empty value is required for the number.")
			}
			v := value.(float64)
			fV, err := govalidator.ToFloat(f.Value)
			if err != nil {
				return false, fmt.Errorf("Invalid value for the number: %v.", f.Value)
			}
			// Here we repeat the piece of code from above because comparison operators
			// in Go do not work with interfaces.
			switch f.Cond {
			case "<":
				return v < fV, nil
			case ">":
				return v > fV, nil
			case "<=":
				return v <= fV, nil
			case ">=":
				return v >= fV, nil
			case "=":
				return v == fV, nil
			}
		case reflect.Bool:
			if f.Value == "" {
				return false, fmt.Errorf("Non-empty value is required for the boolean.")
			}
			if f.Value != "true" && f.Value != "false" {
				return false, fmt.Errorf("Invalid value for the boolean: %v.", f.Value)
			}
			v := value.(bool)
			fV, _ := govalidator.ToBoolean(f.Value)
			return v == fV, nil
		}
	}
	return false, nil
}

func ParseFilter(input interface{}, filter string) (interface{}, error) {
	parsed, err := ParseFilterObject(filter)
	if err != nil {
		return nil, err
	}
	return parseFilter(input, parsed)
}

func parseFilter(model interface{}, filter map[string]Filter) (interface{}, error) {
	if slice, ok := model.([]interface{}); ok {
		result := []interface{}{}
		for _, el := range slice {
			parsedEl, err := parseFilter(el, filter)
			if err != nil {
				return nil, err
			}
			if parsedEl != nil {
				result = append(result, parsedEl)
			}
		}
		if len(result) == 0 {
			return nil, nil
		}
		return result, nil
	} else if m, ok := model.(map[string]interface{}); ok {
		for k, v := range filter {
			key := k
			if _, ok := m[k]; !ok {
				key = NormalizePropertyName(k)
				if _, ok := m[key]; !ok {
					return nil, fmt.Errorf("%s: there is no such field in result.", key)
				}
			}
			ok, err := v.Match(m[key])
			if err != nil {
				return nil, err
			} else if !ok {
				return nil, nil
			}
		}
		return m, nil
	}
	return nil, nil
}
