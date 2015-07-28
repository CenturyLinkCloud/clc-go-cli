package group

import (
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type CreateReq struct {
	Name          string `valid:"required"`
	Description   string
	ParentGroupId string `valid:"required"`
	CustomFields  []server.CustomFieldDef
}
