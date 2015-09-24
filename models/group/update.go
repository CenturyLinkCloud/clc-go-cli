package group

import (
	"encoding/json"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type UpdateReq struct {
	Group          `argument:"composed" URIParam:"GroupId" json:"-"`
	PatchOperation []GroupPatchOperation `argument:"ignore"`

	CustomFields    []customfields.Def
	Name            string
	Description     string
	ParentGroupId   string
	ParentGroupName string
}

type GroupPatchOperation struct {
	Op     string
	Member string
	Value  interface{}
}

func (u *UpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.PatchOperation)
}

func (u *UpdateReq) Validate() error {
	if len(u.PatchOperation) != 0 {
		return fmt.Errorf("Invalid property: patch-operation")
	}

	if err := u.Group.Validate(); err != nil {
		return err
	}

	if u.ParentGroupName != "" && u.ParentGroupId != "" {
		return fmt.Errorf("Only one of parent-group-id and parent-group-name may be specified")
	}

	var any int64
	values := []int64{
		int64(len(u.CustomFields)),
		int64(len(u.Name)),
		int64(len(u.Description)),
		int64(len(u.ParentGroupId)),
		int64(len(u.ParentGroupName)),
	}
	for _, v := range values {
		any += v
	}
	if any == 0 {
		return fmt.Errorf("At least one of the custom-fields, name, description, parent-group-id, parent-group-name must be provided.")
	}
	return nil
}

func (u *UpdateReq) ApplyDefaultBehaviour() error {
	if len(u.CustomFields) != 0 {
		op := GroupPatchOperation{
			Op:     "set",
			Member: "customFields",
			Value:  u.CustomFields,
		}
		u.PatchOperation = append(u.PatchOperation, op)
	}
	if u.Name != "" {
		op := GroupPatchOperation{
			Op:     "set",
			Member: "name",
			Value:  u.Name,
		}
		u.PatchOperation = append(u.PatchOperation, op)
	}
	if u.Description != "" {
		op := GroupPatchOperation{
			Op:     "set",
			Member: "description",
			Value:  u.Description,
		}
		u.PatchOperation = append(u.PatchOperation, op)
	}
	if u.ParentGroupId != "" {
		op := GroupPatchOperation{
			Op:     "set",
			Member: "parentGroupId",
			Value:  u.ParentGroupId,
		}
		u.PatchOperation = append(u.PatchOperation, op)
	}
	return nil
}

func (u *UpdateReq) InferID(cn base.Connection) error {
	err := u.Group.InferID(cn)
	if err != nil {
		return err
	}

	if u.ParentGroupName == "" {
		return nil
	}

	id, err := IDByName(cn, "all", u.ParentGroupName)
	if err != nil {
		return err
	}
	u.ParentGroupId = id
	return nil
}

func (u *UpdateReq) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "ParentGroupName" && property != "GroupName" {
		return nil, nil
	}

	// Group and Parent Group need the same list of names.
	return GetNames(cn, "all")
}
