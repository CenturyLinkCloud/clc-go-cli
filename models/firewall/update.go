package firewall

type UpdateReq struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	FirewallPolicy string `valid:"required" URIParam:"yes"`
	Enabled        bool
	Source         []string
	Destination    []string
	Ports          []string
}

func (u *UpdateReq) Validate() error {
	err := validateSource(u.Source)
	if err != nil {
		return err
	}
	err = validateDestination(u.Destination)
	if err != nil {
		return err
	}
	err = validatePorts(u.Ports)
	if err != nil {
		return err
	}
	return nil
}
