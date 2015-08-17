package autocomplete_test

import (
	"github.com/centurylinkcloud/clc-go-cli/autocomplete"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
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

	args = []string{"server", "create"}
	commands = []string{""}
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
