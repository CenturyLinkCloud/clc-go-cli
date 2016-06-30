package state

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/centurylinkcloud/clc-go-cli/config"
)

func ReadFromFile(name string) ([]byte, error) {
	clcHome := config.GetClcHome()
	if clcHome == "" {
		return nil, errors.New("Unable to save state: no home")
	}
	return ioutil.ReadFile(path.Join(clcHome, name))
}

func WriteToFile(data []byte, name string, perm os.FileMode) error {
	clcHome := config.GetClcHome()
	if clcHome == "" {
		return nil
	}
	if err := config.CreateIfNotExists(); err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(clcHome, name), data, perm)
}

func GetFileInfo(name string) (os.FileInfo, error) {
	clcHome := config.GetClcHome()
	if clcHome == "" {
		return nil, errors.New("Unable to save state: no home")
	}
	return os.Stat(path.Join(clcHome, name))
}
