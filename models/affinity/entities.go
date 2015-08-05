package affinity

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id       string
	Name     string
	Location string
	Links    []models.LinkEntity
}
