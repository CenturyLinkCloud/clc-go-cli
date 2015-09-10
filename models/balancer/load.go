package balancer

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

func Load(cn base.Connection, dataCenter string) ([]Entity, error) {
	var balancers []Entity

	URL := fmt.Sprintf("%s/v2/sharedLoadBalancers/{accountAlias}/%s", base.URL, dataCenter)
	err := cn.ExecuteRequest("GET", URL, nil, &balancers)
	if err != nil {
		return nil, err
	}
	return balancers, nil
}

func IDByName(cn base.Connection, dataCenter string, name string) (string, error) {
	balancers, err := Load(cn, dataCenter)
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, b := range balancers {
		if b.Name == name {
			matched = append(matched, b.Id)
		}
	}

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no load balancers with name %s in %s.", name, dataCenter)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There are more than one balancer with name %s in %s. Please, specify and ID.", name, dataCenter)
	}
}

func GetNames(cn base.Connection, dataCenter string) ([]string, error) {
	balancers, err := Load(cn, dataCenter)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, b := range balancers {
		names = append(names, b.Name)
	}
	return names, nil
}
