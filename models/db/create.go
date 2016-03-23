package db

import "github.com/centurylinkcloud/clc-go-cli/errors"

type Create struct {
	InstanceType        string               `valid:"required" oneOf:"MySQL,MySQL_REPLICATION" json:"instanceType"`
	ExternalId          string               `valid:"required" json:"externalId"`
	MachineConfig       MachineConfig        `json:"machineConfig"`
	BackupRetentionDays int64                `valid:"required" json:"backupRetentionDays"`
	Users               []User               `json:"users"`
	DataCenter          string               `json:"location"`
	Destinations        []DestinationRequest `json:"destinations,omitempty"`
	Instances           []Instance           `json:"instances,omitempty"`
	BackupTime          BackupTime           `json:"backupTime"`
}

func (c *Create) Validate() error {
	if c.Users == nil || len(c.Users) == 0 {
		return errors.EmptyField("--users")
	}
	emptyConfig := MachineConfig{}
	if c.MachineConfig == emptyConfig {
		return errors.EmptyField("--machine-config")
	}
	return nil
}

type MachineConfig struct {
	Cpu     int64 `json:"cpu"`
	Memory  int64 `json:"memory"`
	Storage int64 `json:"storage"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DestinationRequest struct {
	DestinationType string                `oneOf:"EMAIL,SMS" json:"destinationType"`
	DataCenter      string                `json:"location"`
	Notifications   []NotificationRequest `json:"notifications"`
}

type NotificationRequest struct {
	NotificationType string `oneOf:"CPU_UTILIATION,MEMORY_UTILIZATION,STORAGE_UTILIZATION" json:"notificationType"`
}

type Instance struct {
	Name string `json:"name"`
}

type BackupTime struct {
	Hour   int64 `json:"hour"`
	Minute int64 `json:"minute"`
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
	BackupTime string
	Status     string
	Size       int64
}
