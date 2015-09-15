package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type UpdateIPAddressReq struct {
	Server   `argument:"composed" URIParam:"ServerId"`
	PublicIp string `valid:"required" URIParam:"yes"`

	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
}

func (ur *UpdateIPAddressReq) Validate() error {
	err := ur.Server.Validate()
	if err != nil {
		return err
	}
	if len(ur.Ports) == 0 {
		return fmt.Errorf("Ports: non-zero value required")
	}
	return nil
}
