package crossdc_firewall

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id                  string
	Status              string
	Enabled             bool
	SourceCidr          string
	SourceAccount       string
	SourceLocation      string
	DestinationCidr     string
	DestinationAccount  string
	DestinationLocation string
	Links               []models.LinkEntity
}
