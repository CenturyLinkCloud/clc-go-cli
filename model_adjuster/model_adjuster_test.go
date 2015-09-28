package model_adjuster_test

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
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

type testModelIDInferable struct {
	Id   string
	Name string
}

type modelAdjusterTestCase struct {
	model interface{}
	res   interface{}
	err   string
	skip  bool
}

type connStub struct{}

func (t testModel) ApplyDefaultBehaviour() error {
	t.EssentialField = t.AuxiliaryField
	t.AuxiliaryField = ""
	return nil
}

func (t testModelIDInferable) InferID(cn base.Connection) error {
	if t.Name == "unknown" {
		return fmt.Errorf("Unknown name")
	}
	t.Id = t.Name
	return nil
}

func (t testModelIDInferable) GetNames(cn base.Connection, name string) ([]string, error) {
	return nil, nil
}

func (c connStub) ExecuteRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	return nil
}

var applyDefaultTestCases = []modelAdjusterTestCase{
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
var inferIDTestCases = []modelAdjusterTestCase{
	{
		model: testModelIDInferable{
			Name: "name",
		},
		res: testModelIDInferable{
			Id:   "name",
			Name: "name",
		},
	},
	{
		model: testModelIDInferable{
			Name: "unknown",
		},
		err: "Unknown name",
	},
}

func TestDefaultBehaviour(t *testing.T) {
	for i, testCase := range applyDefaultTestCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		err := model_adjuster.ApplyDefaultBehaviour(&testCase.model)
		res := testCase.res
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}

func TestIDInference(t *testing.T) {
	for i, testCase := range inferIDTestCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		err := model_adjuster.InferID(testCase.model, connStub{})
		res := testCase.res

		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		if (errMsg != "" || testCase.err != "") && errMsg != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, errMsg)
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
