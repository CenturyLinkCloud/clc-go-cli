package network

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Network struct {
	NetworkId   string
	NetworkName string
	DataCenter  string
}

func (n *Network) Validate() error {
	if n.DataCenter == "" {
		return fmt.Errorf("DataCenter: non-zero value required.")
	}

	if (n.NetworkId == "") == (n.NetworkName == "") {
		return fmt.Errorf("Exactly one of the network-id and network-name properties must be specified.")
	}
	return nil
}

func (n *Network) InferID(cn base.Connection) error {
	if n.NetworkName == "" {
		return nil
	}

	id, err := IDByName(cn, n.DataCenter, n.NetworkName)
	if err != nil {
		return err
	}
	n.NetworkId = id
	return nil
}

func (n *Network) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "NetworkName" {
		return nil, nil
	}

	return GetNames(cn, n.DataCenter)
}
