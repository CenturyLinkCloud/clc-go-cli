package db

import "github.com/centurylinkcloud/clc-go-cli/errors"

type CreateNotification struct {
	SubscriptionId     string `json:"-" URIParam:"yes" valid:"required"`
	DestinationRequest `argument:"composed"`
}

func (c *CreateNotification) Validate() error {
	if c.DestinationRequest.DestinationType == "" {
		return errors.EmptyField("destination-type")
	}
	if c.DestinationRequest.Location == "" {
		return errors.EmptyField("location")
	}
	if c.DestinationRequest.Notifications == nil || len(c.DestinationRequest.Notifications) == 0 {
		return errors.EmptyField("notifications")
	}
	return nil
}
