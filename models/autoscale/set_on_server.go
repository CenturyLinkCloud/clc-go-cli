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
	err := p.Server.Validate()
	if err != nil {
		return err
	}
	return validatePolicy(p.PolicyId, p.PolicyName)
}

func (p *SetOnServerReq) InferID(cn base.Connection) error {
	err := p.Server.InferID(cn)
	if err != nil {
		return err
	}
	id, err := inferID(p.PolicyName, cn)
	if err == nil && id != "" {
		p.PolicyId = id
	}
	return err
}

func (p *SetOnServerReq) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "ServerName" {
		return p.Server.GetNames(cn, property)
	}
	return getNames(cn, property)
}
