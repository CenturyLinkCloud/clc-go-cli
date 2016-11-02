package firewall

type DeleteReq struct {
	DataCenter     string `valid:"required" URIParam:"yes" json:"-"`
	FirewallPolicy string `valid:"required" URIParam:"yes" json:"-"`
}
