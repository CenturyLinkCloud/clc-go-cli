package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AddIPAddressReq struct {
	ServerId string `valid:"required" URIParam:"true"`

	InternalIpAddress  string
	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
}

func (ar *AddIPAddressReq) Validate() error {
	if len(ar.Ports) == 0 {
		return fmt.Errorf("Ports: non-zero value required.")
	}
	return nil
}
