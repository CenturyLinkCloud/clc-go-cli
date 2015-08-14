package formatters

import (
	"bytes"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	table "github.com/ldmberman/tablewriter"
	"sort"
	"strings"
)

type TableFormatter struct{}

func (f *TableFormatter) FormatOutput(model interface{}) (res string, err error) {
	m, err := parser.ConvertToMapOrSlice(model)
	if err != nil {
		return "", err
	}
	return getTable(m)
}

// getTable constructs a textual table representation of the input model, which
// should be of either []interface{} or map[string]interface{} type.
//
// Every map is transformed into a table with keys in column headers and values
// in the row. A slice of maps is also transformed into a table,
// the one with multiple rows sharing the same set of column headers. Values of
// other types are printed with Go "%v" format. Slices of values of other types
// are joined with the newline.
func getTable(m interface{}) (string, error) {
	if mmap, ok := m.(map[string]interface{}); ok {
		buf := bytes.NewBuffer([]byte{})
		t := table.NewWriter(buf)
		t.SetRowLine(true)
		t.SetAlignment(table.ALIGN_CENTRE)
		t.SetAutoWrapText(false)
		values := []string{}
		keys := sortedKeys(mmap)
		for _, k := range keys {
			value, err := getTable(mmap[k])
			if err != nil {
				return "", err
			}
			values = append(values, value)
		}
		t.Append(keys)
		t.Append(values)
		t.Render()
		return string(buf.Bytes()), nil
	} else if mslice, ok := m.([]interface{}); ok {
		if len(mslice) == 0 {
			return "", nil
		}
		buf := bytes.NewBuffer([]byte{})
		t := table.NewWriter(buf)
		t.SetRowLine(true)
		t.SetAlignment(table.ALIGN_CENTRE)
		t.SetAutoWrapText(false)
		if elmap, ok := mslice[0].(map[string]interface{}); ok {
			values := make([][]string, len(mslice))
			keys := sortedKeys(elmap)
			for i, el := range mslice {
				if elmap, ok = el.(map[string]interface{}); !ok {
					return "", fmt.Errorf("The slice is not uniform - can not be represented as table.")
				}
				for _, k := range keys {
					value := ""
					var err error
					if v, ok := elmap[k]; ok {
						value, err = getTable(v)
						if err != nil {
							return "", err
						}
					}
					values[i] = append(values[i], value)
				}
			}
			t.Append(keys)
			t.AppendBulk(values)
			t.Render()
			return string(buf.Bytes()), nil
		} else {
			values := []string{}
			for _, el := range mslice {
				v, err := getTable(el)
				if err != nil {
					return "", err
				}
				values = append(values, fmt.Sprintf("%v", v))
			}
			return strings.Join(values, "\n"), nil
		}
	} else {
		return fmt.Sprintf("%v", m), nil
	}
}

func sortedKeys(m map[string]interface{}) (keys []string) {
	keys = []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
