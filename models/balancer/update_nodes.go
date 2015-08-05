package balancer

import (
	"encoding/json"
	"fmt"
)

type UpdateNodes struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
	PoolId         string `valid:"required" URIParam:"yes"`
	Nodes          []UpdateNode
}

type UpdateNode struct {
	Status      string
	IpAddress   string `valid:"required"`
	PrivatePort int64  `valid:"required"`
}

func (u *UpdateNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Nodes)
}

func (u *UpdateNodes) Validate() error {
	if len(u.Nodes) == 0 {
		return fmt.Errorf("nodes: non-zero value required.")
	}
	return nil
}
