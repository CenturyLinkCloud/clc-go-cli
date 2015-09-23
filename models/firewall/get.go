package firewall

type GetReq struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	FirewallPolicy string `valid:"required" URIParam:"yes"`
}
