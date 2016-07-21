package datacenter

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type GetReq struct {
	DataCenter        string `valid:"required"`
	WithComputeLimits base.NilField
	WithNetworkLimits base.NilField
	WithAvailableOvfs base.NilField
	WithLoadBalancers base.NilField
}

type GetRes struct {
	Id            string
	Name          string
	ComputeLimits *DcComputeLimits  `json:",omitempty"`
	NetworkLimits *DcNetworkLimits  `json:",omitempty"`
	AvailableOVFs *[]DcAvailableOVF `json:",omitempty"`
	LoadBalancers *[]DcLoadBalancer `json:",omitempty"`
	Links         models.Links
}

type DcComputeLimits struct {
	Cpu       DcResourceLimit
	MemoryGB  DcResourceLimit
	StorageGB DcResourceLimit
}

type DcNetworkLimits struct {
	Networks DcResourceLimit
}

type DcResourceLimit struct {
	Value     int
	Inherited bool
}

type DcAvailableOVF struct {
	Id            string
	Name          string
	StorageSizeGB int
	CpuCount      int
	MemorySizeMB  int
}

type DcLoadBalancer struct {
	Name        string
	Description string
	IpAddress   string
	Status      string
}
