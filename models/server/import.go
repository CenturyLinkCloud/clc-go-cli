package server

import (
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type Import struct {
	Name         string             `valid:"required"`
	Description  string             `json:",omitempty"`
	GroupId      string             `valid:"required"`
	PrimaryDns   string             `json:",omitempty"`
	SecondaryDns string             `json:",omitempty"`
	NetworkId    string             `json:",omitempty"`
	RootPassword string             `valid:"required"`
	Cpu          int64              `valid:"required"`
	MemoryGb     int64              `valid:"required"`
	Type         string             `valid:"required" oneOf:"standard,hyperscale"`
	StorageType  string             `json:",omitempty" oneOf:"standard,premium,hyperscale"`
	CustomFields []customfields.Def `json:",omitempty"`
	OvfId        string             `valid:"required"`
	OvfOsType    string             `valid:"required"`
}
