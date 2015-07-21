package server

import (
	"fmt"
	"time"
)

type CreateReq struct {
	Name                   string `valid:"required"`
	Description            string `json:"omitempty"`
	GroupId                string
	GroupName              string `json:"omitempty"`
	SourceServerId         string
	TemplateId             string
	TemplateName           string              `json:"omitempty"`
	IsManagedOS            bool                `json:"omitempty"`
	PrimaryDns             string              `json:"omitempty"`
	SecondaryDns           string              `json:"omitempty"`
	NetworkId              string              `json:"omitempty"`
	IpAddress              string              `json:"omitempty"`
	Password               string              `json:"omitempty"`
	SourceServerPassword   string              `json:"omitempty"`
	Cpu                    int64               `valid:"required"`
	CpuAutoscalePolicyId   string              `json:"omitempty"`
	MemoryGB               int64               `valid:"required"`
	Type                   string              `valid:"required"`
	StorageType            string              `json:"omitempty"`
	AntiAffinityPolicyId   string              `json:"omitempty"`
	AntiAffinityPolicyName string              `json:"omitempty"`
	CustomFields           []CustomFieldDef    `json:"omitempty"`
	AdditionalDisks        []AdditionalDiskDef `json:"omitempty"`
	Ttl                    time.Time           `json:"omitempty"`
	Packages               []PackageDef        `json:"omitempty"`
}

type CustomFieldDef struct {
	Id    string
	Value string
}

type AdditionalDiskDef struct {
	Path   string
	SizeGB int
	Type   string
}

type PackageDef struct {
	PackageId  string
	Parameters map[string]string
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

	return nil
}

func (c *CreateReq) ApplyDefaultBehaviour() error {
	if c.TemplateId != "" {
		c.SourceServerId = c.TemplateId
	}
	return nil
	//TODO: implement searching templates by name
	//TODO: implement searching groups by names
}
