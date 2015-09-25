package autoscale

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type SetOnServerReq struct {
	server.Server `argument:"composed" URIParam:"ServerId"`
	PolicyId      string `json:"Id"`
	PolicyName    string
}

type SetOnServerRes struct {
	Id    string
	Links []models.LinkEntity
}

func (p *SetOnServerReq) Validate() error {
	return validatePolicy(p.PolicyId, p.PolicyName)
}

func (p *SetOnServerReq) InferID(cn base.Connection) error {
	id, err := inferID(p.PolicyName, cn)
	if err == nil && id != "" {
		p.PolicyId = id
	}
	return err
}

func (p *SetOnServerReq) GetNames(cn base.Connection, property string) ([]string, error) {
	return getNames(cn, property)
}
