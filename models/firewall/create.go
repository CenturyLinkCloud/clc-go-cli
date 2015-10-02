package firewall

import (
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/models"
)

type CreateReq struct {
	DataCenter         string   `json:"-" valid:"required" URIParam:"yes"`
	DestinationAccount string   `valid:"required"`
	Sources            []string `json:"Source"`
	Destinations       []string `json:"Destination"`
	Ports              []string
}

type CreateRes struct {
	Links []models.LinkEntity
}

func (c *CreateReq) Validate() error {
	err := validateSource(c.Sources)
	if err != nil {
		return err
	}
	err = validateDestination(c.Destinations)
	if err != nil {
		return err
	}
	err = validatePorts(c.Ports)
	if err != nil {
		return err
	}
	return nil
}

func validateSource(source []string) error {
	if len(source) == 0 {
		return fmt.Errorf("source: non-zero value required.")
	}
	return nil
}

func validateDestination(dest []string) error {
	if len(dest) == 0 {
		return fmt.Errorf("destination: non-zero value required.")
	}
	return nil
}

func validatePorts(ports []string) error {
	if len(ports) == 0 {
		return fmt.Errorf("ports: non-zero value required.")
	}
	return nil
}
