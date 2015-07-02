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

	p, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(p)
	exist := true
	if os.IsNotExist(err) {
		f, err = os.Create(p)
		exist = false
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
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

var getConfigPath = func() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(u.HomeDir, ".clc", "config.yml"), nil
}














