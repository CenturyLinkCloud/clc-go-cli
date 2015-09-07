package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type CreateReq struct {
	Name                   string `valid:"required"`
	Description            string `json:",omitempty"`
	GroupId                string
	GroupName              string `json:",omitempty"`
	SourceServerId         string
	TemplateId             string
	TemplateName           string             `json:",omitempty"`
	IsManagedOs            bool               `json:",omitempty"`
	IsManagedBackup        bool               `json:",omitempty"`
	PrimaryDns             string             `json:",omitempty"`
	SecondaryDns           string             `json:",omitempty"`
	NetworkId              string             `json:",omitempty"`
	IpAddress              string             `json:",omitempty"`
	RootPassword           string             `json:",omitempty"`
	SourceServerPassword   string             `json:",omitempty"`
	Cpu                    int64              `valid:"required"`
	CpuAutoscalePolicyId   string             `json:",omitempty"`
	MemoryGb               int64              `valid:"required"`
	Type                   string             `valid:"required" oneOf:"standard,hyperscale,bareMetal"`
	StorageType            string             `json:",omitempty" oneOf:"standard,premium,hyperscale"`
	AntiAffinityPolicyId   string             `json:",omitempty"`
	AntiAffinityPolicyName string             `json:",omitempty"`
	CustomFields           []customfields.Def `json:",omitempty"`
	AdditionalDisks        []AddDiskRequest   `json:",omitempty"`
	Ttl                    time.Time          `json:"-"`
	TtlString              string             `json:"Ttl,omitempty"`
	Packages               []PackageDef       `json:",omitempty"`
	ConfigurationId        string             `json:",omitempty"`
	OsType                 string             `json:",omitempty"`
}

func (c *CreateReq) Validate() error {
	serverIdValues := []string{c.SourceServerId, c.TemplateId, c.TemplateName}
	numNonEmpty := 0
	for _, item := range serverIdValues {
		if item != "" {
			numNonEmpty++
		}
	}
	if numNonEmpty > 1 || numNonEmpty == 0 {
		return fmt.Errorf("Exactly one parameter from the following: source-server-id, source-server-name, template-id, template-name must be specified.")
	}

	if (c.GroupName == "") == (c.GroupId == "") {
		return fmt.Errorf("Exactly one parameter from the following: group-id, group-name must be specified.")
	}

	if c.Type == "bareMetal" {
		if c.ConfigurationId == "" {
			return fmt.Errorf("ConfigurationId: required for bare metal servers.")
		}
		if c.OsType == "" {
			return fmt.Errorf("OsType: required for bare metal servers.")
		}
	}

	return nil
}

func (c *CreateReq) ApplyDefaultBehaviour() error {
	if c.TemplateId != "" {
		c.SourceServerId = c.TemplateId
	}

	zeroTime := time.Time{}
	if c.Ttl != zeroTime {
		c.TtlString = c.Ttl.Format(timeFormat)
	}
	return nil

	//TODO: implement searching templates by name
	//TODO: implement searching groups by names
}
