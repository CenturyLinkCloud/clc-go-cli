package state

import (
	"github.com/centurylinkcloud/clc-go-cli/config"
	"io/ioutil"
	"os"
	"path"
)

func readFromFile(name string) ([]byte, error) {
	p, err := config.GetPath()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(path.Join(p, name))
}

func writeToFile(data []byte, name string, perm os.FileMode) error {
	if err := config.CreateIfNotExists(); err != nil {
		return err
	}
	p, err := config.GetPath()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(p, name), data, perm)
}
