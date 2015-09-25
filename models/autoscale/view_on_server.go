package autoscale

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type ViewOnServerReq struct {
	server.Server `argument:"composed" URIParam:"ServerId"`
}

type ViewOnServerRes struct {
	Id    string
	Links []models.LinkEntity
}
