package config_test

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

var configDir, configPath string

func initTest(confPath string, content []byte, err error) {
	config.SetConfigPathFunc(func() (string, error) {
		if confPath != "" {
			return confPath, nil
		}
		if err != nil {
			return "", err
		}
		configDir, err = ioutil.TempDir(os.TempDir(), "")
		if err != nil {
			return "", err
		}
		configPath = path.Join(configDir, "config.yml")
		if content != nil {
			f, err := os.Create(configPath)
			if err != nil {
				return "", err
			}
			defer f.Close()
			_, err = f.Write(content)
			if err != nil {
				return "", err
			}
		}
		return configPath, err
	})
}

func finishTest() {
	os.RemoveAll(configDir)
}

func TestCreateNewConfig(t *testing.T) {
	initTest("", nil, nil)
	defer finishTest()
	c, err := config.LoadConfig()
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(c, &config.Config{}) {
		t.Errorf("Incorrect config obtained: %#v", c)
	}
}

func TestErrorConfigPath(t *testing.T) {
	initTest("", nil, fmt.Errorf("Test error"))
	defer finishTest()
	_, err := config.LoadConfig()
	if err == nil || err.Error() != "Failed to load config file, error: Test error" {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCantCreteConfigDir(t *testing.T) {
	initTest("\000_incorrect_path", nil, nil)
	defer finishTest()
	_, err := config.LoadConfig()
	if err == nil || err.Error() != "Failed to load config file, error: open \000_incorrect_path: invalid argument" {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestInvalidConfigContent(t *testing.T) {
	initTest("", []byte("some invalid content"), nil)
	defer finishTest()
	_, err := config.LoadConfig()
	if err == nil || !strings.HasPrefix(err.Error(), "Failed to load config file, error: yaml: unmarshal errors:") {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCorrectConfigContent(t *testing.T) {
	c := &config.Config{User: "user", Password: "password", Profiles: map[string]config.Profile{}}
	content, err := yaml.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	initTest("", content, nil)
	defer finishTest()
	c1, err := config.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(c, c1) {
		t.Errorf("Expected: %#v, obtained: %#v", c, c1)
	}
}
