package backup

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
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

type AccountPoliciesReq struct {
	Limit         string `URIParam:"yes"`
	Offset        string `URIParam:"yes"`
	WithStatus    string `URIParam:"yes"`
	SortBy        string `URIParam:"yes" oneOf:"status,osType,name,policyId,backupIntervalHours,retentionDays"`
	AscendingSort string `URIParam:"yes" oneOf:"true,false"`
}

type AccountPoliciesRes struct {
	Limit      int64
	NextOffset int64
	Offset     int64
	Results    []AccountPolicy
	TotalCount int64
}

type GetAccountPolicy struct {
	PolicyId string `URIParam:"yes" valid:"required"`
}

type AllowedAccountPoliciesReq struct {
	server.Server `argument:"composed" URIParam:"ServerId"`
	Limit         string `URIParam:"yes"`
	Offset        string `URIParam:"yes"`
	WithStatus    string `URIParam:"yes"`
	SortBy        string `URIParam:"yes" oneOf:"status,osType,name,policyId,backupIntervalHours,retentionDays"`
	AscendingSort string `URIParam:"yes" oneOf:"true,false"`
}
