package datacenter

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type ListReq struct {
	WithComputeLimits base.NilField
	WithNetworkLimits base.NilField
	WithAvailableOvfs base.NilField
	WithLoadBalancers base.NilField
}

type ListRes GetRes
