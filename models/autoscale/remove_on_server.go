package autoscale

import (
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type RemoveOnServerReq struct {
	server.Server `argument:"composed" URIParam:"ServerId"`
}
