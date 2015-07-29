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
type testStructType map[string]interface{}
type testSliceType []testStructType

var testStruct = testStructType{
	"FieldString": "some string",
	"FieldInt":    1,
	"FieldBool":   true,
	"FieldStruct": testStructType{
		"FieldString": "inner string",
		"FieldInt":    1,
	},
	"FieldSlice": testSliceType{
		testStructType{
			"FieldString": "inner slice string 1",
			"FieldInt":    1,
		},
		testStructType{
			"FieldString": "inner slice string 2",
			"FieldInt":    2,
		},
	},
}
var testSlice = testSliceType{
	testStructType{
		"FieldString": "string 1",
		"FieldInt":    1,
		"FieldBool":   true,
		"FieldStruct": testStructType{
			"FieldString": "inner string 1",
			"FieldInt":    1,
		},
	},
	testStructType{
		"FieldString": "string 2",
		"FieldInt":    2,
		"FieldBool":   false,
		"FieldStruct": testStructType{
			"FieldString": "inner string 2",
			"FieldInt":    2,
		},
	},
}
var testQueryCases = []testParam{
	// Applies a query to a struct.
	{
		input: testStruct,
		query: "FieldString",
		res: testStructType{
			"FieldString": "some string",
		},
	},
	// Applies a query to a slice.
	{
		input: testSlice,
		query: "FieldString",
		res: testSliceType{
			testStructType{
				"FieldString": "string 1",
			},
			testStructType{
				"FieldString": "string 2",
			},
		},
	},
	// Applies a query with several params.
	{
		input: testStruct,
		query: "FieldString,FieldInt",
		res: testStructType{
			"FieldString": "some string",
			"FieldInt":    1,
		},
	},
	// Applies a query with non-existent params.
	{
		input: testStruct,
		query: "FieldString,FieldUnknown",
		res: testStructType{
			"FieldString": "some string",
		},
	},
	// Applies a query with all of the params being non-existent.
	{
		input: testStruct,
		query: "FieldUnknown,FieldYetUnknown",
		res:   testStructType{},
	},
	// Queries inner fields in structs.
	{
		input: testStruct,
		query: "FieldStruct.FieldString",
		res: testStructType{
			"FieldString": "inner string",
		},
	},
	// Queries inner fields in slices.
	{
		input: testSlice,
		query: "FieldStruct.FieldString",
		res: testSliceType{
			testStructType{
				"FieldString": "inner string 1",
			},
			testStructType{
				"FieldString": "inner string 2",
			},
		},
	},
	// Queries inner slices.
	{
		input: testStruct,
		query: "FieldSlice.FieldString",
		res: testSliceType{
			testStructType{
				"FieldString": "inner slice string 1",
			},
			testStructType{
				"FieldString": "inner slice string 2",
			},
		},
	},
	// Applies aliases in structs.
	{
		input: testSlice,
		query: "FieldStruct.{MyString:FieldString,MyInt:FieldInt}",
		res: testSliceType{
			testStructType{
				"MyString": "inner string 1",
				"MyInt":    1,
			},
			testStructType{
				"MyString": "inner string 2",
				"MyInt":    2,
			},
		},
	},
	// Applies aliases in slices.
	{
		input: testStruct,
		query: "FieldSlice.{MyString:FieldString,MyInt:FieldInt}",
		res: testSliceType{
			testStructType{
				"MyString": "inner slice string 1",
				"MyInt":    1,
			},
			testStructType{
				"MyString": "inner slice string 2",
				"MyInt":    2,
			},
		},
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
		if testCase.err != "" && err.Error() != testCase.err {
			t.Errorf("Invalid error. \nExpected: %s, \nobtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result. \nexpected %#v, \nobtained %#v", testCase.res, res)
		}
	}
}
