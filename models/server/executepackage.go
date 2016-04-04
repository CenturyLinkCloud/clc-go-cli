package server

import "github.com/centurylinkcloud/clc-go-cli/errors"

type ExecutePackage struct {
	ServerIds []string   `json:"servers"`
	Package   PackageDef `valid:"required"`
}

func (e *ExecutePackage) Validate() error {
	if len(e.ServerIds) == 0 {
		return errors.EmptyField("server-ids")
	}
	return nil
}
