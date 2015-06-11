package server

import (
	"time"
)

type CreateReq struct {
	Name                   string
	Description            string
	GroupId                string
	GroupName              string
	SourceServerId         string
	SourceServerName       string
	TemplateId             string
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
