package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

func LoadConfig() (*Config, error) {
	c, err := loadConfigInner()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config file, error: %s", err.Error())
	}
	return c, nil
}

func Save(c *Config) error {
	p, err := GetPath()
	if err != nil {
		return err
	}
	file := path.Join(p, "config.yml")
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	ioutil.WriteFile(file, data, 0777)
	return nil
}

func loadConfigInner() (*Config, error) {
	c := &Config{}

	p, err := GetPath()
	if err != nil {
		return nil, err
	}
	CreateIfNotExists()
	var f *os.File
	filepath := path.Join(p, "config.yml")
	exist := true
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		f, err = os.Create(filepath)
		defer f.Close()
		if err != nil {
			return nil, err
		}
		exist = false
	} else {
		f, err = os.Open(filepath)
		defer f.Close()
		if err != nil {
			return nil, err
		}
	}
	if !exist {
		content, err := yaml.Marshal(c)
		if err != nil {
			return nil, err
		}
		_, err = f.Write(content)
		if err != nil {
			return nil, err
		}
	} else {
		content, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(content, c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

var GetPath = func() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return "", fmt.Errorf("The HOME environment variable is not set. Please, set it so that we can store your configuration there")
	}

	return path.Join(homeDir, CONFIG_FOLDER_NAME), nil
}

func CreateIfNotExists() error {
	p, err := GetPath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
