package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type GetReq struct {
	ServerId string `valid:"required" URIParam:"yes"`
}

type GetRes struct {
	Id          string
	Name        string
	Description string
	GroupId     string
	IsTemplate  bool
	LocationId  string
	OsType      string
	Status      string
	Details     Details
	Type        string
	StorageType string
	ChangeInfo  ChangeInfo
	Links       []models.LinkEntity
}
