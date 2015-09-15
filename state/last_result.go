package state

import (
	"github.com/centurylinkcloud/clc-go-cli/formatters"
)

func SaveLastResult(r interface{}) error {
	jsonFormatter := formatters.JsonFormatter{}
	str, err := jsonFormatter.FormatOutput(r)
	if err != nil {
		return err
	}
	return WriteToFile([]byte(str), ".last_result", 0777)
}

func LoadLastResult() ([]byte, error) {
	return ReadFromFile(".last_result")
}
