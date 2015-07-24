package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type ServerRes struct {
	Server       string
	IsQueued     bool
	ErrorMessage string
	Links        []models.LinkEntity
}
