package model_validator_test

import (
	"fmt"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/model_validator"
)

type testModel struct {
	FieldNotRequired    string
	FieldRequired       string `valid:"required"`
	FieldRequiredCustom string
	Enumerable          string `oneOf:"v1,v2,v3"`
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
			Enumerable:          "v1",
		},
		err: "",
	},
	{
		model: testModel{
			FieldRequired: "",
		},
		err: "The field-required field must be set and non-empty",
	},
	{
		model: testModel{
			FieldRequired:       "some string",
			FieldRequiredCustom: "",
			Enumerable:          "v3",
		},
		err: "Field required.",
	},
	{
		model: testModelWithoutCustomLogic{
			FieldRequired: "some string",
		},
		err: "",
	},
	{
		model: testModel{
			Enumerable:          "v2",
			FieldRequired:       "checked",
			FieldRequiredCustom: "checked",
		},
		err: "",
	},
	// Test that it is case-insensitive.
	{
		model: testModel{
			Enumerable:          "V2",
			FieldRequired:       "checked",
			FieldRequiredCustom: "checked",
		},
		err: "",
	},
	{
		model: testModel{
			Enumerable:          "v100",
			FieldRequired:       "checked",
			FieldRequiredCustom: "checked",
		},
		err: "Enumerable value must be one of v1, v2, v3.",
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
