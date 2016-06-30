package config_test

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/config"
)

var configDir, configPath string

func initTest(confPath string, content []byte) {
	config.SetConfigPathFunc(func() string {
		var err error

		if confPath != "" {
			return confPath
		}
		if configDir == "" {
			configDir, err = ioutil.TempDir(os.TempDir(), "")
			if err != nil {
				panic(err.Error())
			}
		}

		if content != nil {
			configPath = path.Join(configDir, "config.yml")
			f, err := os.Create(configPath)
			if err != nil {
				panic(err.Error())
			}
			defer f.Close()
			_, err = f.Write(content)
			if err != nil {
				panic(err.Error())
			}
		}
		return configDir
	})
}

func finishTest() {
	os.RemoveAll(configDir)
	configDir = ""
}

func TestCreateNewConfig(t *testing.T) {
	initTest("", nil)
	defer finishTest()
	c, err := config.LoadConfig()
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(c, &config.Config{}) {
		t.Errorf("Incorrect config obtained: %#v", c)
	}
}

func TestLoadConfigWhenFileDoesNotExist(t *testing.T) {
	configDir = "/tmp/non-existent_path"
	initTest(configDir, nil)
	defer finishTest()
	c, err := config.LoadConfig()
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(c, &config.Config{}) {
		t.Errorf("Incorrect config obtained: %#v", c)
	}
}

func TestLoadConfigWithInvalidConfigContent(t *testing.T) {
	initTest("", []byte("some invalid content"))
	defer finishTest()
	_, err := config.LoadConfig()
	if err == nil || !strings.HasPrefix(err.Error(), "Failed to load config file: yaml: unmarshal errors:") {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestLoadConfigWithCorrectConfigContent(t *testing.T) {
	c := &config.Config{User: "user", Password: "password", Profiles: map[string]config.Profile{}}
	initTest("", nil)
	defer finishTest()
	config.Save(c)
	c1, err := config.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(c, c1) {
		t.Errorf("Expected: %#v, obtained: %#v", c, c1)
	}
}
