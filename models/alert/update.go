package alert

type UpdateReq struct {
	Policy   `argument:"composed" URIParam:"PolicyId" json:"-"`
	Name     string `valid:"required"`
	Actions  []Action
	Triggers []Trigger
}

func (u *UpdateReq) Validate() error {
	err := u.Policy.Validate()
	if err != nil {
		return err
	}

	err = validateActions(u.Actions)
	if err != nil {
		return err
	}
	err = validateTriggers(u.Triggers)
	if err != nil {
		return err
	}
	return nil
}
