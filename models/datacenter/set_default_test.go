package datacenter_test

import (
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"reflect"
	"testing"
)

type (
	withoutDataCenter struct{}
	withDataCenter    struct {
		DataCenter string
	}
)

func TestApplyDefaultForNil(t *testing.T) {
	// Simply check that it does not panic.
	datacenter.ApplyDefault(nil, &config.Config{})
}

func TestApplyDefaultForSetDefault(t *testing.T) {
	// Simply check that it does not panic.
	datacenter.ApplyDefault(&datacenter.SetDefault{}, &config.Config{})
}

func TestApplyDefaultForEmptyConfig(t *testing.T) {
	var m interface{}

	m = &withoutDataCenter{}
	// Simply check that it does not panic.
	datacenter.ApplyDefault(m, &config.Config{})

	m = &withDataCenter{DataCenter: "WA1"}
	datacenter.ApplyDefault(m, &config.Config{})
	got := *m.(*withDataCenter)
	expected := withDataCenter{DataCenter: "WA1"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot:%v", got, expected)
	}
}

func TestApplyDefaultForSetConfig(t *testing.T) {
	var m interface{}
	conf := &config.Config{DefaultDataCenter: "GB3"}

	m = &withoutDataCenter{}
	// Simply check that it does not panic.
	datacenter.ApplyDefault(m, conf)

	// Check that the default one does not override the specific one.
	m = &withDataCenter{DataCenter: "WA1"}
	datacenter.ApplyDefault(m, conf)
	got := *m.(*withDataCenter)
	expected := withDataCenter{DataCenter: "WA1"}
	if got != expected {
		t.Errorf("Invalid result.\nExpected: %v\nGot:%v", got, expected)
	}

	m = &withDataCenter{}
	datacenter.ApplyDefault(m, conf)
	got = *m.(*withDataCenter)
	expected = withDataCenter{DataCenter: "GB3"}
	if got != expected {
		t.Errorf("Invalid result.\nExpected: %v\nGot:%v", got, expected)
	}
}
