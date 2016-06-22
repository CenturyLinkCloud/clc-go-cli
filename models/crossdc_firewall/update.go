package crossdc_firewall

type UpdateReq struct {
	DataCenter     string `json:"-" valid:"required" URIParam:"yes"`
	FirewallPolicy string `json:"-" valid:"required" URIParam:"yes"`
	Enabled        string `json:"-" valid:"required" URIParam:"yes" oneOf:"true,false"`
}
