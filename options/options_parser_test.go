package options_test

import (
	"github.com/centurylinkcloud/clc-go-cli/options"
	"reflect"
	"testing"
)

type Case struct {
	input map[string]interface{}
	res   *options.Options
	err   string
	skip  bool
}

var extractFromCases = []Case{
	// Parses valid data.
	{
		input: map[string]interface{}{
			"Help":         nil,
			"User":         "John",
			"Password":     "Snow",
			"Profile":      "default",
			"AccountAlias": "THEWALL",
			"Trace":        nil,
			"Output":       "text",
			"Query":        "location",
			"Filter":       "name=Ygritte",
		},
		res: &options.Options{
			Help:         true,
			User:         "John",
			Password:     "Snow",
			Profile:      "default",
			AccountAlias: "THEWALL",
			Trace:        true,
			Output:       "text",
			Query:        "location",
			Filter:       "name=Ygritte",
		},
	},
	// Complains about the invalid data.
	{
		input: map[string]interface{}{
			"Help": "me",
		},
		err: "help option must not have a value",
	},
	{
		input: map[string]interface{}{
			"User": 1,
		},
		err: "User must be string.",
	},
	{
		input: map[string]interface{}{
			"Password": 1,
		},
		err: "Password must be string",
	},
	{
		input: map[string]interface{}{
			"Profile": 1,
		},
		err: "Profile must be string.",
	},
	{
		input: map[string]interface{}{
			"AccountAlias": []string{"THE", "WALL"},
		},
		err: "Account alias must be string.",
	},
	{
		input: map[string]interface{}{
			"Trace": "it, please",
		},
		err: "trace option must not have a value",
	},
	{
		input: map[string]interface{}{
			"Output": []string{""},
		},
		err: "The --output value must be a string.",
	},
	{
		input: map[string]interface{}{
			"Query": []string{""},
		},
		err: "Query must be string.",
	},
	{
		input: map[string]interface{}{
			"Filter": []string{""},
		},
		err: "Filter must be string.",
	},
}

func TestExtractFrom(t *testing.T) {
	for i, testCase := range extractFromCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := options.ExtractFrom(testCase.input)
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}

func TestAreToBeSaved(t *testing.T) {
	expected := false
	got, err := options.AreToBeSaved(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}

	expected = true
	got, err = options.AreToBeSaved(map[string]interface{}{"GenerateCliSkeleton": nil})
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}

	expected = false
	noValueError := "generate-cli-skeleton option must not have a value"
	got, err = options.AreToBeSaved(map[string]interface{}{"GenerateCliSkeleton": "value"})
	if err == nil || err.Error() != noValueError {
		t.Errorf("Invalid error.\n Expected: %s,\n obtained: %v", noValueError, err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}
}

func TestAreToBeTakenFromFile(t *testing.T) {
	expected := false
	got, file, err := options.AreToBeTakenFromFile(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}
	if file != "" {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", "", file)
	}

	expected = false
	noOtherOptionsError := "No other options are allowed to be with the from-file."
	got, file, err = options.AreToBeTakenFromFile(map[string]interface{}{"FromFile": "/", "Name": "server"})
	if err == nil || err.Error() != noOtherOptionsError {
		t.Errorf("Invalid error.\n Expected: %s,\n obtained: %v", noOtherOptionsError, err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}
	if file != "" {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", "", file)
	}

	expected = false
	invalidFileError := "Invalid file path: 1."
	got, file, err = options.AreToBeTakenFromFile(map[string]interface{}{"FromFile": 1})
	if err == nil || err.Error() != invalidFileError {
		t.Errorf("Invalid error.\n Expected: %s,\n obtained: %v", invalidFileError, err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}
	if file != "" {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", "", file)
	}

	expected = true
	p := "/home/user/file"
	got, file, err = options.AreToBeTakenFromFile(map[string]interface{}{"FromFile": p})
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("Invalid result.\n Expected: %v,\n obtained: %v", expected, got)
	}
	if file != p {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", p, file)
	}
}
