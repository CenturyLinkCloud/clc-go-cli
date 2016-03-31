package backup

import "github.com/centurylinkcloud/clc-go-cli/models/server"

type GetServerPolicies struct {
	server.Server `argument:"compose" URIParam:"ServerId"`
	WithStatus    string `URIParam:"yes" json:"-"`
}
