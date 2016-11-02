package webhook

// DeleteTargetURIReq represents a request to delete a specific target uri associated with a webhook for a given event
type DeleteTargetURIReq struct {
	Event     string `json:"-" valid:"required" URIParam:"yes"`
	TargetUri string `URIParam:"yes" valid:"required"`
}

// Validate provides custom validation for the DeleteTargetURIReq request
func (c *DeleteTargetURIReq) Validate() error {

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
