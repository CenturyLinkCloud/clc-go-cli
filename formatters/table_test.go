package formatters_test

import (
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"testing"
)

func TestTableFormatter(t *testing.T) {
	var m interface{}
	tf := &formatters.TableFormatter{}

	m = map[string]interface{}{
		"Field 1": "Value 1",
		"Field 2": "Value 2",
	}
	formatters.SetTerminalWidthFn(func() uint {
		return uint(80)
	})
	got, err := tf.FormatOutput(m)
	if err != nil {
		t.Error(err)
	}
	expected := `+---------+---------+
| Field 1 | Field 2 |
+---------+---------+
| Value 1 | Value 2 |
+---------+---------+
`
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}

	// Now let's try to narrow our virtual terminal a bit.
	formatters.SetTerminalWidthFn(func() uint {
		return uint(20)
	})
	got, err = tf.FormatOutput(m)
	if err != nil {
		t.Error(err)
	}
	expected = `+---------+---------+
| Field 1 | Value 1 |
+---------+---------+
| Field 2 | Value 2 |
+---------+---------+
`
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}

	// Back to the wide terminal and let's test slices and nested maps.
	formatters.SetTerminalWidthFn(func() uint {
		return uint(80)
	})
	m = []interface{}{
		map[string]interface{}{
			"Field 1": "Value 1",
			"Field 2": "Value 2",
			"Field 3": map[string]interface{}{
				"Inner Field 1": nil,
				"Inner Field 2": 4,
			},
		},
		map[string]interface{}{
			"Field 1": []interface{}{"PUT", "GET", "POST"},
			"Field 2": false,
		},
	}
	got, err = tf.FormatOutput(m)
	if err != nil {
		t.Error(err)
	}
	expected = `+---------------------------------------+
| +---------+---------+                 |
| | Field 1 | Field 2 |                 |
| +---------+---------+                 |
| | Value 1 | Value 2 |                 |
| +---------+---------+                 |
+---------------------------------------+
| +-----------------------------------+ |
| | Field 3                           | |
| +-----------------------------------+ |
| | +---------------+---------------+ | |
| | | Inner Field 1 | Inner Field 2 | | |
| | +---------------+---------------+ | |
| | |               |       4       | | |
| | +---------------+---------------+ | |
| +-----------------------------------+ |
+---------------------------------------+

+--------------------------------+
| +---------+                    |
| | Field 2 |                    |
| +---------+                    |
| |  false  |                    |
| +---------+                    |
+--------------------------------+
| +--------------+               |
| | Field 1      |               |
| +--------------+               |
| | PUT          |               |
| | GET          |               |
| | POST         |               |
| +--------------+               |
+--------------------------------+
`
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}
