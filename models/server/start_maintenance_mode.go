package server

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MaintenanceRequest struct {
	ServerIds []string
}

func (s *MaintenanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ServerIds)
}

func (s *MaintenanceRequest) Validate() error {
	if len(s.ServerIds) == 0 {
		return fmt.Errorf("ServerIds: non-zero value required.")
	}
	return nil
}

// This method replaces all of the ids with their upper case versions because
// the API call only accepts ids written in upper case.
func (s *MaintenanceRequest) ApplyDefaultBehaviour() error {
	for i, id := range s.ServerIds {
		s.ServerIds[i] = strings.ToUpper(id)
	}
	return nil
}
