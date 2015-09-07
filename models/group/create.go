package group

import (
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type CreateReq struct {
	Name          string `valid:"required"`
	Description   string
	ParentGroupId string `valid:"required"`
	CustomFields  []customfields.Def
}
