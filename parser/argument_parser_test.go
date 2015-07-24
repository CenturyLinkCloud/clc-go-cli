package parser_test

import (
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
	"testing"
)

type parserTestParam struct {
	input []string
	err   string
	res   map[string]interface{}
	skip  bool
}

var testCases = []parserTestParam{
	{input: []string{}, res: map[string]interface{}{}},
	{
		input: []string{`{"P1":"val1","P2":[1,2],"P3":true,"P4":{"P41":"val41","P42":["str"]}}`},
		res: map[string]interface{}{
			"P1": "val1",
			"P2": []interface{}{1.0, 2.0},
			"P3": true,
			"P4": map[string]interface{}{
				"P41": "val41",
				"P42": []interface{}{"str"},
			},
		},
	},
	// Parses root values without keys from JSON.
	{input: []string{`{"some-key": "value"}`}, res: map[string]interface{}{"SomeKey": "value"}},
	// Does not allow duplicate keys.
	{input: []string{`{"some-key": "value2"}`, `{"some-key": "value1"}`}, err: "Option 'SomeKey' is specified twice."},
	// Parses --some-key=value.
	{input: []string{"--some-key", "value"}, res: map[string]interface{}{"SomeKey": "value"}},
	// Does not allow duplicate keys.
	{input: []string{"--some-key", "value1", "--some-key", "value2"}, err: "Option 'SomeKey' is specified twice."},
	{input: []string{`{"some-key": "value"}`, "--some-key", "value2"}, err: "Option 'SomeKey' is specified twice."},
	// Does not parse root values not in JSON or a=b,c=d,.. format.
	{input: []string{"value", "value2"}, err: "value is neither in JSON nor in a=b,c=d.. format."},
	// Parses top-level objects in JSON and a=b,c=d,.. format.
	{input: []string{"key-one=value1,key-two=value2", `{"key-three":"value3"}`}, res: map[string]interface{}{
		"KeyOne":   "value1",
		"KeyTwo":   "value2",
		"KeyThree": "value3",
	}},
	// Parses keys without values.
	{input: []string{"--some-key"}, res: map[string]interface{}{"SomeKey": nil}},
	// Does not parse key values from JSON or key1=value1,key2=value2,.. notation.
	{input: []string{"--some-key", `{"key": "value"}`}, res: map[string]interface{}{"SomeKey": `{"key": "value"}`}},
	{input: []string{"--some-key", "p1-key=10,p2-key=true,p3=',=!@=$ ,%^ &\"%<,.=\"'"}, res: map[string]interface{}{
		"SomeKey": "p1-key=10,p2-key=true,p3=',=!@=$ ,%^ &\"%<,.=\"'",
	}},
	// Parses --key element1 element2 element3.
	{input: []string{"--some-key", "value1", "value2", `{"value1":[1,2,3]}`, "a=b"}, res: map[string]interface{}{
		"SomeKey": []interface{}{"value1", "value2", `{"value1":[1,2,3]}`, "a=b"},
	}},
	// Parses --key element1 element2 --another-key.
	{input: []string{"--some-key", `{"key":"value"}`, "value2", "--another-key"}, res: map[string]interface{}{
		"SomeKey": []interface{}{`{"key":"value"}`, "value2"}, "AnotherKey": nil,
	}},
	// Fails with -- argument.
	{input: []string{"--"}, err: "-- is an invalid argument."},
	// Parses nested JSON objects and arrays properly.
	{input: []string{`{"k1":{"k2":{"k3":[1,2,3]}}}`}, res: map[string]interface{}{
		"K1": map[string]interface{}{
			"K2": map[string]interface{}{
				"K3": []interface{}{1.0, 2.0, 3.0},
			},
		},
	}},
	// Parses a complex case.
	{
		input: []string{`{"a":{"b":"c"}}`, "--some-long-key", "--another-key", `{"a":"b"}`, "a=b?,c=d", "--yet-another-key"},
		res: map[string]interface{}{
			"A":             map[string]interface{}{"B": "c"},
			"SomeLongKey":   nil,
			"AnotherKey":    []interface{}{`{"a":"b"}`, "a=b?,c=d"},
			"YetAnotherKey": nil,
		},
	},
	// Parses JSON arrays of objects properly.
	{
		input: []string{`{"k":[{"knested":"value"}]}`},
		res:   map[string]interface{}{"K": []interface{}{map[string]interface{}{"Knested": "value"}}},
	},
}

func TestArgumentParser(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := parser.ParseArguments(testCase.input)
		if testCase.err != "" && err.Error() != testCase.err {
			t.Errorf("Invalid error. \nExpected: %s, \nobtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result. \nexpected %#v, \nobtained %#v", testCase.res, res)
		}
	}
}
