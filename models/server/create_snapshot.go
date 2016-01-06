package server

import "github.com/centurylinkcloud/clc-go-cli/errors"

type CreateSnapshotReq struct {
	SnapshotExpirationDays int64 `valid:"required"`
	ServerIds              []string
}

func (cr *CreateSnapshotReq) Validate() error {
	if len(cr.ServerIds) == 0 {
		return errors.EmptyField("server-ids")
	}
	return nil
}
