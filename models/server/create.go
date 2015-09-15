package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/affinity"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
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
	serverIdValues := []string{c.SourceServerId, c.TemplateName}
	numNonEmpty := 0
	for _, item := range serverIdValues {
		if item != "" {
			numNonEmpty++
		}
	}
	if numNonEmpty > 1 || numNonEmpty == 0 {
		return fmt.Errorf("Exactly one parameter from the following: source-server-id, source-server-name, template-name must be specified.")
	}

	if (c.GroupName == "") == (c.GroupId == "") {
		return fmt.Errorf("Exactly one parameter from the following: group-id, group-name must be specified.")
	}

	if (c.AntiAffinityPolicyId == "") == (c.AntiAffinityPolicyName == "") {
		return fmt.Errorf("Exactly one parameter from the following: anti-affinity-policy-id, anti-affinity-policy-name must be specified.")
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
	zeroTime := time.Time{}
	if c.Ttl != zeroTime {
		c.TtlString = c.Ttl.Format(timeFormat)
	}
	return nil

	//TODO: implement searching groups by names
}

func (c *CreateReq) InferID(cn base.Connection) error {
	if c.TemplateName != "" {
		c.SourceServerId = c.TemplateName
	}

	if c.GroupName != "" {
		g := &group.Group{GroupName: c.GroupName}
		err := g.InferID(cn)
		if err != nil {
			return err
		}
		c.GroupId = g.GroupId
	}

	if c.AntiAffinityPolicyName != "" {
		p := &affinity.Policy{PolicyName: c.AntiAffinityPolicyName}
		err := p.InferID(cn)
		if err != nil {
			return err
		}
		c.AntiAffinityPolicyId = p.PolicyId
	}
	return nil
}

func (c *CreateReq) GetNames(cn base.Connection, property string) ([]string, error) {
	switch property {
	case "TemplateName":
		return LoadTemplates(cn)
	case "GroupName":
		return group.GetNames(cn, "all")
	case "AntiAffinityPolicyName":
		p := &affinity.Policy{}
		return p.GetNames(cn, "PolicyName")
	}
	return nil, nil
}
