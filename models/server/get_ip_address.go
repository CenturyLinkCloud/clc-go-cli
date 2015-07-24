package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type GetIPAddressReq struct {
	ServerId string `valid:"required" URIParam:"yes"`
	PublicIp string `valid:"required" URIParam:"yes"`
}

type GetIPAddressRes struct {
	InternalIPAddress  string
	Ports              []models.PortRestriction
	SourceRestrictions []models.SourceRestriction
}
