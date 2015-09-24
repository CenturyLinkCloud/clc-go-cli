package balancer

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id          string
	Name        string
	Description string
	IpAddress   string
	Status      string
	Pools       []Pool
	Links       []models.LinkEntity
}

type Pool struct {
	Id          string
	Port        int64
	Method      string
	Persistence string
	Nodes       []Node
	Links       []models.LinkEntity
}

type Node struct {
	Name        string `json:",omitempty"`
	Status      string
	IpAddress   string
	PrivatePort int64
	Links       []models.LinkEntity `json:",omitempty"`
}
