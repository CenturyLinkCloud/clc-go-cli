package crossdc_firewall

import (
	"fmt"
	"net"
)

type CreateReq struct {
	DataCenter            string `json:"-" valid:"required" URIParam:"yes"`
	DestinationAccountId  string `valid:"required"`
	DestinationLocationId string `valid:"required"`
	DestinationCidr       string `valid:"required"`
	Enabled               string `oneOf:"true,false"`
	SourceCidr            string `valid:"required"`
}

func (c *CreateReq) Validate() error {
	var err error

	if c.Enabled == "" {
		c.Enabled = "true"
	}

	err = validateCidr(c.SourceCidr)
	if err != nil {
		return fmt.Errorf("source-cidr: %s", err.Error())
	}

	err = validateCidr(c.DestinationCidr)
	if err != nil {
		return fmt.Errorf("destination-cidr: %s", err.Error())
	}

	return nil
}

func validateCidr(cidr string) error {
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("%s is invalid CIDR", cidr)
	}

	return nil
}
