package webhook

// AddTargetURIReq represents a request to add a webhook. It will add a target uri to be called on a specific event.
type AddTargetURIReq struct {
	Event     string `json:"-" valid:"required" URIParam:"yes"`
	TargetUri string `valid:"required" json:"targetUri"`
}

// Validate provides custom validation for the AddTargetURIReq request
func (c *AddTargetURIReq) Validate() error {

	err := ValidateTargetURI(c.TargetUri)
	if err != nil {
		return err
	}

	err = ValidateEvent(c.Event)
	if err != nil {
		return err
	}

	return nil
}
