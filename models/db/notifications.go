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

type UpdateNotification struct {
	DestinationId      string `json:"-" URIParam:"yes" valid:"required"`
	SubscriptionId     string `json:"-" URIParam:"yes" valid:"required"`
	DestinationRequest `argument:"composed"`
}

func (c *UpdateNotification) Validate() error {
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

type DeleteNotification struct {
	DestinationId  string `json:"-" URIParam:"yes" valid:"required"`
	SubscriptionId string `json:"-" URIParam:"yes" valid:"required"`
}

type VerifyDestination struct {
	DestinationId  string `json:"-" URIParam:"yes" valid:"required"`
	SubscriptionId string `json:"-" URIParam:"yes" valid:"required"`
	Token          string `json:"-" URIParam:"yes" valid:"required"`
}
