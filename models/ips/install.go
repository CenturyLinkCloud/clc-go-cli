package ips

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type Install struct {
	ServerName   string `json:"hostName" valid:"required"`
	AccountAlias string `argument:"ignore" json:"accountAlias"`
}

func (i *Install) InferID(cn base.Connection) error {
	i.AccountAlias = cn.GetAccountAlias()
	return nil
}

func (i *Install) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "ServerName" {
		return nil, nil
	}

	return server.GetNames(cn, "all")
}
