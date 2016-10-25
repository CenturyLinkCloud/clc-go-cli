package webhook

// UpdateReq represents a request to update a webhook for a given event
type UpdateReq struct {
	Event      string      `json:"-" valid:"required" URIParam:"yes"`
	Recursive  string      `json:"recursive" oneOf:"true,false"`
	TargetUris []TargetUri `json:"targetUris,omitempty"`
}

// Validate provides custom validation for the UpdateReq request
func (c *UpdateReq) Validate() error {

	if c.TargetUris != nil {
		if len(c.TargetUris) > 0 {

			for _, targetURI := range c.TargetUris {
				err := ValidateTargetURI(targetURI.TargetUri)
				if err != nil {
					return err
				}
			}
		}
	}

	err := ValidateEvent(c.Event)
	if err != nil {
		return err
	}

	return nil
}
