package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type GetReq struct {
	Server `argument:"composed" URIParam:"ServerId"`
}

type GetRes struct {
	Id          string
	Name        string
	DisplayName string
	Description string
	GroupId     string
	IsTemplate  bool
	LocationId  string
	OsType      string
	Os          string
	Status      string
	Details     Details
	Type        string
	ChangeInfo  models.ChangeInfo
	Links       []models.LinkEntity
}
