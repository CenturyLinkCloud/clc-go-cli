package group

import (
	"fmt"
	"strings"
	"time"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
)

var (
	ListTimeout = 200
)

func Load(cn base.Connection, dataCenter string) ([]Entity, error) {
	datacenters := []datacenter.GetRes{}

	if dataCenter == "all" {
		dcURL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", base.URL)
		err := cn.ExecuteRequest("GET", dcURL, nil, &datacenters)
		if err != nil {
			return nil, err
		}
	} else {
		dc := datacenter.GetRes{}
		dcURL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}/%s?groupLinks=true", base.URL, dataCenter)
		err := cn.ExecuteRequest("GET", dcURL, nil, &dc)
		if err != nil {
			return nil, err
		}
		datacenters = append(datacenters, dc)
	}

	done := make(chan error)
	groups := make([]Entity, len(datacenters))
	for i, ref := range datacenters {
		go loadGroups(ref, groups, i, cn, done)
	}

	received := 0
	for {
		select {
		case err := <-done:
			if err != nil {
				return nil, err
			}
			received += 1
			if received == len(datacenters) {
				return groups, nil
			}
		case <-time.After(time.Second * time.Duration(ListTimeout)):
			return nil, fmt.Errorf("Request timeout error.")
		}
	}
}

func IDByName(cn base.Connection, dataCenter string, name string) (string, error) {
	groups, err := Load(cn, dataCenter)
	if err != nil {
		return "", err
	}

	matched := []string{}
	var searchForID func(groups []Entity)
	searchForID = func(groups []Entity) {
		for _, group := range groups {
			if strings.ToLower(group.Name) == strings.ToLower(name) {
				matched = append(matched, group.Id)
			}
			searchForID(group.Groups)
		}
	}
	searchForID(groups)

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no groups with name '%s'", name)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There are more than one group with name '%s'. Please, specify an ID.", name)
	}
}

func GetNames(cn base.Connection, dataCenter string) ([]string, error) {
	var names []string
	groups, err := Load(cn, dataCenter)
	if err != nil {
		return nil, err
	}

	var collectNames func(groups []Entity)
	collectNames = func(groups []Entity) {
		for _, group := range groups {
			names = append(names, group.Name)
			collectNames(group.Groups)
		}
	}
	collectNames(groups)

	return names, nil
}

func loadGroups(ref datacenter.GetRes, groups []Entity, dcnumber int, cn base.Connection, done chan<- error) {
	// Get detailed DC info.
	link, err := models.GetLink(ref.Links, "self")
	if err != nil {
		done <- err
		return
	}
	d := datacenter.GetRes{}
	dcURL := fmt.Sprintf("%s%s?groupLinks=true", base.URL, link)
	err = cn.ExecuteRequest("GET", dcURL, nil, &d)
	if err != nil {
		done <- err
		return
	}
	// Get the root group of the given DC.
	link, err = models.GetLink(d.Links, "group")
	if err != nil {
		done <- err
		return
	}
	g := Entity{}
	gURL := fmt.Sprintf("%s%s", base.URL, link)
	err = cn.ExecuteRequest("GET", gURL, nil, &g)
	if err != nil {
		done <- err
		return
	}
	groups[dcnumber] = g
	done <- nil
}
