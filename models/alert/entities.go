package alert

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AlertPolicy struct {
	Id    string
	Name  string
	Links []models.LinkEntity
}

type Entity struct {
	Id       string
	Name     string
	Actions  []Action
	Triggers []Trigger
	Links    []models.LinkEntity
}

type Action struct {
	Action   string
	Settings Settings
}

type Settings struct {
	Recipients []string
}

type Trigger struct {
	Metric    string
	Duration  string
	Threshold float64
}
