package server

import (
	"encoding/json"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
)

type Disk struct {
	Id             string
	SizeGB         float64
	PartitionPaths []string
}

type Partition struct {
	SizeGB float64
	Path   string
}

type AddDiskRequest struct {
	Path   string
	SizeGB int64
	Type   string
}

type KeepDiskRequest struct {
	DiskId string
	SizeGB int64
}

type PackageDef struct {
	PackageId  string
	Parameters map[string]string
}

type Details struct {
	IpAddresses          []IPAddresses
	SecondaryIPAddresses []IPAddresses
	AlertPolicies        []alert.AlertPolicy
	Cpu                  int64
	DiskCount            int64
	HostName             string
	InMaintenanceMode    bool
	MemoryMB             int64
	PowerState           string
	StorageGB            int64
	Disks                []Disk
	Partitions           []Partition
	Snapshots            []Snapshot
	CustomFields         []customfields.FullDef
	ProcessorDescription string `json:",omitempty"`
	StorageDescription   string `json:",omitempty"`
	IsManagedBackup      bool   `json:",omitempty"`
}

type IPAddresses struct {
	Public   string `json:",omitempty"`
	Internal string
}

type Snapshot struct {
	Name  string
	Id    string
	Links []models.LinkEntity
}

func (s *Snapshot) UnmarshalJSON(data []byte) error {
	// Here we are getting the snapshot ID out of the server links. Looks
	// rather hacky but there doesn't seem to be another way for retrieving it.
	id := func() string {
		m := struct {
			Links []struct {
				Rel  string
				Href string
			}
		}{}
		if err := json.Unmarshal(data, &m); err != nil {
			return ""
		}
		for _, link := range m.Links {
			if link.Rel != "self" {
				continue
			}
			tokens := strings.Split(link.Href, "/")
			return tokens[len(tokens)-1]
		}
		return ""
	}()

	if id != "" {
		s.Id = id
	}

	type SimpleSnapshot Snapshot
	return json.Unmarshal(data, &struct {
		*SimpleSnapshot
	}{
		SimpleSnapshot: (*SimpleSnapshot)(s),
	})
}
