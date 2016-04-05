package model_loader_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/model_loader"
)

var two int64 = 2

type modelLoaderTestCase struct {
	args map[string]interface{}
	res  interface{}
	err  string
	skip bool
}

type testModel struct {
	FieldString          string
	FieldInt             int64
	FieldFloat           float64
	FieldBool            bool
	FieldDateTime        time.Time
	FieldObject          testFieldObject
	FieldArray           []testFieldObject
	FieldNil             base.NilField
	testInnerObject      `argument:"composed"`
	FieldIntPtr          *int64
	FieldMapStringString map[string]string
}

type testFieldObject struct {
	FieldString          string
	FieldInnerObject     testFieldInnerObject
	FieldMapStringString map[string]string
	FieldInnerArray      []testFieldObject
}

type testFieldInnerObject struct {
	FieldString string
}

type testInnerObject struct{}

var testCases = []modelLoaderTestCase{
	// Loads simple fields.
	{
		args: map[string]interface{}{
			"FieldString": "some string",
			"FieldInt":    "4",
			"FieldFloat":  "0.0234",
			"FieldBool":   "true",
		},
		res: testModel{
			FieldString: "some string",
			FieldInt:    4,
			FieldFloat:  0.0234,
			FieldBool:   true,
		},
	},
	{
		args: map[string]interface{}{
			"FieldInt":   5.0,
			"FieldFloat": 5.5,
			"FieldBool":  true,
		},
		res: testModel{
			FieldInt:   5,
			FieldFloat: 5.5,
			FieldBool:  true,
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
	// Handles nil values.
	{
		args: map[string]interface{}{
			"FieldNil": nil,
		},
		res: testModel{FieldNil: base.NilField{Set: true}},
	},
	// Handles pointers
	{
		args: map[string]interface{}{
			"FieldIntPtr": "2",
		},
		res: testModel{
			FieldIntPtr: &two,
		},
	},
	// Handles map[string]string objects.
	{
		args: map[string]interface{}{
			"FieldMapStringString": `{"key1":"value1","key2":"value2"}`,
		},
		res: testModel{
			FieldMapStringString: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	},
	{
		args: map[string]interface{}{
			"FieldMapStringString": `key1=value1,key2=value2`,
		},
		res: testModel{
			FieldMapStringString: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	},
	{
		args: map[string]interface{}{
			"FieldObject": `{"FieldMapStringString":{"key1":"value1","key2":"value2"}}`,
		},
		res: testModel{
			FieldObject: testFieldObject{
				FieldMapStringString: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
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
	// Normalizes keys in parsed JSON.
	{
		args: map[string]interface{}{
			"FieldObject": `{"field-inner-object":{"field-string":"some string"}}`,
		},
		res: testModel{
			FieldObject: testFieldObject{
				FieldInnerObject: testFieldInnerObject{
					FieldString: "some string",
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
	// Fails to load JSON if it is of the wrong structure.
	{
		args: map[string]interface{}{
			"FieldObject": `{"FieldInnerObject":[{"FieldString":"some string"}]}`,
		},
		err: "Type mismatch: FieldInnerObject must be an object.",
	},
	{
		args: map[string]interface{}{
			"FieldArray": `[{"FieldInnerArray":{"FieldString":"some string"}}]`,
		},
		err: "Type mismatch: FieldInnerArray must be an array.",
	},
	// Fails to load string into object field if it is neither valid JSON nor k1=v1,.. notation.
	{
		args: map[string]interface{}{
			"FieldObject": `can not be parsed into object`,
		},
		err: "`can not be parsed into object` must be an object specified either in JSON or in key=value,.. format",
	},
	// Loads an array value as is if it can't be parsed as JSON.
	{
		args: map[string]interface{}{
			"FieldArray": `{"FieldString":"some string"}`,
		},
		res: testModel{
			FieldArray: []testFieldObject{{
				FieldString: "some string",
			}},
		},
	},
	// Fails to load slices into fields of simple type.
	{
		args: map[string]interface{}{
			"FieldInt": []int{1, 2, 3},
		},
		err: "Type mismatch: FieldInt value must be integer.",
	},
	{
		args: map[string]interface{}{
			"FieldString": []string{"one", "two", "three"},
		},
		err: "Type mismatch: FieldString value must be string.",
	},
	{
		args: map[string]interface{}{
			"FieldDateTime": []float64{.1, .2},
		},
		err: "Type mismatch: FieldDateTime value must be datetime in `YYYY-MM-DD hh:mm:ss` format.",
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
		err: "Unknown option or argument: `UnknownField`.",
	},
	{
		args: map[string]interface{}{
			"testInnerObject": "some value",
		},
		err: "Unknown option or argument: `testInnerObject`.",
	},
	// Fails with numbers out of range.
	{
		args: map[string]interface{}{
			"FieldInt": "99223372036854775808",
		},
		err: "Value `99223372036854775808` is too big.",
	},
	{
		args: map[string]interface{}{
			"FieldFloat": strings.Repeat("9", 310),
		},
		err: fmt.Sprintf("Value `%s` is too big.", strings.Repeat("9", 310)),
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
			"FieldFloat": "Fal",
		},
		err: "Type mismatch: FieldFloat value must be float.",
	},
	{
		args: map[string]interface{}{
			"FieldBool": "False",
		},
		err: "Type mismatch: FieldBool value must be either true or false.",
	},
	{
		args: map[string]interface{}{
			"FieldDateTime": "2012 04 05",
		},
		err: "Type mismatch: FieldDateTime value must be datetime in `YYYY-MM-DD hh:mm:ss` format.",
	},
	{
		args: map[string]interface{}{
			"FieldMapStringString": `just string`,
		},
		err: "Type mismatch: FieldMapStringString must be an object",
	},
	{
		args: map[string]interface{}{
			"FieldMapStringString": `[{"key":"value"}]`,
		},
		err: "Type mismatch: FieldMapStringString must be an object",
	},
	{
		args: map[string]interface{}{
			"FieldObject": `{"FieldMapStringString":{"key":5}}`,
		},
		err: "Type mismatch: `key` must be string",
	},
	// Does not accept any values for base.NilField's.
	{
		args: map[string]interface{}{
			"FieldNil": "",
		},
		err: "FieldNil does not accept any value.",
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
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		if (err != nil || testCase.err != "") && errMsg != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, errMsg)
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
