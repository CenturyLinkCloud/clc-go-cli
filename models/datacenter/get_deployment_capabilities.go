package datacenter

type GetDCReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}

type GetDCRes struct {
	SupportsSharedLoadBalancer bool
	SupportsBareMetalServers   bool
	DeployableNetworks         []DeployableNetwork
	Templates                  []Template
	ImportableOsTypes          []ImportableOSType
}

type DeployableNetwork struct {
	Name      string
	NetworkId string
	Type      string
	AccountId string
}

type Template struct {
	Name               string
	Description        string
	StorageSizeGB      int64
	Capabilities       []string
	ReservedDrivePaths []string
	DrivePathLength    int64
}

type ImportableOSType struct {
	Id                 int64
	Description        string
	LabProductCode     string
	PremiumProductCode string
	Type               string
}
