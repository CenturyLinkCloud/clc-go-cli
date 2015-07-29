package group

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type RestoreReq struct {
	GroupId       string `valid:"required" URIParam:"yes"`
	TargetGroupId string `valid:"required"`
}

type RestoreRes struct {
	IsQueued     bool
	Links        []models.LinkEntity
	ErrorMessage string
}
