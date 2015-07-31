package parser_test

import (
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
	"testing"
)

type testParam struct {
	input interface{}
	query string
	err   string
	res   interface{}
	skip  bool
}

var testStruct = map[string]interface{}{
	"FieldString": "some string",
	"FieldInt":    1.,
	"FieldBool":   true,
	"FieldStruct": map[string]interface{}{
		"FieldString": "inner string",
		"FieldInt":    1.,
	},
	"FieldStruct2": map[string]interface{}{
		"FieldAnotherString": "another inner string",
	},
	"FieldSlice": []interface{}{
		map[string]interface{}{
			"FieldString": "inner slice string 1",
			"FieldInt":    1.,
		},
		map[string]interface{}{
			"FieldString": "inner slice string 2",
			"FieldInt":    2.,
		},
	},
}
var testSlice = []interface{}{
	map[string]interface{}{
		"FieldString": "string 1",
		"FieldInt":    1.,
		"FieldBool":   true,
		"FieldStruct": map[string]interface{}{
			"FieldString": "inner string 1",
			"FieldInt":    1.,
		},
	},
	map[string]interface{}{
		"FieldString": "string 2",
		"FieldInt":    2.,
		"FieldBool":   false,
		"FieldStruct": map[string]interface{}{
			"FieldString": "inner string 2",
			"FieldInt":    2.,
		},
	},
}
var testQueryCases = []testParam{
	// Applies a query to a struct.
	{
		input: testStruct,
		query: "FieldString",
		res: map[string]interface{}{
			"FieldString": "some string",
		},
	},
	// Applies a query to a slice.
	{
		input: testSlice,
		query: "FieldString",
		res: []interface{}{
			map[string]interface{}{
				"FieldString": "string 1",
			},
			map[string]interface{}{
				"FieldString": "string 2",
			},
		},
	},
	// Applies a query with several params.
	{
		input: testStruct,
		query: "FieldString,FieldInt",
		res: map[string]interface{}{
			"FieldString": "some string",
			"FieldInt":    1.,
		},
	},
	// Applies a query with non-existent params.
	{
		input: testStruct,
		query: "FieldString,FieldUnknown",
		res: map[string]interface{}{
			"FieldString": "some string",
		},
	},
	// Applies a query with all of the params being non-existent.
	{
		input: testStruct,
		query: "FieldUnknown,FieldYetUnknown",
		res:   nil,
	},
	// Queries inner fields in structs.
	{
		input: testStruct,
		query: "FieldStruct.FieldString",
		res: map[string]interface{}{
			"FieldString": "inner string",
		},
	},
	// Queries inner fields in slices.
	{
		input: testSlice,
		query: "FieldStruct.FieldString",
		res: []interface{}{
			map[string]interface{}{
				"FieldString": "inner string 1",
			},
			map[string]interface{}{
				"FieldString": "inner string 2",
			},
		},
	},
	// Queries several inner fields.
	{
		input: testSlice,
		query: "FieldStruct.{FieldString,FieldInt}",
		res: []interface{}{
			map[string]interface{}{
				"FieldString": "inner string 1",
				"FieldInt":    1.,
			},
			map[string]interface{}{
				"FieldString": "inner string 2",
				"FieldInt":    2.,
			},
		},
	},
	// Queries inner slices.
	{
		input: testStruct,
		query: "FieldSlice.FieldString",
		res: []interface{}{
			map[string]interface{}{
				"FieldString": "inner slice string 1",
			},
			map[string]interface{}{
				"FieldString": "inner slice string 2",
			},
		},
	},
	// Applies aliases in structs.
	{
		input: testSlice,
		query: "FieldStruct.{MyString:FieldString,MyInt:FieldInt}",
		res: []interface{}{
			map[string]interface{}{
				"MyString": "inner string 1",
				"MyInt":    1.,
			},
			map[string]interface{}{
				"MyString": "inner string 2",
				"MyInt":    2.,
			},
		},
	},
	// Applies aliases in slices.
	{
		input: testStruct,
		query: "FieldSlice.{MyString:FieldString,MyInt:FieldInt}",
		res: []interface{}{
			map[string]interface{}{
				"MyString": "inner slice string 1",
				"MyInt":    1.,
			},
			map[string]interface{}{
				"MyString": "inner slice string 2",
				"MyInt":    2.,
			},
		},
	},
	// Reports if queries are invalid.
	{
		input: testStruct,
		query: "FieldSlice.{MyInt:FieldInt}.SomeField",
		err:   "Invalid query: the alias close must end with } and no symbols are allowed to follow it.",
	},
	{
		input: testStruct,
		query: "FieldSlice.{MyInt:FieldInt,SomeField",
		err:   "Invalid query: the alias close must end with } and no symbols are allowed to follow it.",
	},
	{
		input: testStruct,
		query: "FieldInt.{FieldString.{",
		err:   "Invalid query: .{ was encountered more than once.",
	},
	{
		input: testStruct,
		query: "FieldSlice.{MyInt:FieldInt:}",
		err:   "Invalid query: more than one semicolon was encountered within the alias expression.",
	},
}

func TestQueryParser(t *testing.T) {
	for i, testCase := range testQueryCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := parser.ParseQuery(testCase.input, testCase.query)
		var errMsg string
		if err == nil {
			errMsg = ""
		} else {
			errMsg = err.Error()
		}
		if testCase.err != "" && errMsg != testCase.err {
			t.Errorf("Invalid error. \nExpected: %s, \nobtained %s", testCase.err, errMsg)
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result. \nexpected %#v, \nobtained %#v", testCase.res, res)
		}
	}
}
