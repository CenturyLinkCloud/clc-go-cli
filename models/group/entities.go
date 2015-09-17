package group

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type Entity struct {
	Id            string
	Name          string
	Description   string
	LocationId    string
	Type          string
	Status        string
	ServersCount  int64
	Groups        []Entity
	Links         []models.LinkEntity
	ChangeInfo    server.ChangeInfo
	CustromFields []server.FullCustomFieldDef
}
