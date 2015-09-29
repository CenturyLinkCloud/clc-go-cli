package models

import "github.com/centurylinkcloud/clc-go-cli/base"

type ChangeInfo struct {
	CreatedDate  base.Time
	CreatedBy    string
	ModifiedDate base.Time
	ModifiedBy   string
}
