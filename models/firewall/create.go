package firewall

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type CreateReq struct {
	SourceAccountAlias string `valid:"required" URIParam:"yes"`
	DataCenter         string `valid:"required" URIParam:"yes"`
	DestinationAccount string `valid:"required"`
	Source             []string
	Destination        []string
	Ports              []string
}

type CreateRes struct {
	Links []models.LinkEntity
}

func (c *CreateReq) Validate() error {
	err := validateSource(c.Source)
	if err != nil {
		return err
	}
	err = validateDestination(c.Destination)
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
