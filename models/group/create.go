package group

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type CreateReq struct {
	Name            string `valid:"required"`
	Description     string
	ParentGroupId   string
	ParentGroupName string
	CustomFields    []customfields.Def
}

func (c *CreateReq) Validate() error {
	if (c.ParentGroupId == "") == (c.ParentGroupName == "") {
		return fmt.Errorf("Exactly one of the parent-group-id and parent-group-name parameters must be specified")
	}
	return nil
}

func (c *CreateReq) InferID(cn base.Connection) error {
	if c.ParentGroupName == "" {
		return nil
	}

	id, err := IDByName(cn, "all", c.ParentGroupName)
	if err != nil {
		return err
	}
	c.ParentGroupId = id
	return nil
}

func (c *CreateReq) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "ParentGroupName" {
		return nil, nil
	}
	return GetNames(cn, "all")
}
