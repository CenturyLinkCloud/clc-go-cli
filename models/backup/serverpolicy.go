package backup

import "github.com/centurylinkcloud/clc-go-cli/models/server"

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
	Results    []struct {
		AccountPolicyId  string
		ClcAccountAlias  string
		ExpirationDate   int64
		ServerId         string
		ServerPolicyId   string
		Status           string
		StorageAccountId string
		StorageRegion    string
		UnsubscribedDate int64
	}
}
