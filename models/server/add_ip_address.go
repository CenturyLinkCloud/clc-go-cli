package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AddIPAddressReq struct {
	ServerId string `json:"-" valid:"required" URIParam:"true"`

	InternalIpAddress  string                     `json:",omitempty"`
	Ports              []models.PortRestriction   `json:",omitempty"`
	SourceRestrictions []models.SourceRestriction `json:",omitempty"`
}

func (ar *AddIPAddressReq) Validate() error {
	if len(ar.Ports) == 0 {
		return fmt.Errorf("Ports: non-zero value required.")
	}
	return nil
}
