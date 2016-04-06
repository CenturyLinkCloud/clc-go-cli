package backup

type DataCenters []string

type OSTypes []string

type GetServers struct {
	DataCenterName string `valid:"required" URIParam:"yes"`
}

type DataCenterServers []string
