package balancer

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Pool struct {
	Id          string
	Port        int64
	Method      string
	Persistence string
	Nodes       []Node
	Links       []models.LinkEntity
}

type Node struct {
	Status      string
	IpAddress   string
	PrivatePort int64
	Links       []models.LinkEntity
}
