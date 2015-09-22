package group

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Group struct {
	GroupId   string `json:"-"`
	GroupName string `json:"-"`
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

	id, err := IDByName(cn, "all", g.GroupName)
	if err != nil {
		return err
	}

	g.GroupId = id
	return nil
}

func (g *Group) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "GroupName" {
		return nil, nil
	}
	return GetNames(cn, "all")
}
