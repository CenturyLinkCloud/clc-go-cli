package group

import (
	"encoding/json"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

type UpdateReq struct {
	GroupId        string                `valid:"required" URIParam:"yes"`
	PatchOperation []GroupPatchOperation `argument:"ignore"`

	CustomFields  []server.CustomFieldDef
	Name          string
	Description   string
	ParentGroupId string
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

	var any int64
	values := []int64{
		int64(len(u.CustomFields)),
		int64(len(u.Name)),
		int64(len(u.Description)),
		int64(len(u.ParentGroupId)),
	}
	for _, v := range values {
		any += v
	}
	if any == 0 {
		return fmt.Errorf("At least one of the custom-fields, name, description, parent-group-id must be provided.")
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
