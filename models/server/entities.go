package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"time"
)

type CustomFieldDef struct {
	Id           string
	Value        string
	Name         string
	DisplayValue string
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

type DiskRequest struct {
	Path   string
	SizeGB int64
	Type   string
}

type PackageDef struct {
	PackageId  string
	Parameters map[string]string
}

type Details struct {
	IPAddresses       []IPAddresses
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
	CustomFields      []CustomFieldDef
}

type IPAddresses struct {
	Public   string
	Internal string
}

type Snapshot struct {
	Name  string
	Links []models.LinkEntity
}

type ChangeInfo struct {
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedData time.Time
	ModifiedBy   string
}
