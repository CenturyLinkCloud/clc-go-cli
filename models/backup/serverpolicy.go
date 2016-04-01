package backup

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type GetServerPolicies struct {
	server.Server `argument:"compose" URIParam:"ServerId"`
	WithStatus    string `URIParam:"yes" json:"-"`
}

type GetAccountServerPolicies struct {
	PolicyId      string `URIParam:"yes" valid:"required"`
	Limit         string `URIParam:"yes"`
	Offset        string `URIParam:"yes"`
	WithStatus    string `URIParam:"yes"`
	SortBy        string `URIParam:"yes" oneOf:"status,storageRegion,serverId,serverPolicyId"`
	AscendingSort string `URIParam:"yes" oneOf:"true,false"`
}

type AccountServerPoliciesRes struct {
	Limit      int64
	NextOffset int64
	Offset     int64
	TotalCount int64
	Results    []ServerPolicy
}

type GetAccountServerPolicy struct {
	AccountPolicyId string `URIParam:"yes" valid:"required"`
	ServerPolicyId  string `URIParam:"yes" valid:"required"`
}

type CreateServerPolicy struct {
	AccountPolicyId  string `URIParam:"yes" valid:"required" json:"accountPolicyId"`
	AccountAlias     string `json:"clcAccountAlias" argument:"ignore"`
	server.Server    `argument:"composed" json:"-"`
	ServerID         string `json:"serverId" argument:"ignore"`
	StorageAccountId string `json:"storageAccountId"`
	StorageRegion    string `json:"storageRegion" valid:"required"`
}

func (c *CreateServerPolicy) InferID(cn base.Connection) error {
	c.AccountAlias = cn.GetAccountAlias()
	if err := c.Server.InferID(cn); err != nil {
		return err
	}
	c.ServerID = c.Server.ServerId
	return nil
}

type DeleteServerPolicy struct {
	AccountPolicyId string `URIParam:"yes" valid:"required"`
	ServerPolicyId  string `URIParam:"yes" valid:"required"`
}
