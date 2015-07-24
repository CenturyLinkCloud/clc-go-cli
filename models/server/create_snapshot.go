package server

import (
	"fmt"
)

type CreateSnapshotReq struct {
	SnapshotExpirationDays int64 `valid:"required"`
	ServerIds              []string
}

func (cr *CreateSnapshotReq) Validate() error {
	if len(cr.ServerIds) == 0 {
		return fmt.Errorf("ServerIds: non zero value required")
	}
	return nil
}
