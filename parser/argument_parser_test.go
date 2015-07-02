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
}

var testCases = []parserTestParam{
	{input: []string{}, res: map[string]interface{}{}},
	{
		input: []string{`{"P1":"val1","P2":[1,2],"P3":true,"P4":{"P41":"val41","P42":["str"]}}`},
		res: map[string]interface{}{
			"P1": "val1",
			"P2": []interface{}{1., 2.},
			"P3": true,
			"P4": map[string]interface{}{"P41": "val41", "P42": []interface{}{"str"}},
		},
	},
	{input: []string{`{"some-key": "value"}`}, res: map[string]interface{}{"SomeKey": "value"}},
	{input: []string{`{"some-key": "value2"}`, `{"some-key": "value1"}`}, err: "Option 'SomeKey' is specified twice."},
	{input: []string{"--some-key", "value"}, res: map[string]interface{}{"SomeKey": "value"}},
	{input: []string{"--some-key", "value1", "--some-key", "value2"}, err: "Option 'SomeKey' is specified twice."},
	{input: []string{"some-key", "value"}, err: "Invalid option format, option 'some-key' should start with '--'."},
	{input: []string{"--some-key"}, res: map[string]interface{}{"SomeKey": nil}},
	{input: []string{"--some-key", "p1-key=10,p2-key=true,p3=',=!@=$ ,%^ &\"%<,.=\"'"}, res: map[string]interface{}{
		"SomeKey": map[string]interface{}{"P1Key": 10., "P2Key": true, "P3": ",=!@=$ ,%^ &\"%<,.=\""},
	}},
	{input: []string{"--some-key", "'unfinished-key"}, res: map[string]interface{}{"SomeKey": "'unfinished-key"}},
	{input: []string{"--some-key", "p1=v1,'p2'v2"}, res: map[string]interface{}{"SomeKey": "p1=v1,'p2'v2"}},
	{input: []string{"--some-key", "p1='unfinished value"}, res: map[string]interface{}{"SomeKey": "p1='unfinished value"}},
	{input: []string{"--some-key", "p1='v1'p2=v2"}, res: map[string]interface{}{"SomeKey": "p1='v1'p2=v2"}},
}

func TestArgumentParser(t *testing.T) {
	for i, testCase := range testCases {
		t.Logf("Executing %d test case.", i+1)
		res, err := parser.ParseArguments(testCase.input)
		if testCase.err != "" && err.Error() != testCase.err {
			t.Errorf("Invalid error. Expected: %s, obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result. expected %#v, obtained %#v", testCase.res, res)
		}
	}
}
