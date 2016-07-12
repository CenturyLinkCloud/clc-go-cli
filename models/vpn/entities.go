package vpn

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id     string
	Local  LocalProperties
	Remote RemoteProperties
	Ike    Ike
	Ipsec  IpSec
	Links  []models.LinkEntity
}

type LocalProperties struct {
	LocationAlias       string
	LocationDescription string
	Address             string
	Subnets             []string
}

type RemoteProperties struct {
	SiteName   string
	DeviceType string
	Address    string
	Subnets    []string
}

type Ike struct {
	Encryption        string
	Hashing           string
	DiffieHelmanGroup string
	Lifetime          int64
	Mode              string
	DeadPeerDetection bool
	NatTraversal      bool
	RemoteIdentity    string
}

type IpSec struct {
	Encryption string
	Hashing    string
	Protocol   string
	Pfs        string
	Lifetime   int64
}
