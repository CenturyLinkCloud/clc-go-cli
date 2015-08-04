package alert

import (
	"fmt"
)

type CreateReq struct {
	Name     string `valid:"required"`
	Actions  []Action
	Triggers []Trigger
}

func (c *CreateReq) Validate() error {
	if len(c.Actions) == 0 {
		return fmt.Errorf("Actions: non-zero value required.")
	} else if len(c.Triggers) == 0 {
		return fmt.Errorf("Triggers: non-zero value required.")
	}
	return nil
}
