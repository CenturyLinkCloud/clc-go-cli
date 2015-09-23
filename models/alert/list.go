package alert

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type ListRes struct {
	Items []Entity
	Links []models.LinkEntity
}
