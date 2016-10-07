package server

import (
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/network"
)

type AddNetwork struct {
	ServerId    string `json:"-" URIParam:"yes"`
	ServerName  string `json:"-"`
	NetworkId   string
	NetworkName string `json:"-"`
	IpAddress   string
}

type RemoveNetwork struct {
	ServerId    string `URIParam:"yes" json:"-"`
	ServerName  string `json:"-"`
	NetworkId   string `URIParam:"yes" json:"-"`
	NetworkName string `json:"-"`
}

func (a *AddNetwork) Validate() error {
	if (a.ServerId == "") == (a.ServerName == "") {
		return fmt.Errorf("Exactly one of the parameters server-id and server-name must be specified.")
	}

	if (a.NetworkId == "") == (a.NetworkName == "") {
		return fmt.Errorf("Exactly one of the parameters network-id and network-name must be specified.")
	}
	return nil
}

func (a *AddNetwork) InferID(cn base.Connection) error {
	if a.ServerName != "" {
		s := &Server{ServerName: a.ServerName}
		if err := s.InferID(cn); err != nil {
			return err
		}
		a.ServerId = s.ServerId
	}

	if a.NetworkName != "" {
		ID, err := network.IDByName(cn, "all", a.NetworkName)
		if err != nil {
			return err
		}
		a.NetworkId = ID
	}
	return nil
}

func (a *AddNetwork) GetNames(cn base.Connection, property string) ([]string, error) {
	switch property {
	case "ServerName":
		return GetNames(cn, "all")
	case "NetworkName":
		return network.GetNames(cn, "all")
	default:
		return nil, nil
	}
}

func (r *RemoveNetwork) Validate() error {
	if (r.ServerId == "") == (r.ServerName == "") {
		return fmt.Errorf("Exactly one of the parameters server-id and server-name must be specified.")
	}

	if (r.NetworkId == "") == (r.NetworkName == "") {
		return fmt.Errorf("Exactly one of the parameters network-id and network-name must be specified.")
	}
	return nil
}

func (r *RemoveNetwork) InferID(cn base.Connection) error {
	if r.ServerName != "" {
		s := &Server{ServerName: r.ServerName}
		if err := s.InferID(cn); err != nil {
			return err
		}
		r.ServerId = s.ServerId
	}

	if r.NetworkName != "" {
		ID, err := network.IDByName(cn, "all", r.NetworkName)
		if err != nil {
			return err
		}
		r.NetworkId = ID
	}
	return nil
}

func (r *RemoveNetwork) GetNames(cn base.Connection, property string) ([]string, error) {
	switch property {
	case "ServerName":
		return GetNames(cn, "all")
	case "NetworkName":
		return network.GetNames(cn, "all")
	default:
		return nil, nil
	}
}
