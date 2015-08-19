package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseQuery creates either a map[string]interface{} or []interface{} out of some struct or a slice of structs.
// If the function returns a slice, its elements are either map[string]interface{} or []interface{} recursively.
// Only the keys parsed from the given query string are left in the resulting maps.
func ParseQuery(input interface{}, query string) (interface{}, error) {
	path, fields, aliases, err := parseQueryFields(query)
	if err != nil {
		return nil, err
	}
	var current string
	if len(path) == 0 {
		current = fields[0]
	} else {
		current = path[0]
	}
	return parseModelByQuery(path, fields, input, current, getNextStep(path, current), aliases), nil
}

func ConvertToMapOrSlice(input interface{}) (interface{}, error) {
	var model interface{}
	modelSlice := make([]interface{}, 0)
	modelStruct := make(map[string]interface{}, 0)

	bytes, err := json.Marshal(input)
	if err = json.Unmarshal(bytes, &modelSlice); err == nil {
		model = modelSlice
	} else if err = json.Unmarshal(bytes, &modelStruct); err == nil {
		model = modelStruct
	} else {
		return nil, err
	}
	return model, nil
}

func parseQueryFields(query string) ([]string, []string, map[string]string, error) {
	path, fields := []string{}, []string{}
	aliases := map[string]string{}
	var err error

	withAliases := strings.Split(query, ".{")
	if len(withAliases) > 2 {
		return nil, nil, nil, fmt.Errorf("Invalid query: .{ was encountered more than once.")
	} else if len(withAliases) == 2 {
		fields, aliases, err = parseQueryAliases(withAliases[1])
		if err != nil {
			return nil, nil, nil, err
		}
		parts := splitAndTrim(withAliases[0], ".")
		path = append(path, parts...)
		path = append(path, fields[0])
	} else {
		parts := splitAndTrim(query, ",")
		if len(parts) == 1 {
			parts = splitAndTrim(query, ".")
			if len(parts) > 1 {
				path, fields = parts, parts[len(parts)-1:]
			} else {
				fields = parts
			}
		} else {
			fields = parts
		}
	}
	return path, fields, aliases, err
}

func parseQueryAliases(raw string) (fields []string, aliases map[string]string, err error) {
	if string(raw[len(raw)-1]) != "}" {
		return nil, nil, fmt.Errorf("Invalid query: the alias clause must end with } and no symbols are allowed to follow it.")
	}
	aliases = map[string]string{}
	parts := strings.Split(raw[:len(raw)-1], ",")
	for _, part := range parts {
		if strings.Contains(part, ":") {
			m := strings.Split(part, ":")
			if len(m) != 2 {
				return nil, nil, fmt.Errorf("Invalid query: more than one semicolon was encountered within the alias expression.")
			}
			alias, field := m[0], m[1]
			aliases[NormalizePropertyName(field)] = alias
			fields = append(fields, field)
		} else {
			fields = append(fields, strings.Trim(part, "\t "))
		}
	}
	return
}

func parseModelByQuery(path, fields []string, model interface{}, current, next string, aliases map[string]string) interface{} {
	if slice, ok := model.([]interface{}); ok {
		var result []interface{}
		for _, el := range slice {
			nextEl := parseModelByQuery(path, fields, el, current, next, aliases)
			if nextEl != nil {
				result = append(result, nextEl)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	} else if hash, ok := model.(map[string]interface{}); ok {
		if next == "" {
			filterFields(hash, fields, aliases)
			if len(hash) == 0 {
				return nil
			}
			return hash
		} else {
			var sub interface{}
			if val, ok := hash[current]; ok {
				sub = val
			} else if val, ok := hash[NormalizePropertyName(current)]; ok {
				sub = val
			} else {
				return nil
			}
			return parseModelByQuery(path, fields, sub, next, getNextStep(path, next), aliases)
		}
	}
	return nil
}

func filterFields(m map[string]interface{}, fields []string, aliases map[string]string) {
	for k, v := range m {
		if !contains(fields, k) {
			delete(m, k)
		} else if alias, ok := inAliases(aliases, k); ok {
			delete(m, k)
			m[alias] = v
		}
	}
}

func contains(where []string, what string) bool {
	for _, s := range where {
		if NormalizePropertyName(s) == NormalizePropertyName(what) {
			return true
		}
	}
	return false
}

func inAliases(aliases map[string]string, k string) (string, bool) {
	if alias, ok := aliases[k]; ok {
		return alias, true
	}
	if alias, ok := aliases[NormalizePropertyName(k)]; ok {
		return alias, true
	}
	return "", false
}

func getNextStep(path []string, next string) string {
	for i, f := range path {
		if next == f {
			if i == len(path)-1 {
				return ""
			}
			return path[i+1]
		}
	}
	return ""
}

func splitAndTrim(s string, sym string) []string {
	parts := strings.Split(s, sym)
	for i, p := range parts {
		parts[i] = strings.Trim(p, "\t ")
	}
	return parts
}
