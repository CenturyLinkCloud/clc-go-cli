package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
)

type Import struct {
	Name         string `valid:"required"`
	Description  string `json:",omitempty"`
	GroupId      string
	GroupName    string
	PrimaryDns   string             `json:",omitempty"`
	SecondaryDns string             `json:",omitempty"`
	NetworkId    string             `json:",omitempty"`
	RootPassword string             `valid:"required"`
	Cpu          int64              `valid:"required"`
	MemoryGb     int64              `valid:"required"`
	Type         string             `valid:"required" oneOf:"standard,hyperscale"`
	StorageType  string             `json:",omitempty" oneOf:"standard,premium,hyperscale"`
	CustomFields []customfields.Def `json:",omitempty"`
	OvfId        string             `valid:"required"`
	OvfOsType    string             `valid:"required"`
}

func (i *Import) Validate() error {
	if (i.GroupId == "") == (i.GroupName == "") {
		return fmt.Errorf("Exactly one of the parameters group-id and group-name must be specified")
	}
	return nil
}

func (i *Import) InferID(cn base.Connection) error {
	if i.GroupName != "" {
		g := &group.Group{GroupName: i.GroupName}
		err := g.InferID(cn)
		if err != nil {
			return err
		}
		i.GroupId = g.GroupId
	}
	return nil
}

func (i *Import) GetNames(cn base.Connection, property string) ([]string, error) {
	switch property {
	case "GroupName":
		return group.GetNames(cn, "all")
	default:
		return nil, nil
	}
}
