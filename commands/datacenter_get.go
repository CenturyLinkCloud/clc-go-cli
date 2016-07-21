package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

type DatacenterGet struct {
	CommandBase
}

func NewDatacenterGet(info CommandExcInfo) *DatacenterGet {
	dcGet := DatacenterGet{}
	dcGet.ExcInfo = info
	dcGet.Input = &datacenter.GetReq{}
	return &dcGet
}

func (dcGet *DatacenterGet) Execute(cn base.Connection) error {
	var err error

	input := dcGet.Input.(*datacenter.GetReq)
	dcGet.Output, err = datacenter.Get(cn, input.DataCenter, input.WithComputeLimits.Set,
		input.WithNetworkLimits.Set, input.WithAvailableOvfs.Set, input.WithLoadBalancers.Set)

	return err
}
