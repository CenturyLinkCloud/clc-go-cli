package server

import (
	"encoding/json"
	"fmt"
)

type PowerReq struct {
	ServerIds []string
}

func (pr *PowerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(pr.ServerIds)
}

func (pr *PowerReq) Validate() error {
	if len(pr.ServerIds) == 0 {
		return fmt.Errorf("ServerIds: non zero value required")
	}
	return nil
}
