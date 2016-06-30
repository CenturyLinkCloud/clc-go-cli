package proxy

import (
	"io/ioutil"
	"os"

	"github.com/centurylinkcloud/clc-go-cli/config"
)

var (
	configDir    string
	configPathFn = config.GetClcHome
)

// Config replaces the configuration directory with a temporary one. It is a
// caller's responsibility to call CloseConfig to release resources.
// Useful for testing modules relied on the config file.
func Config() {
	var err error
	configDir, err = ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	config.SetConfigPathFunc(func() string {
		return configDir
	})
}

// CloseConfig removes the temporary directory created by Config.
func CloseConfig() {
	os.RemoveAll(configDir)
	config.SetConfigPathFunc(configPathFn)
}
