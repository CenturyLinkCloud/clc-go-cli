package backup

import (
	"fmt"
	"time"

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

type GetStoredData struct {
	AccountPolicyId string `URIParam:"yes" valid:"required"`
	ServerPolicyId  string `URIParam:"yes" valid:"required"`
	SearchDate      string `URIParam:"yes" valid:"required"`
}

func (g *GetStoredData) Validate() error {
	if _, err := time.Parse(base.DATE_FORMAT, g.SearchDate); err != nil {
		return fmt.Errorf("The search-date value must be a valid date in YYYY-MM-DD format")
	}
	return nil
}

type UpdateServerPolicy struct {
	AccountPolicyId string                        `URIParam:"yes" valid:"required"`
	ServerPolicyId  string                        `URIParam:"yes" valid:"required"`
	Operations      []UpdateServerPolicyOperation `argument:"ignore" json:"operations"`
	Status          string                        `json:"-" oneOf:"ACTIVE,INACTIVE" valid:"required"`
}

func (u *UpdateServerPolicy) Validate() error {
	u.Operations = append(u.Operations, UpdateServerPolicyOperation{
		Op:    "replace",
		Path:  "/status",
		Value: u.Status,
	})
	return nil
}

type UpdateServerPolicyOperation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type GetRestoreDetails struct {
	AccountPolicyId         string `URIParam:"yes" valid:"required"`
	ServerPolicyId          string `URIParam:"yes" valid:"required"`
	BackupFinishedStartDate string `URIParam:"yes" valid:"required"`
	BackupFinishedEndDate   string `URIParam:"yes" valid:"required"`
	Limit                   string `URIParam:"yes"`
	Offset                  string `URIParam:"yes"`
	InRetentionOnly         string `URIParam:"yes" oneOf:"true,false"`
	SortBy                  string `URIParam:"yes" oneOf:"policyId,retentionDay,backupStartedDate,backupFinishedDate,retentionExpiredDate,backupStatus,filesTransferredToStorage,bytesTransferredToStorage,filesFailedTransferToStorage,bytesFailedToTransfer,unchangedFilesNotTransferred,unchangedBytesNotTransferred,filesRemovedFromDisk,bytesRemovedFromDisk"`
	AscendingSort           string `URIParam:"yes" oneOf:"true,false"`
}

func (g *GetRestoreDetails) Validate() error {
	if _, err := time.Parse(base.DATE_FORMAT, g.BackupFinishedEndDate); err != nil {
		return fmt.Errorf("The backup-finished-end-date value must be a valid date in YYYY-MM-DD format")
	}
	if _, err := time.Parse(base.DATE_FORMAT, g.BackupFinishedStartDate); err != nil {
		return fmt.Errorf("The backup-finished-start-date value must be a valid date in YYYY-MM-DD format")
	}
	return nil
}

type GetRestoreDetailsRes struct {
	Limit      int64
	NextOffset int64
	Offset     int64
	TotalCount int64
	Results    []RestoreDetails
}
