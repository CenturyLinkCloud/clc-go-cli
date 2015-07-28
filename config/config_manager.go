package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

func LoadConfig() (*Config, error) {
	c, err := loadConfigInner()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config file, error: %s", err.Error())
	}
	return c, nil
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
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(u.HomeDir, ".clc"), nil
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
