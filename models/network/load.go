package network

import (
	"fmt"
	"strings"
	"time"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

const (
	NETWORKS_LIST_TIMEOUT = 200
)

type loadResult struct {
	Err      error
	Networks []Entity
}

func Load(cn base.Connection, dataCenter string) ([]Entity, error) {
	var datacenters []string

	if dataCenter == "all" {
		l := []datacenter.ListRes{}
		URL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", base.URL)
		err := cn.ExecuteRequest("GET", URL, nil, &l)
		if err != nil {
			return nil, err
		}
		if len(l) == 0 {
			return nil, nil
		}
		for _, d := range l {
			datacenters = append(datacenters, d.Id)
		}
	} else {
		datacenters = []string{dataCenter}
	}

	done := make(chan loadResult)
	for _, d := range datacenters {
		go load(cn, d, done)
	}

	count := 0
	networks := []Entity{}
	for {
		select {
		case res := <-done:
			if res.Err != nil {
				return nil, res.Err
			}
			count += 1
			networks = append(networks, res.Networks...)
			if count == len(datacenters) {
				return networks, nil
			}
		case <-time.After(time.Second * time.Duration(NETWORKS_LIST_TIMEOUT)):
			return nil, fmt.Errorf("Networks list request timeout error.")
		}
	}
}

func IDByName(cn base.Connection, dataCenter string, name string) (string, error) {
	networks, err := Load(cn, dataCenter)
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, n := range networks {
		if strings.ToLower(n.Name) == strings.ToLower(name) {
			matched = append(matched, n.Id)
		}
	}

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no networks with name %s in %s.", name, dataCenter)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There are more than one network with name %s in %s", name, dataCenter)
	}
}

func GetNames(cn base.Connection, dataCenter string) ([]string, error) {
	networks, err := Load(cn, dataCenter)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, n := range networks {
		names = append(names, n.Name)
	}
	return names, nil
}

func load(cn base.Connection, dataCenter string, done chan<- loadResult) {
	networks := []Entity{}
	URL := fmt.Sprintf("%s/v2-experimental/networks/{accountAlias}/%s", base.URL, dataCenter)
	err := cn.ExecuteRequest("GET", URL, nil, &networks)
	if err != nil {
		done <- loadResult{Err: err, Networks: nil}
	}
	done <- loadResult{Err: err, Networks: networks}
}
