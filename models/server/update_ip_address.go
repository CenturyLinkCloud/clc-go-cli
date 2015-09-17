package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type UpdateIPAddressReq struct {
	ServerId string `json:"-" valid:"required" URIParam:"yes"`
	PublicIp string `json:"-" valid:"required" URIParam:"yes"`

	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
}

func (ur *UpdateIPAddressReq) Validate() error {
	if len(ur.Ports) == 0 {
		return fmt.Errorf("Ports: non-zero value required")
	}
	return nil
}
