package group

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type Entity struct {
	Id           string
	Name         string
	Description  string
	LocationId   string
	Type         string
	Status       string
	ServersCount int64
	Groups       []Entity
	Links        []models.LinkEntity
	ChangeInfo   models.ChangeInfo
	CustomFields []customfields.FullDef
}

type LBForHA struct {
	Id          string
	Name        string
	PublicPort  int64
	PrivatePort int64
	PublicIp    string
}

type HAPolicy struct {
	GroupId             string
	PolicyId            string
	LocationId          string
	Name                string
	AvailableServers    int64
	TargetSize          int64
	ScaleDirection      string
	ScaleResourceReason string
	LoadBalancer        LBForHA
	Timestamp           base.Time
}
