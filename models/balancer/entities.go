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
	Nodes       []PoolNode
	Links       []models.LinkEntity
}

type PoolNode struct {
	Name        string
	Status      string
	IpAddress   string
	PrivatePort int64
	Links       []models.LinkEntity `json:",omitempty"`
}

type Node struct {
	Status      string
	IpAddress   string
	PrivatePort int64
	Links       []models.LinkEntity `json:",omitempty"`
}
