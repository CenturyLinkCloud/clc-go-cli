package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type AddIPAddressReq struct {
	Server `argument:"composed" URIParam:"ServerId"`

	InternalIpAddress  string
	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
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
