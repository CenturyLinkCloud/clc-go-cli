package commands

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"time"
)

const (
	BaseURL          = "https://api.ctl.io"
	GroupListTimeout = 200
)

type GroupList struct {
	CommandBase
}

func NewGroupList(info CommandExcInfo) *GroupList {
	g := GroupList{}
	g.ExcInfo = info
	g.Input = &group.List{}
	return &g
}

func (g *GroupList) Execute(cn base.Connection) error {
	var err error
	var code string
	input := g.Input.(*group.List)

	code = input.DataCenter
	if input.All.Set {
		code = "all"
	}

	g.Output, err = GetGroups(cn, code)
	if err != nil {
		return err
	}
	return nil
}

func GetGroups(cn base.Connection, code string) ([]group.Entity, error) {
	datacenters := []datacenter.GetRes{}

	if code == "all" {
		dcURL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", BaseURL)
		err := cn.ExecuteRequest("GET", dcURL, nil, &datacenters)
		if err != nil {
			return nil, err
		}
	} else {
		dc := datacenter.GetRes{}
		dcURL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}/%s?groupLinks=true", BaseURL, code)
		err := cn.ExecuteRequest("GET", dcURL, nil, &dc)
		if err != nil {
			return nil, err
		}
		datacenters = append(datacenters, dc)
	}

	done := make(chan error)
	groups := make([]group.Entity, len(datacenters))
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
		case <-time.After(time.Second * GroupListTimeout):
			return nil, fmt.Errorf("Request timeout error.")
		}
	}
}

func GetLink(links []models.LinkEntity, resource string) string {
	for _, link := range links {
		if link.Rel == resource {
			return link.Href
		}
	}
	panic(fmt.Sprintf("No %s link found", resource))
}

func loadGroups(ref datacenter.GetRes, groups []group.Entity, dcnumber int, cn base.Connection, done chan<- error) {
	// Get detailed DC info.
	d := datacenter.GetRes{}
	dcURL := fmt.Sprintf("%s%s?groupLinks=true", BaseURL, GetLink(ref.Links, "self"))
	err := cn.ExecuteRequest("GET", dcURL, nil, &d)
	if err != nil {
		done <- err
		return
	}
	// Get the root group of the given DC.
	g := group.Entity{}
	gURL := fmt.Sprintf("%s%s", BaseURL, GetLink(d.Links, "group"))
	err = cn.ExecuteRequest("GET", gURL, nil, &g)
	if err != nil {
		done <- err
		return
	}
	groups[dcnumber] = g
	done <- nil
}
