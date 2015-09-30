package state_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
	"github.com/centurylinkcloud/clc-go-cli/state"
)

func TestReadWrite(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	data := "Lorem ipsum"
	err := state.WriteToFile([]byte(data), "f", 0666)
	if err != nil {
		t.Fatal(err)
	}
	var bytes []byte
	bytes, err = state.ReadFromFile("f")
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != data {
		t.Errorf("Invalid result\n Expected %s\nGot %s", data, string(bytes))
	}
}

func TestGetFileInfo(t *testing.T) {
	proxy.Config()
	defer proxy.Config()

	p, err := config.GetPath()
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile(path.Join(p, "test"), []byte{}, 0666)
	defer os.Remove(path.Join(p, "test"))
	var info os.FileInfo
	info, err = state.GetFileInfo("test")
	if err != nil {
		t.Fatal(err)
	}
	if info.Name() != "test" {
		t.Errorf("Invalid result\nGot a wrong file name: %s", info.Name())
	}
}
