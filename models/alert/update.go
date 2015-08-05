package alert

type UpdateReq struct {
	PolicyId string `valid:"required" URIParam:"yes"`
	Name     string `valid:"required"`
	Actions  []Action
	Triggers []Trigger
}

func (u *UpdateReq) Validate() error {
	err := validateActions(u.Actions)
	if err != nil {
		return err
	}
	err = validateTriggers(u.Triggers)
	if err != nil {
		return err
	}
	return nil
}
