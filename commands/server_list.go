package commands

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
)

type ServerList struct {
	CommandBase
}

func NewServerList(info CommandExcInfo) *ServerList {
	sl := ServerList{}
	sl.ExcInfo = info
	return &sl
}

func (sl *ServerList) Execute(cn base.Connection) error {
	var servers []models.LinkEntity

	groups, err := GetGroups(cn)
	if err != nil {
		return err
	}
	for _, g := range groups {
		err := extractServers(g, &servers)
		if err != nil {
			return err
		}
	}
	sl.Output = servers
	return nil
}

func extractServers(g group.Entity, servers *[]models.LinkEntity) error {
	if g.ServersCount == 0 {
		return nil
	}
	for _, link := range g.Links {
		if link.Rel == "server" {
			*servers = append(*servers, link)
		}
	}
	if len(g.Groups) != 0 {
		for _, gnested := range g.Groups {
			err := extractServers(gnested, servers)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
