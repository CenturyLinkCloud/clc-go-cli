package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AddIPAddressReq struct {
	Server `json:"-" argument:"composed" URIParam:"ServerId"`

	InternalIpAddress  string                     `json:",omitempty"`
	Ports              []models.PortRestriction   `json:",omitempty"`
	SourceRestrictions []models.SourceRestriction `json:",omitempty"`
}

func (ar *AddIPAddressReq) Validate() error {
	err := ar.Server.Validate()
	if err != nil {
		return err
	}
	if len(ar.Ports) == 0 {
		return fmt.Errorf("Ports: non-zero value required.")
	}
	return nil
}
