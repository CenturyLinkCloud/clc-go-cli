package formatters

import (
	"fmt"
	"sort"
	"strings"
)

type TextFormatter struct{}

func (f *TextFormatter) FormatOutput(model interface{}) (res string, err error) {
	text := []string{}
	printText(model, "", &text)
	return strings.Join(text, "\n"), nil
}

func printText(m interface{}, header string, text *[]string) {
	if mmap, ok := m.(map[string]interface{}); ok {
		flatValues, nestedKeys, nestedValues := splitMap(mmap)
		if len(flatValues) > 0 {
			*text = append(*text, textRow(header, flatValues...))
		}
		for i, k := range nestedKeys {
			printText(nestedValues[i], k, text)
		}
	} else if mslice, ok := m.([]interface{}); ok {
		for _, el := range mslice {
			printText(el, header, text)
		}
	} else {
		*text = append(*text, textRow(header, m))
	}
}

func splitMap(m map[string]interface{}) ([]interface{}, []string, []interface{}) {
	flatValues, nestedKeys, nestedValues := []interface{}{}, []string{}, []interface{}{}
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := m[k]
		_, isMap := v.(map[string]interface{})
		_, isSlice := v.([]interface{})
		if isMap || isSlice {
			nestedKeys = append(nestedKeys, k)
			nestedValues = append(nestedValues, v)
		} else {
			flatValues = append(flatValues, v)
		}
	}
	return flatValues, nestedKeys, nestedValues
}

func textRow(header string, values ...interface{}) string {
	res := ""
	if header != "" {
		res += fmt.Sprintf("%s\t", strings.ToUpper(header))
	}
	for i, v := range values {
		if v != nil {
			res += fmt.Sprintf("%v", v)
		}
		if i != len(values)-1 {
			res += "\t"
		}
	}
	return res
}
