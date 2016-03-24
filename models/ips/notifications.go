package ips

import (
	"encoding/json"
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/errors"
)

type SetNotifications struct {
	ServerName               string `json:"-" URIParam:"yes" valid:"required"`
	NotificationDestinations []NotificationDestination
}

func (s *SetNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.NotificationDestinations)
}

func (s *SetNotifications) Validate() error {
	if s.NotificationDestinations == nil || len(s.NotificationDestinations) == 0 {
		return errors.EmptyField("notification-destinations")
	}
	for _, n := range s.NotificationDestinations {
		if n.TypeCode == "" {
			return errors.EmptyField("notification-destinations::type-code")
		}
		if n.TypeCode != "WEBHOOK" && n.TypeCode != "EMAIL" && n.TypeCode != "SYSLOG" {
			return fmt.Errorf("type-code has to be EMAIL or WEBHOOK or SYSLOG")
		}
	}
	return nil
}

type NotificationDestination struct {
	Url            string         `json:"url"`
	TypeCode       string         `json:"typeCode"`
	EmailAddress   string         `json:"emailAddress"`
	SysLogSettings SysLogSettings `json:"sysLogSettings"`
}

type SysLogSettings struct {
	IpAddress string `json:"ipAddress"`
	UdpPort   int64  `json:"udpPort"`
	Facility  int64  `json:"facility"`
}
