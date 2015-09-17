package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type GetIPAddressReq struct {
	Server   `argument:"composed" URIParam:"ServerId"`
	PublicIp string `valid:"required" URIParam:"yes"`
}

type GetIPAddressRes struct {
	InternalIPAddress  string
	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
}
