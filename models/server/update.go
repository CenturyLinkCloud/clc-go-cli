package server

import (
	"encoding/json"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
)

type UpdateReq struct {
	Server         `argument:"composed" URIParam:"ServerId"`
	PatchOperation []ServerPatchOperation `argument:"ignore"`

	Cpu          int64
	MemoryGb     int64
	RootPassword []string
	CustomFields []customfields.Def
	Description  string
	GroupId      string
	GroupName    string
	Disks        UpdateDisksDescription
}

type UpdateDisksDescription struct {
	Add  []AddDiskRequest
	Keep []KeepDiskRequest
}

func (udd *UpdateDisksDescription) Flatten() []interface{} {
	array := make([]interface{}, 0)
	add := func(el interface{}) {
		array = append(array, el)
	}
	for _, a := range udd.Add {
		add(a)
	}
	for _, k := range udd.Keep {
		add(k)
	}
	return array
}

type ServerPatchOperation struct {
	Op     string
	Member string
	Value  interface{}
}

func (ur *UpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(ur.PatchOperation)
}

func (ur *UpdateReq) Validate() error {
	err := ur.Server.Validate()
	if err != nil {
		return err
	}
	if len(ur.PatchOperation) != 0 {
		return fmt.Errorf("Invalid property: patch-operation")
	}
	if ur.Cpu < 0 || ur.MemoryGb < 0 {
		return fmt.Errorf("cpu and memory must be positive integers.")
	}
	if len(ur.RootPassword) != 0 && len(ur.RootPassword) != 2 {
		return fmt.Errorf("root-password field must consist of 2 elements - the old password and the new one.")
	}
	var any int64
	values := []int64{
		ur.Cpu,
		ur.MemoryGb,
		int64(len(ur.RootPassword)),
		int64(len(ur.CustomFields)),
		int64(len(ur.Description)),
		int64(len(ur.GroupId)),
		int64(len(ur.GroupName)),
		int64(len(ur.Disks.Add)),
		int64(len(ur.Disks.Keep)),
	}
	for _, v := range values {
		any += v
	}
	if any == 0 {
		return fmt.Errorf("At least one of the cpu, memory, root-password, custom-fields, description, group-id, group-name, disks must be provided.")
	}
	if ur.GroupId != "" && ur.GroupName != "" {
		return fmt.Errorf("Only one of group-id and group-name may be specified")
	}
	return nil
}

func (ur *UpdateReq) ApplyDefaultBehaviour() error {
	if ur.Cpu != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "cpu",
			Value:  ur.Cpu,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if ur.MemoryGb != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "memory",
			Value:  ur.MemoryGb,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if len(ur.RootPassword) != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "password",
			Value: map[string]string{
				"current":  ur.RootPassword[0],
				"password": ur.RootPassword[1],
			},
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if len(ur.CustomFields) != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "customFields",
			Value:  ur.CustomFields,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if ur.Description != "" {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "description",
			Value:  ur.Description,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if ur.GroupId != "" || ur.GroupName != "" {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "groupId",
			Value:  ur.GroupId,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if len(ur.Disks.Add) != 0 && len(ur.Disks.Keep) != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "disks",
			Value:  ur.Disks.Flatten(),
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	return nil
}

func (u *UpdateReq) InferID(cn base.Connection) error {
	err := u.Server.InferID(cn)
	if err != nil {
		return err
	}

	if u.GroupName != "" {
		var patch *ServerPatchOperation
		patched := false
		for i, op := range u.PatchOperation {
			if op.Member == "groupId" {
				patched = true
				patch = &u.PatchOperation[i]
				break
			}
		}
		if patched {
			ID, err := group.IDByName(cn, "all", u.GroupName)
			if err != nil {
				return err
			}
			patch.Value = ID
		}
	}
	return nil
}

func (u *UpdateReq) GetNames(cn base.Connection, property string) ([]string, error) {
	switch property {
	case "ServerName":
		return u.Server.GetNames(cn, "ServerName")
	case "GroupName":
		return group.GetNames(cn, "all")
	default:
		return nil, nil
	}
}
