package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func StoreApi(api []*ApiDef) error {
	data, err := json.Marshal(api)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("api.json", data, os.ModeExclusive)
	return err
}

func LoadApi() ([]*ApiDef, error) {
	data, err := ioutil.ReadFile("api.json")
	if err != nil {
		return nil, err
	}
	var api []*ApiDef
	err = json.Unmarshal(data, api)
	return api, err
}
