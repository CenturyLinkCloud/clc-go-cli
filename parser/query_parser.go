package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseQuery accepts a map[string]interface{} or an []interface and
// returns an object with only keys parsed from the given query string left.
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
	return parseModelByQuery(path, fields, input, current, getNextStep(path, current), aliases)
}

func ConvertToMapOrSlice(input interface{}) (interface{}, error) {
	var model interface{}

	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &model)
	removeLinks(model)
	return model, err
}

func removeLinks(input interface{}) {
	switch input.(type) {
	case map[string]interface{}:
		m := input.(map[string]interface{})
		delete(m, "Links")
		for _, value := range m {
			removeLinks(value)
		}
	case []interface{}:
		array := input.([]interface{})
		for _, child := range array {
			removeLinks(child)
		}
	}
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
		for _, p := range parts {
			if strings.Contains(p, ",") {
				return nil, nil, nil, fmt.Errorf("If nested fields are queried, multiple fields can only be specified in a .{...} clause")
			}
		}
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
			for _, p := range parts {
				if strings.Contains(p, ".") {
					return nil, nil, nil, fmt.Errorf("If nested fields are queried, multiple fields can only be specified in a .{...} clause")
				}
			}
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
			if strings.Contains(field, ".") {
				return nil, nil, fmt.Errorf("%s: you can not query inner objects in the .{...} clause", field)
			}
			aliases[NormalizePropertyName(field)] = alias
			fields = append(fields, field)
		} else {
			if strings.Contains(part, ".") {
				return nil, nil, fmt.Errorf("%s: you can not query inner objects in the .{...} clause", part)
			}
			fields = append(fields, strings.Trim(part, "\t "))
		}
	}
	return
}

func parseModelByQuery(path, fields []string, model interface{}, current, next string, aliases map[string]string) (interface{}, error) {
	if slice, ok := model.([]interface{}); ok {
		var result []interface{}
		for _, el := range slice {
			nextEl, err := parseModelByQuery(path, fields, el, current, next, aliases)
			if err != nil {
				return nil, err
			}
			result = append(result, nextEl)
		}
		return result, nil
	} else if hash, ok := model.(map[string]interface{}); ok {
		if next == "" {
			for _, f := range fields {
				if _, ok := hash[NormalizePropertyName(f)]; !ok {
					return nil, noSuchPath(path, f)
				}
			}
			return filterFields(hash, fields, aliases), nil
		} else {
			var sub interface{}
			if val, ok := hash[current]; ok {
				sub = val
			} else if val, ok := hash[NormalizePropertyName(current)]; ok {
				sub = val
			} else {
				return nil, noSuchPath(path, current)
			}
			return parseModelByQuery(path, fields, sub, next, getNextStep(path, next), aliases)
		}
	}
	return nil, noSuchPath(path, current)
}

func filterFields(m map[string]interface{}, fields []string, aliases map[string]string) map[string]interface{} {
	r := map[string]interface{}{}
	for k, v := range m {
		if alias, ok := inAliases(aliases, k); ok {
			r[alias] = v
		} else if contains(fields, k) {
			r[k] = v
		}
	}
	return r
}

func noSuchPath(path []string, current string) error {
	p := ""
	finished := false
	for _, s := range path {
		if s == current {
			p += current
			finished = true
			break
		}
		p = p + s + "."
	}
	if !finished {
		p += current
	}

	return fmt.Errorf("%s: there is no such field in the result", p)
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
