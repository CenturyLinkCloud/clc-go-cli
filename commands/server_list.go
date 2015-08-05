package commands

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
	"time"
)

const (
	ServerListTimeout = 200
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
	var links []models.LinkEntity

	groups, err := GetGroups(cn)
	if err != nil {
		return err
	}
	for _, g := range groups {
		err := extractServers(g, &links)
		if err != nil {
			return err
		}
	}

	servers := make([]server.GetRes, len(links))
	done := make(chan error)
	for i, link := range links {
		go loadServer(link, servers, i, done, cn)
	}

	serversLoaded := 0
	for {
		select {
		case err := <-done:
			if err != nil {
				return err
			}
			serversLoaded += 1
			if serversLoaded == len(servers) {
				sl.Output = servers
				return nil
			}
		case <-time.After(time.Second * ServerListTimeout):
			return fmt.Errorf("Request timeout error")
		}
	}
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

func loadServer(link models.LinkEntity, servers []server.GetRes, index int, done chan<- error, cn base.Connection) {
	serverURL := fmt.Sprintf("%s%s", BaseURL, GetLink([]models.LinkEntity{link}, "server"))
	d := server.GetRes{}
	err := cn.ExecuteRequest("GET", serverURL, nil, &d)
	if err != nil {
		done <- err
		return
	}
	servers[index] = d
	done <- nil
}
