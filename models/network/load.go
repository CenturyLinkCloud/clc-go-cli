package network

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

func Load(cn base.Connection, dataCenter string) ([]Entity, error) {
	var networks []Entity

	URL := fmt.Sprintf("%s/v2-experimental/networks/{accountAlias}/%s", base.URL, dataCenter)
	err := cn.ExecuteRequest("GET", URL, nil, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func IDByName(cn base.Connection, dataCenter string, name string) (string, error) {
	networks, err := Load(cn, dataCenter)
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, n := range networks {
		if n.Name == name {
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
