package webhook

// DeleteReq represents a request to delete all the target uris associated with a webhook for a given event
type DeleteReq struct {
	Event string `json:"-" valid:"required" URIParam:"yes"`
}

// Validate provides custom validation for the AddTargetURIReq request
func (c *DeleteReq) Validate() error {

	err := ValidateEvent(c.Event)
	if err != nil {
		return err
	}

	return nil
}
