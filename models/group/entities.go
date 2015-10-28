package group

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type Entity struct {
	Id           string
	Name         string
	Description  string
	LocationId   string
	Type         string
	Status       string
	ServersCount int64
	Groups       []Entity
	Links        []models.LinkEntity
	ChangeInfo   models.ChangeInfo
	CustomFields []customfields.FullDef
}

type LBForHA struct {
	Id          string
	Name        string
	PublicPort  int64
	PrivatePort int64
	PublicIp    string
}

type HAPolicy struct {
	GroupId             string
	PolicyId            string
	LocationId          string
	Name                string
	AvailableServers    int64
	TargetSize          int64
	ScaleDirection      string
	ScaleResourceReason string
	LoadBalancer        LBForHA
	Timestamp           base.Time
}

type ScheduledActivities struct {
	Id                    string
	LocationId            string
	ChangeInfo            models.ChangeInfo
	Status                string
	Type                  string
	BeginDateUTC          base.Time
	Repeat                string
	CustomWeeklyDays      []string
	Expire                string
	ExpireCount           int64
	ExpireDateUTC         base.Time
	TimeZoneOffset        string
	IsExpired             bool
	LastOccurrenceDateUTC base.Time
	OccurrenceCount       int64
	NextOccurrenceDateUTC base.Time
}

type CPU struct {
	Value     int64
	Inherited bool
}

type Memory struct {
	Value     int64
	Inherited bool
}

type Network struct {
	Value     string
	Inherited bool
}

type PrimaryDNS struct {
	Value     string
	Inherited bool
}

type SecondaryDNS struct {
	Value     string
	Inherited bool
}

type TemplateName struct {
	Value     string
	Inherited bool
}

type Defaults struct {
	Cpu          CPU
	MemoryGB     Memory
	NetworkId    Network
	PrimaryDns   PrimaryDNS
	SecondaryDns SecondaryDNS
	TemplateName TemplateName
}
