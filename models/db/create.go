package db

import (
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/errors"
)

type Create struct {
	InstanceType        string               `valid:"required" oneOf:"MySQL,MySQL_REPLICATION" json:"instanceType"`
	ExternalId          string               `valid:"required" json:"externalId"`
	MachineConfig       MachineConfig        `json:"machineConfig"`
	BackupRetentionDays int64                `valid:"required" json:"backupRetentionDays"`
	Users               []User               `json:"users"`
	DataCenter          string               `json:"location"`
	Destinations        []DestinationRequest `json:"destinations,omitempty"`
	Instances           []Instance           `json:"instances,omitempty"`
	BackupTime          BackupTime           `json:"-"`
	BackupTimeValue     *BackupTime          `json:"backupTime,omitempty" argument:ignore`
}

func (c *Create) Validate() error {
	if c.Users == nil || len(c.Users) == 0 {
		return errors.EmptyField("--users")
	}
	emptyConfig := MachineConfig{}
	if c.MachineConfig == emptyConfig {
		return errors.EmptyField("--machine-config")
	}

	m := c.MachineConfig
	if m.Cpu == nil || m.Memory == nil || m.Storage == nil {
		return fmt.Errorf("All of the cpu, memory, and storage have to be set")
	}

	emptyBackup := BackupTime{}
	if c.BackupTime != emptyBackup {
		b := c.BackupTime
		if b.Hour == nil || b.Minute == nil {
			return fmt.Errorf("Both hour and minute have to be set")
		}
		c.BackupTimeValue = &b
	}
	return nil
}

type MachineConfig struct {
	Cpu     *int64 `json:"cpu,omitempty"`
	Memory  *int64 `json:"memory,omitempty"`
	Storage *int64 `json:"storage,omitempty"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DestinationRequest struct {
	DestinationType string                `oneOf:"EMAIL,SMS" json:"destinationType"`
	Location        string                `json:"location"`
	Notifications   []NotificationRequest `json:"notifications"`
}

type NotificationRequest struct {
	NotificationType string `oneOf:"CPU_UTILIATION,MEMORY_UTILIZATION,STORAGE_UTILIZATION" json:"notificationType"`
}

type Instance struct {
	Name string `json:"name"`
}

type BackupTime struct {
	Hour   *int64 `json:"hour,omitempty"`
	Minute *int64 `json:"minute,omitempty"`
}

type CreateRes struct {
	Id                  int64
	Location            string
	InstanceType        string
	ExternalId          string
	Status              string
	BackupTime          string
	BackupRetentionDays int64
	OptionGroup         string
	ParameterGroup      string
	Users               []User
	Instances           []Instance
	Servers             []Server
	Host                string
	Port                int64
	Certificate         string
	Backups             []BackupResponse
}

type Server struct {
	Id          int64
	Alias       string
	Location    string
	Cpu         int64
	Memory      int64
	Storage     int64
	Attributes  map[string]string
	Connections int64
}

type BackupResponse struct {
	Id         int64
	FileName   string
	BackupTime int64
	Status     string
	Size       int64
}
