package state_test

import (
	"encoding/json"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/proxy"
	"github.com/centurylinkcloud/clc-go-cli/state"
)

type Result struct {
	Status string
}

func TestLastResult(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	r := Result{Status: "http://.../status"}
	err := state.SaveLastResult(r)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := state.LoadLastResult()
	if err != nil {
		t.Fatal(err)
	}
	var got Result
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Errorf("Result unmarshalling failed: %s", err.Error())
	}
	if got != r {
		t.Errorf("Invalid result\nExpected %v\nGot %v", r, got)
	}

	// Repeat to check that we only store one result at a time.
	r.Status = "http://"
	err = state.SaveLastResult(r)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err = state.LoadLastResult()
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(bytes, &got)
	if err != nil {
		t.Errorf("Result unmarshalling failed: %s", err.Error())
	}
	if got != r {
		t.Errorf("Invalid result\nExpected %v\nGot %v", r, got)
	}
}
