package firewall

type UpdateReq struct {
	SourceAccountAlias string `json:"-" valid:"required" URIParam:"yes"`
	DataCenter         string `json:"-" valid:"required" URIParam:"yes"`
	FirewallPolicy     string `json:"-" valid:"required" URIParam:"yes"`
	Enabled            bool
	Source             []string
	Destination        []string
	Ports              []string
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
