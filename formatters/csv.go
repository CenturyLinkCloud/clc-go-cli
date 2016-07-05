package formatters

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

type CsvFormatter struct{}

func (f *CsvFormatter) FormatOutput(model interface{}) (res string, err error) {
	buffer := bytes.NewBufferString("")
	csvWriter := csv.NewWriter(buffer)

	if modelAsMap, ok := model.(map[string]interface{}); ok {
		header := f.extractHeader(modelAsMap)
		csvWriter.Write(header)

		row := f.mapToTableRow(modelAsMap)
		csvWriter.Write(row)
	} else if models, ok := model.([]interface{}); ok {
		if len(models) > 0 {
			first := models[0].(map[string]interface{})
			header := f.extractHeader(first)
			csvWriter.Write(header)

			for _, model := range models {
				modelAsMap := model.(map[string]interface{})
				row := f.mapToTableRow(modelAsMap)
				csvWriter.Write(row)
			}
		}
	} else {
		return "", fmt.Errorf("Don't know how to convert %v into CSV", model)
	}

	csvWriter.Flush()

	return buffer.String(), nil
}

func (f *CsvFormatter) mapToTableRow(modelAsMap map[string]interface{}) []string {
	row := []string{}
	for _, key := range sortedKeys(modelAsMap) {
		value := modelAsMap[key]
		row = append(row, f.formatValue(value, 1)...)
	}
	return row
}

func (f *CsvFormatter) extractHeader(model map[string]interface{}) []string {
	header := []string{}
	for _, key := range sortedKeys(model) {
		value := model[key]

		if valueMap, ok := value.(map[string]interface{}); ok {
			for _, nestedKey := range sortedKeys(valueMap) {
				header = append(header, fmt.Sprintf("%s.%s", key, nestedKey))
			}
		} else if valueSlice, ok := value.([]interface{}); ok {
			if len(valueSlice) > 0 {
				if firstItemMap, ok := valueSlice[0].(map[string]interface{}); ok {
					// get first item of slice
					subheaders := f.extractHeader(firstItemMap)
					for _, subheader := range subheaders {
						header = append(header, fmt.Sprintf("%s.%s", key, subheader))
					}
				} else {
					header = append(header, key)
				}
			} else {
				header = append(header, key)
			}
		} else {
			header = append(header, key)
		}
	}

	return header
}

func (f *CsvFormatter) formatValue(value interface{}, depth int) []string {
	values := []string{}
	nextDepth := depth + 1

	if valueSlice, ok := value.([]interface{}); ok {
		// value is a slice of maps
		if len(valueSlice) == 0 {
			values = append(values, "")
		} else if depth > 1 {
			b, err := json.Marshal(valueSlice)
			if err != nil {
				panic(err)
			}
			values = append(values, string(b))
		} else {
			for _, item := range valueSlice {
				values = append(values, f.formatValue(item, nextDepth)...)
			}
		}
	} else if valueSlice, ok := value.([]string); ok {
		// value is a slice of strigs
		values = append(values, strings.Join(valueSlice, ","))
	} else if valueMap, ok := value.(map[string]interface{}); ok {
		// value in nested map
		if depth > 2 {
			b, err := json.Marshal(valueMap)
			if err != nil {
				panic(err)
			}
			values = append(values, string(b))
		} else {
			for _, k := range sortedKeys(valueMap) {
				values = append(values, f.formatValue(valueMap[k], nextDepth)...)
			}
		}
	} else {
		values = append(values, fmt.Sprintf("%v", value))
	}

	return values
}
