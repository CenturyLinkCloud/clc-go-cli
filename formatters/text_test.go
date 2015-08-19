package formatters_test

import (
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"testing"
)

func TestTextFormatter(t *testing.T) {
	i := map[string]interface{}{
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
	expected := `1	Some name
TAGS	tag 1	value 1
TAGS	tag 2	value 2
VERBS	GET
VERBS	POST
VERBS	PUT`
	tf := formatters.TextFormatter{}
	got, _ := tf.FormatOutput(i)
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}
