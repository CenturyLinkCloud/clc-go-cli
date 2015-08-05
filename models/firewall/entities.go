package firewall

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id                 string
	Status             string
	Enabled            bool
	Source             []string
	Destination        []string
	DestinationAccount string
	Ports              []string
	Links              []models.LinkEntity
}
