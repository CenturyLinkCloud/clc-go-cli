package alert

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AlertPolicy struct {
	Id    string
	Name  string
	Links []models.LinkEntity
}
