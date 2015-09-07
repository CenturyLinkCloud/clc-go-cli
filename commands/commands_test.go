package commands_test

import (
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"reflect"
	"sort"
	"testing"
)

type (
	testEntity struct {
		PropertyId   string
		PropertyName string
	}

	testCommandInput struct {
		Property1 string
		Property2 testEntity
	}
)

func TestCommandBaseArguments(t *testing.T) {
	c := &commands.CommandBase{
		Input: nil,
	}
	got := c.Arguments()
	expected := []string{}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}

	c = &commands.CommandBase{
		Input: "",
	}
	got = c.Arguments()
	expected = []string{}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}

	input := "Input"
	c = &commands.CommandBase{
		Input: &input,
	}
	got = c.Arguments()
	expected = []string{}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}

	c = &commands.CommandBase{
		Input: &testCommandInput{},
	}
	got = c.Arguments()
	expected = []string{"--property1", "--property2"}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}
}
