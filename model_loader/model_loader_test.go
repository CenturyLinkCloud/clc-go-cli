package model_loader_test

import (
	"github.com/centurylinkcloud/clc-go-cli/model_loader"
	"reflect"
	"testing"
	"time"
)

type modelLoaderTestCase struct {
	args map[string]interface{}
	res  interface{}
	err  string
	skip bool
}

type testModel struct {
	FieldString   string
	FieldInt      int64
	FieldFloat    float64
	FieldBool     bool
	FieldDateTime time.Time
	FieldObject   testFieldObject
	FieldArray    []testFieldObject
}

type testFieldObject struct {
	FieldString      string
	FieldInnerObject testFieldInnerObject
}

type testFieldInnerObject struct {
	FieldString string
}

var testCases = []modelLoaderTestCase{
	// Loads string field.
	{
		args: map[string]interface{}{
			"FieldString": "some string",
		},
		res: testModel{
			FieldString: "some string",
		},
	},
	// Loads time.Time field.
	{
		args: map[string]interface{}{
			"FieldDateTime": "2012-02-13 15:40:00",
		},
		res: testModel{
			FieldDateTime: time.Date(2012, time.February, 13, 15, 40, 0, 0, time.UTC),
		},
	},
	// Parses JSON and loads it into object field.
	{
		args: map[string]interface{}{
			"FieldObject": `{"FieldString": "some string","FieldInnerObject":{"FieldString":"another string"}}`,
		},
		res: testModel{
			FieldObject: testFieldObject{
				FieldString: "some string",
				FieldInnerObject: testFieldInnerObject{
					FieldString: "another string",
				},
			},
		},
	},
	{
		args: map[string]interface{}{
			"FieldArray": `[{"FieldString": "string 1"},{"FieldString": "string 2"}]`,
		},
		res: testModel{
			FieldArray: []testFieldObject{
				testFieldObject{FieldString: "string 1"},
				testFieldObject{FieldString: "string 2"},
			},
		},
	},
	// Parses k1=v1,k2=v2,.. notation and loads it into object field.
	{
		args: map[string]interface{}{
			"FieldObject": `FieldString=some string`,
		},
		res: testModel{
			FieldObject: testFieldObject{
				FieldString: "some string",
			},
		},
	},
	// Fails to load string into object field if it is neither JSON nor k1=v1,.. notation.
	{
		args: map[string]interface{}{
			"FieldObject": `can not be parsed into object`,
		},
		err: "`can not be parsed into object` is neither in JSON nor in key=value,.. format.",
	},
	// Loads JSON into string field as string.
	{
		args: map[string]interface{}{
			"FieldString": `{"a":"b"}`,
		},
		res: testModel{
			FieldString: `{"a":"b"}`,
		},
	},
	// Loads k1=v1,k2=v2.. notation into string field as string.
	{
		args: map[string]interface{}{
			"FieldString": `a=b,c=d`,
		},
		res: testModel{
			FieldString: `a=b,c=d`,
		},
	},
	// Fails with unknown fields.
	{
		args: map[string]interface{}{
			"UnknownField": "some value",
		},
		err: "Field `UnknownField` does not exist.",
	},
	// Fails with different type mismatches.
	{
		args: map[string]interface{}{
			"FieldInt": "string",
		},
		err: "Type mismatch: FieldInt value must be integer.",
	},
	{
		args: map[string]interface{}{
			"FieldBool": "Fal",
		},
		err: "Type mismatch: FieldBool value must be either true or false.",
	},
	{
		args: map[string]interface{}{
			"FieldDateTime": "2012 04 05",
		},
		err: "Type mismatch: FieldDateTime value must be datetime in `YYYY-MM-DD hh:mm:ss` format.",
	},
}

func TestModelLoader(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res := testModel{}
		err := model_loader.LoadModel(testCase.args, &res)
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
