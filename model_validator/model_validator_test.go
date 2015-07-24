package model_validator_test

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"testing"
)

type testModel struct {
	FieldNotRequired    string
	FieldRequired       string `valid:"required"`
	FieldRequiredCustom string
}

func (t testModel) Validate() error {
	if t.FieldRequiredCustom == "" {
		return fmt.Errorf("Field required.")
	}
	return nil
}

type testModelWithoutCustomLogic struct {
	FieldNotRequired string
	FieldRequired    string `valid:"required"`
}

type modelValidatorTestCase struct {
	model interface{}
	err   string
	skip  bool
}

var testCases = []modelValidatorTestCase{
	{
		model: testModel{
			FieldRequired:       "some string",
			FieldRequiredCustom: "another string",
		},
		err: "",
	},
	{
		model: testModel{
			FieldRequired: "",
		},
		err: "FieldRequired: non zero value required",
	},
	{
		model: testModel{
			FieldRequired:       "some string",
			FieldRequiredCustom: "",
		},
		err: "Field required.",
	},
	{
		model: testModelWithoutCustomLogic{
			FieldRequired: "some string",
		},
		err: "",
	},
}

func TestModelValidator(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		err := model_validator.ValidateModel(testCase.model)
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
	}
}
