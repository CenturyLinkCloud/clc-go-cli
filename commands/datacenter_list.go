package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

type DatacenterList struct {
	CommandBase
}

func NewDatacenterList(info CommandExcInfo) *DatacenterList {
	dcList := DatacenterList{}
	dcList.ExcInfo = info
	dcList.Input = &datacenter.ListReq{}
	return &dcList
}

func (dcList *DatacenterList) Execute(cn base.Connection) error {
	var err error

	input := dcList.Input.(*datacenter.ListReq)
	dcList.Output, err = datacenter.All(cn, input.WithComputeLimits.Set,
		input.WithNetworkLimits.Set, input.WithAvailableOvfs.Set, input.WithLoadBalancers.Set)

	return err
}
