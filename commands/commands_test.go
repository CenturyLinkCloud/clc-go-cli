package commands_test

import (
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
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

	testComposedInput struct {
		Property   string
		testEntity `argument:"composed"`
	}

	testComplexInput struct {
		Property          string
		AuxiliaryProperty string `argument:"ignore"`
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

	c = &commands.CommandBase{
		Input: &testComposedInput{},
	}
	got = c.Arguments()
	expected = []string{"--property", "--property-id", "--property-name"}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}

	c = &commands.CommandBase{
		Input: &testComplexInput{},
	}
	got = c.Arguments()
	expected = []string{"--property"}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected: %v\nGot: %v", expected, got)
	}
}

func TestLogin(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	c := commands.NewLogin(commands.CommandExcInfo{})
	conf := &config.Config{}
	opts := &options.Options{}

	// Test with no options.
	got := c.Login(opts, conf)
	expected := "Either a profile or a user and a password must be specified."
	assert(t, got, expected)

	// Try specifying a user.
	opts.User = "John@Snow"
	got = c.Login(opts, conf)
	expected = "Both --user and --password options must be specified."
	assert(t, got, expected)

	// Then provide a password.
	opts.Password = "1gr1tte"
	got = c.Login(opts, conf)
	expected = "Logged in as John@Snow."
	assert(t, got, expected)
	var err error
	conf, err = config.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	assert(t, conf.User, "John@Snow")
	assert(t, conf.Password, "1gr1tte")

	// Try to switch a profile.
	opts.User, opts.Password = "", ""
	opts.Profile = "friend"
	got = c.Login(opts, conf)
	expected = "Profile friend does not exist."
	assert(t, got, expected)
	// Oops, lets create one.
	conf.Profiles["friend"] = config.Profile{User: "Sam@Tarly", Password: "g1lly"}
	got = c.Login(opts, conf)
	expected = "Logged in as Sam@Tarly."
}

func TestSetDefaultDataCenter(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	c := commands.NewSetDefaultDC(commands.CommandExcInfo{})
	c.Input = &datacenter.SetDefault{DataCenter: "CA1"}
	if c.IsOffline() != true {
		t.Errorf("Invalid result. The command must be offline.")
	}

	got, err := c.ExecuteOffline()
	if err != nil {
		t.Error(err)
	}
	assert(t, got, "CA1 is now the default data center.")
	var conf *config.Config
	conf, err = config.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	assert(t, conf.DefaultDataCenter, "CA1")
}

func TestShowDefaultDataCenter(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	c := commands.NewShowDefaultDC(commands.CommandExcInfo{})
	if c.IsOffline() != true {
		t.Errorf("Invalid result. The command must be offline.")
	}

	got, err := c.ExecuteOffline()
	if err != nil {
		t.Error(err)
	}
	assert(t, got, "No data center is currently set as default.")

	conf := &config.Config{DefaultDataCenter: "CA1"}
	err = config.Save(conf)
	if err != nil {
		t.Error(err)
	}
	got, err = c.ExecuteOffline()
	if err != nil {
		t.Error(err)
	}
	assert(t, got, "CA1")
}

func TestUnsetDefaultDataCenter(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	conf := &config.Config{DefaultDataCenter: "CA1"}
	err := config.Save(conf)
	if err != nil {
		t.Error(err)
	}

	c := commands.NewUnsetDefaultDC(commands.CommandExcInfo{})
	if c.IsOffline() != true {
		t.Errorf("Invalid result. The command must be offline.")
	}

	var got string
	got, err = c.ExecuteOffline()
	if err != nil {
		t.Error(err)
	}
	assert(t, got, "The default data center is unset.")
	conf, err = config.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	assert(t, conf.DefaultDataCenter, "")
}

func assert(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Invalid result. Expected: %s\nGot: %s", expected, got)
	}
}
