package main

import (
	"encoding/json"
	"io/ioutil"
)

func StoreApi(api []*ApiDef, apiPath string) error {
	data, err := json.MarshalIndent(api, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(apiPath, data, 0664)
	return err
}

func LoadApi(apiPath string) ([]*ApiDef, error) {
	data, err := ioutil.ReadFile(apiPath)
	if err != nil {
		return nil, err
	}
	var api []*ApiDef
	err = json.Unmarshal(data, api)
	return api, err
}
