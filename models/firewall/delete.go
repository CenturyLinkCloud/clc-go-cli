package firewall

type DeleteReq struct {
	SourceAccountAlias string `valid:"required" URIParam:"yes"`
	DataCenter         string `valid:"required" URIParam:"yes"`
	FirewallPolicy     string `valid:"required" URIParam:"yes"`
}
