package network

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id          string
	Cidr        string
	Description string
	Gateway     string
	Name        string
	Netmask     string
	Type        string
	Vlan        int64
	Links       []models.LinkEntity
}

type IpAddress struct {
	Address string
	Claimed bool
	Server  string
	Type    string
}
