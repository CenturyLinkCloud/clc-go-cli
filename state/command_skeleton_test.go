package state_test

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"reflect"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
	"github.com/centurylinkcloud/clc-go-cli/state"
)

type Model struct {
	A string
	B string
}

type ModelWithNil struct {
	X string
	N base.NilField
}

func TestArgumentsToJSON(t *testing.T) {
	// Test with no data.
	empty := map[string]interface{}{}
	data, err := state.ArgumentsToJSON(empty, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal([]byte(data), &empty)
	if err != nil {
		t.Errorf("Unmarshalling failed: %s", err.Error())
	}
	if len(empty) != 0 {
		t.Errorf("Invalid result\nExpected an empty map\nGot %v", empty)
	}

	args := map[string]interface{}{
		"B": "B value overriden",
		"C": "C value",
	}
	m := Model{A: "A value", B: "B value"}
	data, err = state.ArgumentsToJSON(args, &m)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]interface{}{}
	err = json.Unmarshal([]byte(data), &got)
	if err != nil {
		t.Errorf("Unmarshalling failed: %s", err.Error())
	}
	expected := map[string]interface{}{
		"A": "A value",
		"B": "B value overriden",
		"C": "C value",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result\nExpected %v\nGot %v", expected, got)
	}

	// Test a model with a nil field
	args = map[string]interface{}{
		"X": "X value",
	}
	mWithNil := ModelWithNil{}
	data, err = state.ArgumentsToJSON(args, &mWithNil)
	if err != nil {
		t.Fatal(err)
	}
	got = map[string]interface{}{}
	err = json.Unmarshal([]byte(data), &got)
	if err != nil {
		t.Errorf("Unmarshalling failed: %s", err.Error())
	}
	expected = map[string]interface{}{
		"X": "X value",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result\nExpected %v\nGot %v", expected, got)
	}

	args = map[string]interface{}{
		"X": "X value",
		"N": nil,
	}
	mWithNil = ModelWithNil{}
	data, err = state.ArgumentsToJSON(args, &mWithNil)
	if err != nil {
		t.Fatal(err)
	}
	got = map[string]interface{}{}
	err = json.Unmarshal([]byte(data), &got)
	if err != nil {
		t.Errorf("Unmarshalling failed: %s", err.Error())
	}
	expected = map[string]interface{}{
		"X": "X value",
		"N": nil,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result\nExpected %v\nGot %v", expected, got)
	}
}

func TestArgumentsFromJSON(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	m := Model{A: "A value", B: "B value"}
	p, err := config.GetPath()
	if err != nil {
		t.Fatal(err)
	}
	var bytes []byte
	bytes, err = json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(path.Join(p, "test"), bytes, 0666)
	if err != nil {
		t.Fatal(err)
	}

	got := map[string]interface{}{}
	got, err = state.ArgumentsFromJSON(path.Join(p, "test"))
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]interface{}{
		"A": "A value",
		"B": "B value",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result\nExpected %v\nGot %v", expected, got)
	}
}
