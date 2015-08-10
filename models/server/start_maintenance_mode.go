package server

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StartMaintenance struct {
	ServerIds []string
}

func (s *StartMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ServerIds)
}

func (s *StartMaintenance) Validate() error {
	if len(s.ServerIds) == 0 {
		return fmt.Errorf("ServerIds: non-zero value required.")
	}
	return nil
}

// This method replaces all of the ids with their upper case versions because
// the API call only accepts ids written in upper case.
func (s *StartMaintenance) ApplyDefaultBehaviour() error {
	for i, id := range s.ServerIds {
		s.ServerIds[i] = strings.ToUpper(id)
	}
	return nil
}
