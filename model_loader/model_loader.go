package model_loader

import (
	"encoding/json"
)

func LoadModel(parsedArgs map[string]interface{}, inputModel interface{}) error {
	str, err := json.Marshal(parsedArgs)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, inputModel)
	return err
}
