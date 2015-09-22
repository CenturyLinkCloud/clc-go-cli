package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/affinity"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/models/network"
	"time"
)

const (
	apiTimeFormat = "2006-01-02T15:04:05Z"
)

type CreateReq struct {
	Name                   string `valid:"required"`
	Description            string `json:",omitempty"`
	GroupId                string
	GroupName              string `json:",omitempty"`
	SourceServerId         string
	SourceServerName       string `json:"-"`
	TemplateName           string `json:"-"`
	IsManagedOs            bool   `json:"IsManagedOS"`
	IsManagedBackup        bool
	PrimaryDns             string             `json:",omitempty"`
	SecondaryDns           string             `json:",omitempty"`
	NetworkId              string             `json:",omitempty"`
	NetworkName            string             `json:",omitempty"`
	IpAddress              string             `json:",omitempty"`
	RootPassword           string             `json:"Password,omitempty"`
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
	TtlString              string             `json:"Ttl,omitempty" argument:"ignore"`
	Packages               []PackageDef       `json:",omitempty"`
	ConfigurationId        string             `json:",omitempty"`
	OsType                 string             `json:",omitempty"`
}

func (c *CreateReq) Validate() error {
	serverIdValues := []string{c.SourceServerId, c.SourceServerName, c.TemplateName}
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

	if c.NetworkName != "" && c.NetworkId != "" {
		return fmt.Errorf("Only one parameter from the following: network-id, network-name may be specified.")
	}

	if c.AntiAffinityPolicyId != "" && c.AntiAffinityPolicyName != "" {
		return fmt.Errorf("Only one parameter from the following: anti-affinity-policy-id, anti-affinity-policy-name may be specified.")
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
		c.TtlString = c.Ttl.Format(apiTimeFormat)
	}
	return nil
}

func (c *CreateReq) InferID(cn base.Connection) error {
	if c.TemplateName != "" {
		c.SourceServerId = c.TemplateName
	}

	if c.SourceServerName != "" {
		s := &Server{ServerName: c.SourceServerName}
		err := s.InferID(cn)
		if err != nil {
			return err
		}
		c.SourceServerId = s.ServerId
	}

	if c.GroupName != "" {
		g := &group.Group{GroupName: c.GroupName}
		err := g.InferID(cn)
		if err != nil {
			return err
		}
		c.GroupId = g.GroupId
	}

	if c.NetworkName != "" {
		ID, err := network.IDByName(cn, "all", c.NetworkName)
		if err != nil {
			return err
		}
		c.NetworkId = ID
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
	case "SourceServerName":
		return GetNames(cn, "all")
	case "GroupName":
		return group.GetNames(cn, "all")
	case "AntiAffinityPolicyName":
		p := &affinity.Policy{}
		return p.GetNames(cn, "PolicyName")
	case "NetworkName":
		return network.GetNames(cn, "all")
	}
	return nil, nil
}
