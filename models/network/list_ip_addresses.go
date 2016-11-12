package network

type ListIpAddresses struct {
	Network `argument:"composed" URIParam:"NetworkId,DataCenter" json:"-"`
	Type    string `URIParam:"yes" oneOf:"claimed,free,all,optional"`
}
