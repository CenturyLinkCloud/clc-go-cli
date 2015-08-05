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
	err := validateActions(c.Actions)
	if err != nil {
		return err
	}
	err = validateTriggers(c.Triggers)
	if err != nil {
		return err
	}
	return nil
}

func validateActions(actions []Action) error {
	if len(actions) == 0 {
		return fmt.Errorf("Actions: non-zero value required.")
	}
	return nil
}

func validateTriggers(triggers []Trigger) error {
	if len(triggers) == 0 {
		return fmt.Errorf("Triggers: non-zero value required.")
	}
	return nil
}
