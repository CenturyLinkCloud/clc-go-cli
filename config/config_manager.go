package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const configFileName = "config.yml"

func LoadConfig() (*Config, error) {
	c, err := loadConfigInner()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config file: %s", err.Error())
	}
	return c, nil
}

func Save(c *Config) error {
	err := saveInner(c)
	if err != nil {
		return fmt.Errorf("Failed to save config file: %s", err.Error())
	}
	return nil
}

func saveInner(c *Config) error {
	clcHome := GetClcHome()
	if clcHome == "" {
		return nil
	}

	file := path.Join(clcHome, configFileName)
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0777)
}

func loadConfigInner() (*Config, error) {
	c := &Config{}

	clcHome := GetClcHome()
	if clcHome == "" {
		return c, nil
	}

	CreateIfNotExists()

	filepath := path.Join(clcHome, configFileName)
	_, err := os.Stat(filepath)
	if err != nil {
		// Config file does not exist. It's ok
		if os.IsNotExist(err) {
			return c, nil
		}
		return nil, fmt.Errorf("error stat file %s: %s", filepath, err)
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %s", filepath, err)
	}
	defer f.Close()

	// Read config file
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", filepath, err)
	}
	// Unmarshal config
	err = yaml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

var GetClcHome = func() string {
	clcHome := os.Getenv("CLC_HOME")
	if clcHome != "" {
		return clcHome
	}

	homeDir := os.Getenv(HOME_VAR)
	if homeDir == "" {
		return ""
	}

	return path.Join(homeDir, CONFIG_FOLDER_NAME)
}

func CreateIfNotExists() error {
	clcHome := GetClcHome()
	if _, err := os.Stat(clcHome); os.IsNotExist(err) {
		err := os.MkdirAll(clcHome, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
