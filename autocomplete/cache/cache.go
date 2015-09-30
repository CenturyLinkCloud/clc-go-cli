package cache

import (
	"encoding/json"
	"github.com/centurylinkcloud/clc-go-cli/state"
	"time"
)

var (
	LONG_AUTOCOMPLETE_REFRESH_TIMEOUT = 30 // seconds
)

func Put(key string, opts []string) {
	data, err := json.Marshal(opts)
	if err == nil {
		state.WriteToFile(data, key, 0666)
	}
}

func Get(key string) ([]string, bool) {
	info, err := state.GetFileInfo(key)
	if err != nil {
		return nil, false
	}

	if time.Now().Sub(info.ModTime()) > time.Second*time.Duration(LONG_AUTOCOMPLETE_REFRESH_TIMEOUT) {
		return nil, false
	}

	data, err := state.ReadFromFile(key)
	if err != nil {
		return nil, false
	}

	opts := []string{}
	err = json.Unmarshal(data, &opts)
	if err != nil {
		return nil, false
	}

	return opts, true
}
