package model_adjuster_test

import (
	"github.com/centurylinkcloud/clc-go-cli/model_adjuster"
	"reflect"
	"testing"
)

type testModel struct {
	AuxiliaryField string
	EssentialField string
}

type testModelNotAdjustable struct {
	Field string
}

type modelAdjusterTestCase struct {
	model interface{}
	res   interface{}
	err   string
	skip  bool
}

func (t testModel) ApplyDefaultBehaviour() error {
	t.EssentialField = t.AuxiliaryField
	t.AuxiliaryField = ""
	return nil
}

var testCases = []modelAdjusterTestCase{
	{
		model: testModel{
			AuxiliaryField: "some string",
			EssentialField: "",
		},
		res: testModel{
			AuxiliaryField: "",
			EssentialField: "some string",
		},
	},
	{
		model: testModelNotAdjustable{
			Field: "some string",
		},
		res: testModelNotAdjustable{
			Field: "some string",
		},
	},
}

func TestModelAdjuster(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		err := model_adjuster.ApplyDefaultBehaviour(testCase.model)
		res := testCase.res
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
