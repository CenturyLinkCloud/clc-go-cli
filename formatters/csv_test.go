package formatters_test

import (
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/formatters"
)

func TestCsvFormatterSimple(t *testing.T) {
	i := map[string]interface{}{
		"Id":    1,
		"Name":  "Some name",
		"Price": 10.5,
	}
	expected := `Id,Name,Price
1,Some name,10.5
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}

func TestCsvFormatterNestedArrayWithPrimitives(t *testing.T) {
	i := map[string]interface{}{
		"Id":    1,
		"Name":  "Some name",
		"Price": 10.5,
		"Tags":  []string{"foo", "bar", "baz"},
	}
	expected := `Id,Name,Price,Tags
1,Some name,10.5,"foo,bar,baz"
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}

func TestCsvFormatterNestedEmptyArray(t *testing.T) {
	i := map[string]interface{}{
		"Id":   1,
		"Name": "Some name",
		"Tags": []interface{}{},
	}
	expected := `Id,Name,Tags
1,Some name,
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}

func TestCsvFormatterNestedArrayWithMaps(t *testing.T) {
	i := map[string]interface{}{
		"Id":    1,
		"Name":  "Some name",
		"Price": 10.5,
		"Tags": []interface{}{
			map[string]interface{}{
				"Name":  "tag 1",
				"Value": "value 1",
			},
			map[string]interface{}{
				"Name":  "tag 2",
				"Value": "value 2",
			},
		},
	}
	expected := `Id,Name,Price,Tags.Name,Tags.Value
1,Some name,10.5,tag 1,value 1,tag 2,value 2
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}

func TestCsvFormatterNestedMap(t *testing.T) {
	i := map[string]interface{}{
		"Id":   1,
		"Name": "Some name",
		"Xtra": map[string]interface{}{
			"Author":      "John Doe",
			"Description": "some description",
		},
	}
	expected := `Id,Name,Xtra.Author,Xtra.Description
1,Some name,John Doe,some description
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}

func TestCsvFormatterNestedMapWithSubarray(t *testing.T) {
	i := map[string]interface{}{
		"Id":   1,
		"Name": "Some name",
		"Xtra": map[string]interface{}{
			"Author": "John Doe",
			"Disks": []interface{}{
				map[string]interface{}{
					"Id":   "foo",
					"Size": 1,
				},
				map[string]interface{}{
					"Id":   "bar",
					"Size": 2,
				},
			},
		},
	}
	expected := `Id,Name,Xtra.Author,Xtra.Disks
1,Some name,John Doe,"[{""Id"":""foo"",""Size"":1},{""Id"":""bar"",""Size"":2}]"
`
	tf := formatters.CsvFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: '%s'\n obtained: '%s'", expected, got)
	}
}
