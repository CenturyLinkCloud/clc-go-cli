package server

import (
	"encoding/json"

	"github.com/centurylinkcloud/clc-go-cli/errors"
)

type PowerReq struct {
	ServerIds []string
}

func (pr *PowerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(pr.ServerIds)
}

func (pr *PowerReq) Validate() error {
	if len(pr.ServerIds) == 0 {
		return errors.EmptyField("server-ids")
	}
	return nil
}
