package network

type ListIpAddresses struct {
	DataCenter string `valid:"required" URIParam:"yes"`
	Network    string `valid:"required" URIParam:"yes"`
	Type       string `URIParam:"yes" oneOf:"claimed,free,all"`
}
