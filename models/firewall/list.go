package firewall

type ListReq struct {
	SourceAccountAlias      string `valid:"required" URIParam:"yes"`
	DataCenter              string `valid:"required" URIParam:"yes"`
	DestinationAccountAlias string `URIParam:"yes"`
}
