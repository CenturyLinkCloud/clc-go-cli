package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"time"
)

type FullCustomFieldDef struct {
	Id           string
	Value        string
	Name         string
	DisplayValue string
}

type CustomFieldDef struct {
	Id    string
	Value string
}

type Disk struct {
	Id             string
	SizeGB         float64
	PartitionPaths []string
}

type Partition struct {
	SizeGB float64
	Path   string
}

type AddDiskRequest struct {
	Path   string
	SizeGB int64
	Type   string
}

type KeepDiskRequest struct {
	DiskId string
	SizeGB int64
}

type PackageDef struct {
	PackageId  string
	Parameters map[string]string
}

type Details struct {
	IpAddresses       []IPAddresses
	AlertPolicies     []alert.AlertPolicy
	Cpu               int64
	DiskCount         int64
	HostName          string
	InMaintenanceMode bool
	MemoryMB          int64
	PowerState        string
	StorageGB         int64
	Disks             []Disk
	Partitions        []Partition
	Snapshots         []Snapshot
	CustomFields      []FullCustomFieldDef
}

type IPAddresses struct {
	Public   string `json:",omitempty"`
	Internal string
}

type Snapshot struct {
	Name  string
	Links []models.LinkEntity
}

type ChangeInfo struct {
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedDate time.Time
	ModifiedBy   string
}
