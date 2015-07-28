package commands

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
)

const (
	BaseURL = "https://api.ctl.io"
)

type GroupList struct {
	CommandBase
}

func NewGroupList(info CommandExcInfo) *GroupList {
	g := GroupList{}
	g.ExcInfo = info
	return &g
}

func (g *GroupList) Execute(cn base.Connection) error {
	var err error

	g.Output, err = Get(cn)
	if err != nil {
		return err
	}
	return nil
}

func Get(cn base.Connection) ([]group.Entity, error) {
	var err error
	var groups []group.Entity

	datacenters := []datacenter.GetRes{}
	dcURL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", BaseURL)
	err = cn.ExecuteRequest("GET", dcURL, nil, &datacenters)
	if err != nil {
		return nil, err
	}

	for _, ref := range datacenters {
		// Get detailed DC info.
		d := datacenter.GetRes{}
		dcURL = fmt.Sprintf("%s/%s?groupLinks=true", BaseURL, GetLink(ref.Links, "self"))
		err = cn.ExecuteRequest("GET", dcURL, nil, &d)
		if err != nil {
			return nil, err
		}
		// Get the root group of the given DC.
		g := group.Entity{}
		gURL := fmt.Sprintf("%s/%s", BaseURL, GetLink(d.Links, "group"))
		err = cn.ExecuteRequest("GET", gURL, nil, &g)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	return groups, nil
}

func GetLink(links []models.LinkEntity, resource string) string {
	for _, link := range links {
		if link.Rel == resource {
			return link.Href
		}
	}
	panic(fmt.Sprintf("No %s link found", resource))
}
