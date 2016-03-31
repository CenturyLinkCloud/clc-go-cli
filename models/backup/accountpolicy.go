package backup

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type AccountPolicyReq struct {
	AccountAlias string `argument:"ignore" json:"clcAccountAlias"`

	BackupIntervalHours    *int64   `json:"backupIntervalHours" valid:"required"`
	ExcludedDirectoryPaths []string `json:"excludedDirectoryPaths,omitempty"`
	Name                   string   `json:"name" valid:"required"`
	OsType                 string   `json:"osType" oneOf:"Linux,Windows" valid:"required"`
	Paths                  []string `json:"paths" valid:"required"`
	RetentionDays          *int64   `json:"retentionDays" valid:"required"`
}

func (a *AccountPolicyReq) InferID(cn base.Connection) error {
	a.AccountAlias = cn.GetAccountAlias()
	return nil
}

func (a *AccountPolicyReq) GetNames(cn base.Connection, property string) ([]string, error) {
	return nil, nil
}
