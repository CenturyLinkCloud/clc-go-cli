package datacenter

type GetDCReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}

type GetDCRes struct {
	SupportsPremiumStorage     bool
	SupportsSharedLoadBalancer bool
	DeployableNetworks         []DeployableNetwork
	Templates                  []Template
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
