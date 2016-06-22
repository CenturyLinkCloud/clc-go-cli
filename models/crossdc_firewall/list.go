package crossdc_firewall

type ListReq struct {
	DataCenter              string `valid:"required" URIParam:"yes"`
	DestinationAccountAlias string `URIParam:"yes"`
}
