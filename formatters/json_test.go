package formatters_test

import (
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	m := map[string]interface{}{
		"Id":   1.,
		"Name": "Some name",
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
		"Verbs": []interface{}{
			"GET",
			"POST",
			"PUT",
		},
	}
	expected := `{
    "Id": 1,
    "Name": "Some name",
    "Tags": [
        {
            "Name": "tag 1",
            "Value": "value 1"
        },
        {
            "Name": "tag 2",
            "Value": "value 2"
        }
    ],
    "Verbs": [
        "GET",
        "POST",
        "PUT"
    ]
}
`
	tf := formatters.JsonFormatter{}
	got, err := tf.FormatOutput(m)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}

	slice := []interface{}{
		"GET",
		"POST",
		"DELETE",
	}
	expected = `[
    "GET",
    "POST",
    "DELETE"
]
`
	got, err = tf.FormatOutput(slice)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}

	str := "GET"
	expected = `"GET"
`
	got, err = tf.FormatOutput(str)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}

	expected = `{}
`
	got, err = tf.FormatOutput(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}
