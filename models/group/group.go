package group

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Group struct {
	GroupId   string
	GroupName string
}

func (g *Group) Validate() error {
	if (g.GroupId == "") == (g.GroupName == "") {
		return fmt.Errorf("Exactly one of the group-id and group-name parameters must be specified")
	}
	return nil
}

func (g *Group) InferID(cn base.Connection) error {
	if g.GroupName == "" {
		return nil
	}

	groups, err := Load(cn, "all")
	if err != nil {
		return err
	}

	matched := []string{}
	var searchForID func(groups []Entity)
	searchForID = func(groups []Entity) {
		for _, group := range groups {
			if group.Name == g.GroupName {
				matched = append(matched, group.Id)
			}
			searchForID(group.Groups)
		}
	}
	searchForID(groups)

	switch len(matched) {
	case 0:
		return fmt.Errorf("There are no groups with name '%s'", g.GroupName)
	case 1:
		g.GroupId = matched[0]
		return nil
	default:
		return fmt.Errorf("There are more than one group with name '%s'. Please, specify an ID.", g.GroupName)
	}
}

func (g *Group) GetNames(cn base.Connection, name string) ([]string, error) {
	var names []string

	if name != "GroupName" {
		return nil, nil
	}

	groups, err := Load(cn, "all")
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
