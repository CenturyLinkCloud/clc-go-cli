package parser_test

import (
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"reflect"
	"testing"
)

type filterTestParam struct {
	input  interface{}
	filter string
	err    string
	res    interface{}
	skip   bool
}

var filterTestStruct = map[string]interface{}{
	"Str":  "some string",
	"Num":  1.,
	"Bool": true,
	"Struct": map[string]interface{}{
		"Str": "inner string",
	},
	"Slice": []interface{}{"element"},
}
var filterTestSlice = []interface{}{
	map[string]interface{}{
		"Str":  "some string",
		"Num":  1.,
		"Bool": true,
	},
	map[string]interface{}{
		"Str":  "another string",
		"Num":  2.,
		"Bool": false,
	},
}
var filterTestCases = []filterTestParam{
	// Applies valid filters to structs.
	{
		input:  filterTestStruct,
		filter: `str="some string"`,
		res:    filterTestStruct,
	},
	{
		input:  filterTestStruct,
		filter: `str="no such string"`,
		res:    nil,
	},
	{
		input:  filterTestStruct,
		filter: `num=1`,
		res:    filterTestStruct,
	},
	{
		input:  filterTestStruct,
		filter: `num<3`,
		res:    filterTestStruct,
	},
	{
		input:  filterTestStruct,
		filter: `num>="7.2"`,
		res:    nil,
	},
	{
		input:  filterTestStruct,
		filter: `bool=true`,
		res:    filterTestStruct,
	},
	{
		input:  filterTestStruct,
		filter: `bool="true"`,
		res:    filterTestStruct,
	},
	{
		input:  filterTestStruct,
		filter: `bool="false"`,
		res:    nil,
	},
	// Applies valid filters to slices.
	{
		input:  filterTestSlice,
		filter: `str~=ome`,
		res:    []interface{}{filterTestSlice[0]},
	},
	{
		input:  filterTestSlice,
		filter: `str^=another`,
		res:    []interface{}{filterTestSlice[1]},
	},
	{
		input:  filterTestSlice,
		filter: `str$=string`,
		res:    filterTestSlice,
	},
	{
		input:  filterTestSlice,
		filter: `str>z`,
		res:    []interface{}{},
	},
	{
		input:  filterTestSlice,
		filter: `num>=2`,
		res:    []interface{}{filterTestSlice[1]},
	},
	{
		input:  filterTestSlice,
		filter: `num<="7"`,
		res:    filterTestSlice,
	},
	// Complains about the invalid filter.
	{
		input:  filterTestStruct,
		filter: `str`,
		err:    "Invalid filter format.",
	},
	{
		input:  filterTestStruct,
		filter: `str,num,bool`,
		err:    "Invalid filter format.",
	},
	{
		input:  filterTestStruct,
		filter: `str=`,
		err:    "Invalid filter format.",
	},
	{
		input:  filterTestStruct,
		filter: `"str~=some,num=2"`,
		err:    "Invalid filter format.",
	},
	// Complains about incompatibilies between the certain operations and certain field types.
	{
		input:  filterTestStruct,
		filter: `num~=1`,
		err:    "Operations ~=, ^= and $= can only be used with strings.",
	},
	{
		input:  filterTestStruct,
		filter: `num^=1`,
		err:    "Operations ~=, ^= and $= can only be used with strings.",
	},
	{
		input:  filterTestStruct,
		filter: `num$=1`,
		err:    "Operations ~=, ^= and $= can only be used with strings.",
	},
	{
		input:  filterTestSlice,
		filter: `bool<=true`,
		err:    "Operations <,>,<= and >= can not be used with booleans.",
	},
	{
		input:  filterTestSlice,
		filter: `bool~=true`,
		err:    "Operations ~=, ^= and $= can only be used with strings.",
	},
	{
		input:  filterTestStruct,
		filter: `struct={}`,
		err:    "Structs are not supported in filters.",
	},
	{
		input:  filterTestStruct,
		filter: `slice~={"key":"value"}`,
		err:    "Slices are not supported in filters.",
	},
	// Applies valid filters with multiple conditions.
	{
		input:  filterTestSlice,
		filter: `str~=string,num<5.34`,
		res:    filterTestSlice,
	},
	{
		input:  filterTestSlice,
		filter: `str~=str,num>1.85`,
		res:    []interface{}{filterTestSlice[1]},
	},
	{
		input:  filterTestSlice,
		filter: `str$=g,bool=false`,
		res:    []interface{}{filterTestSlice[1]},
	},
	{
		input:  filterTestSlice,
		filter: `str^=ano,bool=true`,
		res:    nil,
	},
	// Complains about the unknown fields.
	{
		input:  filterTestSlice,
		filter: `str$=ing,unknown=5`,
		err:    "Unknown: there is no such field in result.",
	},
	// Complains about the inappropriate values for the fields of the certain types.
	{
		input:  filterTestStruct,
		filter: `num=abc`,
		err:    "Invalid value for the number: abc.",
	},
	{
		input:  filterTestStruct,
		filter: `num=""`,
		err:    "num: non-empty value required.",
	},
	{
		input:  filterTestStruct,
		filter: `bool=0`,
		err:    "Invalid value for the boolean: 0.",
	},
	{
		input:  filterTestStruct,
		filter: `bool=""`,
		err:    "bool: non-empty value required.",
	},
}

func TestFilterParser(t *testing.T) {
	for i, testCase := range filterTestCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := parser.ParseFilter(testCase.input, testCase.filter)
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
