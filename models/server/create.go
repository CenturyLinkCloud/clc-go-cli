package server

import (
	"fmt"
	"time"
)

type CreateReq struct {
	Name                   string `valid:required`
	Description            string
	GroupId                string
	GroupName              string
	SourceServerId         string
	TemplateId             string `json:"-"`
	TemplateName           string
	IsManagedOS            bool
	PrimaryDns             string
	SecondaryDns           string
	NetworkId              string
	IpAddress              string
	Password               string
	SourceServerPassword   string
	Cpu                    int
	CpuAutoscalePolicyId   string
	MemoryGB               int
	Type                   string
	StorageType            string
	AntiAffinityPolicyId   string
	AntiAffinityPolicyName string
	CustomFields           []CustomFieldDef
	AdditionalDisks        []AdditionalDiskDef
	Ttl                    time.Time
	Packages               []PackageDef
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

	return nil
}

func (c *CreateReq) ApplyDefaultBehaviour() error {
	if c.TemplateId != "" {
		c.SourceServerId = c.TemplateId
	}
	return nil
	//TODO: implement searching templates by name
}
