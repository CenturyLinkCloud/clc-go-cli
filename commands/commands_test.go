package commands_test

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/auth"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
	"github.com/centurylinkcloud/clc-go-cli/state"
	"reflect"
	"sort"
	"testing"
	"time"
)

type (
	// Types for TestCommandBaseArguments.
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

	// Types for TestWait.
	footprintType1 struct {
		Links []models.LinkEntity
	}
	footprintType2 models.LinkEntity
	footprintType3 models.Status
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
		t.Fatal(err)
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
		t.Fatal(err)
	}
	assert(t, got, "CA1 is now the default data center.")
	var conf *config.Config
	conf, err = config.LoadConfig()
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	assert(t, got, "No data center is currently set as default.")

	conf := &config.Config{DefaultDataCenter: "CA1"}
	err = config.Save(conf)
	if err != nil {
		t.Fatal(err)
	}
	got, err = c.ExecuteOffline()
	if err != nil {
		t.Fatal(err)
	}
	assert(t, got, "CA1")
}

func TestUnsetDefaultDataCenter(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	conf := &config.Config{DefaultDataCenter: "CA1"}
	err := config.Save(conf)
	if err != nil {
		t.Fatal(err)
	}

	c := commands.NewUnsetDefaultDC(commands.CommandExcInfo{})
	if c.IsOffline() != true {
		t.Errorf("Invalid result. The command must be offline.")
	}

	var got string
	got, err = c.ExecuteOffline()
	if err != nil {
		t.Fatal(err)
	}
	assert(t, got, "The default data center is unset.")
	conf, err = config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert(t, conf.DefaultDataCenter, "")
}

func TestWait(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	status := commands.StatusResponse{}
	proxy.Server([]proxy.Endpoint{
		{"/authentication/login", proxy.LoginResponse},
		{"/get/status", &status},
	})
	defer proxy.CloseServer()

	cn, err := auth.AuthenticateCommand(&options.Options{User: "_", Password: "_"}, &config.Config{})
	if err != nil {
		t.Fatal(err)
	}
	commands.PING_INTERVAL = time.Duration(200)
	w := commands.NewWait(commands.CommandExcInfo{})

	// At first check an idle run.
	err = w.Execute(cn)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Nothing to wait for."
	if !reflect.DeepEqual(w.Output, &expected) {
		t.Errorf("Invalid result. Expected: %v\nGot: %v", expected, w.Output)
	}

	// Then add a footprint of the "previous" command.
	// There can be different types of footprints and we test all of them here.
	f1 := footprintType1{Links: []models.LinkEntity{
		{
			Rel:  "status",
			Href: "/get/status",
		},
	}}
	f2 := footprintType2{
		Rel:  "status",
		Href: "/get/status",
	}
	f3 := footprintType3{
		URI: "/get/status",
	}
	for _, f := range []interface{}{f1, f2, f3} {
		err := state.SaveLastResult(f)
		if err != nil {
			t.Fatal(err)
		}

		done := make(chan error)
		status.Status = "notStarted"
		go func(w *commands.Wait, cn base.Connection, done chan<- error) {
			err := w.Execute(cn)
			if err != nil {
				t.Fatal(err)
			}
			done <- nil
		}(w, cn, done)

		step := 0
	Wait:
		for {
			select {
			case <-done:
				if step < 3 {
					t.Error("Invalid result. The command finished prematurily.")
				}
				break Wait
			case <-time.After(time.Millisecond * 500):
				step += 1
				switch step {
				case 1:
					status.Status = "executing"
				case 2:
					status.Status = "resumed"
				default:
					status.Status = "succeeded"
				}
			}
		}
	}
}

func TestLoadGroups(t *testing.T) {
	ca1 := datacenter.GetRes{
		Links: []models.LinkEntity{
			{
				Rel:  "self",
				Href: "/v2/datacenters/ALIAS/CA1",
			},
			{
				Rel:  "group",
				Href: "/get/group/ca1",
			},
		},
	}
	ca2 := datacenter.GetRes{
		Links: []models.LinkEntity{
			{
				Rel:  "self",
				Href: "/v2/datacenters/ALIAS/CA2",
			},
			{
				Rel:  "group",
				Href: "/get/group/ca2",
			},
		},
	}
	allDatacenters := []datacenter.GetRes{ca1, ca2}
	group1 := group.Entity{Name: "Group 1"}
	group2 := group.Entity{Name: "Group 2"}
	proxy.Server([]proxy.Endpoint{
		{"/v2/datacenters/ALIAS/CA1", &ca1},
		{"/v2/datacenters/ALIAS/CA2", &ca2},
		{"/v2/datacenters/ALIAS", &allDatacenters},
		{"/get/group/ca1", &group1},
		{"/get/group/ca2", &group2},
		{"/authentication/login", proxy.LoginResponse},
	})
	defer proxy.CloseServer()

	cn, err := auth.AuthenticateCommand(&options.Options{User: "_", Password: "_"}, &config.Config{})
	if err != nil {
		t.Fatal(err)
	}
	c := commands.NewGroupList(commands.CommandExcInfo{})

	// Load groups for all data centers.
	c.Input.(*group.List).All.Set = true
	err = c.Execute(cn)
	if err != nil {
		t.Fatal(err)
	}
	expected := fmt.Sprintf("%v", []group.Entity{group1, group2})
	got := fmt.Sprintf("%v", c.Output.([]group.Entity))
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Invalid result.\nExpected: %s\nGot: %s", expected, got)
	}

	// Load groups for one data center.
	c.Input.(*group.List).All.Set = false
	c.Input.(*group.List).DataCenter = "CA1"
	err = c.Execute(cn)
	if err != nil {
		t.Fatal(err)
	}
	expected = fmt.Sprintf("%v", []group.Entity{group1})
	got = fmt.Sprintf("%v", c.Output.([]group.Entity))
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Invalid result.\nExpected: %s\nGot: %s", expected, got)
	}
}

func assert(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Invalid result. Expected: %s\nGot: %s", expected, got)
	}
}
