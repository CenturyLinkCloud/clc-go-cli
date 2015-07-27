package server

import (
	"encoding/json"
	"fmt"
)

type UpdateReq struct {
	ServerId       string `valid:"required" URIParam:"true"`
	PatchOperation []ServerPatchOperation

	Cpu          int64
	Memory       int64
	Password     []string
	CustomFields []CustomFieldDef
	Description  string
	GroupId      string
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
	if len(ur.PatchOperation) != 0 {
		return fmt.Errorf("Invalid property: patch-operation")
	}
	if ur.Cpu < 0 || ur.Memory < 0 {
		return fmt.Errorf("cpu and memory must be positive integers.")
	}
	if len(ur.Password) != 0 && len(ur.Password) != 2 {
		return fmt.Errorf("password field must consist of 2 elements - the old password and the new one.")
	}
	var any int64
	values := []int64{
		ur.Cpu,
		ur.Memory,
		int64(len(ur.Password)),
		int64(len(ur.CustomFields)),
		int64(len(ur.Description)),
		int64(len(ur.GroupId)),
		int64(len(ur.Disks.Add)),
		int64(len(ur.Disks.Keep)),
	}
	for _, v := range values {
		any += v
	}
	if any == 0 {
		return fmt.Errorf("At least one of the cpu, memory, password, custom-fields, description, group-id, disks must be provided.")
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
	if ur.Memory != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "memory",
			Value:  ur.Memory,
		}
		ur.PatchOperation = append(ur.PatchOperation, op)
	}
	if len(ur.Password) != 0 {
		op := ServerPatchOperation{
			Op:     "set",
			Member: "password",
			Value: map[string]string{
				"current":  ur.Password[0],
				"password": ur.Password[1],
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
	if ur.GroupId != "" {
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
