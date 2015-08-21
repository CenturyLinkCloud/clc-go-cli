package autocomplete_test

import (
	"github.com/centurylinkcloud/clc-go-cli/autocomplete"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestResourceAutocomplete(t *testing.T) {
	resources := command_loader.GetResources()
	sort.Strings(resources)

	args := []string{""}
	opts := strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}

	args = []string{"serve"}
	opts = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}

	args = []string{"a"}
	opts = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}
}

func TestCommandAutocomplete(t *testing.T) {
	commands := command_loader.GetCommands("server")
	sort.Strings(commands)

	args := []string{"server"}
	opts := strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, commands) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", commands, opts)
	}

	args = []string{"server", "cr"}
	opts = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, commands) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", commands, opts)
	}

	args = []string{"serv", "create"}
	commands = []string{""}
	opts = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, commands) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", commands, opts)
	}
}

func TestArgumentsAutocomplete(t *testing.T) {
	opts := options.Get()
	r, _ := command_loader.LoadResource("server")
	c, _ := command_loader.LoadCommand(r, "create")
	arguments := append(c.Arguments(), opts...)
	sort.Strings(arguments)

	args := []string{"server", "create"}
	got := strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--user", "test-user"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu", "0"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--trace"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "not", "valid", "arguments"}
	got = strings.Split(autocomplete.Run(args), " ")
	arguments = []string{""}
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--user"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--from-file"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--trace", "--from-file"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu", "0", "--memoryGB"}
	got = strings.Split(autocomplete.Run(args), " ")
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}
}

func TestEnumerablesAutocomplete(t *testing.T) {
	args := []string{"server", "create", "--type"}
	expected, _ := model_validator.FieldOptions(&server.CreateReq{}, "Type")
	got := strings.Split(autocomplete.Run(args), " ")
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}

func TestOutputOptionAutocomplete(t *testing.T) {
	args := []string{"server", "create", "--output"}
	expected := "json table text"
	got := autocomplete.Run(args)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}
