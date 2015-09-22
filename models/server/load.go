package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"time"
)

var (
	ListTimeout = 200
)

func Load(cn base.Connection, dataCenter string) ([]GetRes, error) {
	groups, err := group.Load(cn, dataCenter)
	if err != nil {
		return nil, err
	}

	links := []models.LinkEntity{}
	for _, g := range groups {
		err := extractServers(g, &links)
		if err != nil {
			return nil, err
		}
	}

	servers := make([]GetRes, len(links))
	if len(links) == 0 {
		return servers, nil
	}
	done := make(chan error)
	for i, link := range links {
		go loadServer(link, servers, i, done, cn)
	}

	serversLoaded := 0
	for {
		select {
		case err := <-done:
			if err != nil {
				return nil, err
			}
			serversLoaded += 1
			if serversLoaded == len(servers) {
				return servers, nil
			}
		case <-time.After(time.Second * time.Duration(ListTimeout)):
			return nil, fmt.Errorf("Request timeout error")
		}
	}
}

func IDByName(cn base.Connection, name string) (string, error) {
	servers, err := Load(cn, "all")
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, s := range servers {
		if s.Name == name {
			matched = append(matched, s.Id)
		}
	}

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no servers with name %s.", name)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There is more than one server with name %s. Please, specify an ID.", name)
	}
}

func GetNames(cn base.Connection, dataCenter string) ([]string, error) {
	servers, err := Load(cn, "all")
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, s := range servers {
		names = append(names, s.Name)
	}
	return names, nil
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

func loadServer(link models.LinkEntity, servers []GetRes, index int, done chan<- error, cn base.Connection) {
	href, err := models.GetLink([]models.LinkEntity{link}, "server")
	if err != nil {
		done <- err
		return
	}

	serverURL := fmt.Sprintf("%s%s", base.URL, href)
	d := GetRes{}
	err = cn.ExecuteRequest("GET", serverURL, nil, &d)
	if err != nil {
		done <- err
		return
	}

	servers[index] = d
	done <- nil
}
