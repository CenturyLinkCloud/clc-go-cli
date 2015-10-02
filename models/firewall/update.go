package firewall

type UpdateReq struct {
	DataCenter     string `json:"-" valid:"required" URIParam:"yes"`
	FirewallPolicy string `json:"-" valid:"required" URIParam:"yes"`
	Enabled        bool
	Sources        []string `json:"Source"`
	Destinations   []string `json:"Destination"`
	Ports          []string
}

func (u *UpdateReq) Validate() error {
	err := validateSource(u.Sources)
	if err != nil {
		return err
	}
	err = validateDestination(u.Destinations)
	if err != nil {
		return err
	}
	err = validatePorts(u.Ports)
	if err != nil {
		return err
	}
	return nil
}
