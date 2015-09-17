package autocomplete_test

import (
	"fmt"
	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/autocomplete"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"
)

type (
	inputModel struct {
		Property string
	}
	inputModelDCCentric struct {
		DataCenter string
		Property   string
		Property2  string
	}
)

var (
	testCommand = &commands.CommandBase{
		Input: &inputModel{},
		ExcInfo: commands.CommandExcInfo{
			Resource: "resource",
			Command:  "DCagnostic",
		},
	}
	testCommandDCCentric = &commands.CommandBase{
		Input: &inputModelDCCentric{},
		ExcInfo: commands.CommandExcInfo{
			Resource: "resource",
			Command:  "DCcentric",
		},
	}
)

func (i *inputModel) InferID(cn base.Connection) error {
	return nil
}

func (i *inputModel) GetNames(cn base.Connection, property string) ([]string, error) {
	return []string{"Value 1", "Value2"}, nil
}

func (i *inputModelDCCentric) InferID(cn base.Connection) error {
	return nil
}

func (i *inputModelDCCentric) GetNames(cn base.Connection, property string) ([]string, error) {
	if i.DataCenter == "" {
		return nil, fmt.Errorf("A data center must be set.")
	}
	return []string{"Value 1", "Value2"}, nil
}

func TestResourceAutocomplete(t *testing.T) {
	resources := command_loader.GetResources()
	sort.Strings(resources)

	args := []string{""}
	opts := strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}

	args = []string{"serve"}
	opts = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}

	args = []string{"a"}
	opts = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, resources) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", resources, opts)
	}
}

func TestCommandAutocomplete(t *testing.T) {
	commands := command_loader.GetCommands("server")
	sort.Strings(commands)

	args := []string{"server"}
	opts := strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, commands) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", commands, opts)
	}

	args = []string{"server", "cr"}
	opts = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(opts)
	if !reflect.DeepEqual(opts, commands) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", commands, opts)
	}

	args = []string{"serv", "create"}
	commands = []string{""}
	opts = strings.Split(autocomplete.Run(args), autocomplete.SEP)
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
	got := strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--user", "test-user"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu", "0"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--trace"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "not", "valid", "arguments"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	arguments = []string{""}
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--user"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--from-file"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--trace", "--from-file"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}

	args = []string{"server", "create", "--cpu", "0", "--memory-gb"}
	got = strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, arguments) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", arguments, got)
	}
}

func TestEnumerablesAutocomplete(t *testing.T) {
	args := []string{"server", "create", "--type"}
	expected, _ := model_validator.FieldOptions(&server.CreateReq{}, "Type")
	got := strings.Split(autocomplete.Run(args), autocomplete.SEP)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}

func TestOutputOptionAutocomplete(t *testing.T) {
	args := []string{"server", "create", "--output"}
	expected := strings.Join([]string{"json", "table", "text"}, autocomplete.SEP)
	got := autocomplete.Run(args)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}

func TestProfileOptionAutocomplete(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	config.SetConfigPathFunc(func() (string, error) {
		return dir, nil
	})
	defer func() {
		os.RemoveAll(dir)
	}()

	f, err := os.Create(path.Join(dir, "config.yml"))
	defer f.Close()
	if err != nil {
		t.Error(err)
	}

	profiles := map[string]config.Profile{
		"Default": config.Profile{
			User:     "Vincent",
			Password: "Vega",
		},
		"Empty": config.Profile{},
		"Profile2": config.Profile{
			User:     "Mia",
			Password: "Wallace",
		},
	}
	bytes, err := yaml.Marshal(config.Config{Profiles: profiles})
	if err != nil {
		t.Error(err)
	}
	f.Write(bytes)

	args := []string{"server", "create", "--profile"}
	expected := []string{"Default", "Empty", "Profile2"}
	got := strings.Split(autocomplete.Run(args), autocomplete.SEP)
	sort.Strings(got)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\n Expected: %s,\n obtained: %s", expected, got)
	}
}

func TestAPIRelatedPropertiesAutocomplete(t *testing.T) {
	var got, expected interface{}

	proxy.Config()
	defer proxy.CloseConfig()
	proxy.Login()
	defer proxy.CloseLogin()

	cli.AllCommands = append(cli.AllCommands, testCommand)
	cli.AllCommands = append(cli.AllCommands, testCommandDCCentric)

	c := &config.Config{User: "user", Password: "password"}
	err := config.Save(c)
	if err != nil {
		panic(err)
	}

	// Test a data-center-agnostic command.
	got = autocomplete.Run([]string{"resource", "DCagnostic", "--property"})
	expected = "Value 1\nValue2"
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected:%v\nGot:%v", expected, got)
	}

	// Test a data-center-centric command with empty config.
	got = autocomplete.Run([]string{"resource", "DCcentric", "--property"})
	expected = ""
	if got != expected {
		t.Errorf("Invalid result.\nExpected nothing\nGot:%v", got)
	}

	// Test a data-center-centric command with a default data center set.
	c.DefaultDataCenter = "DC"
	config.Save(c)
	got = autocomplete.Run([]string{"resource", "DCcentric", "--property2"})
	expected = "Value 1\nValue2"
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Invalid result.\nExpected:%v\nGot:%v", expected, got)
	}
}
