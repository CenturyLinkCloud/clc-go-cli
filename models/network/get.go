package network

import (
	"fmt"
)

type GetReq struct {
	DataCenter  string `valid:"required" URIParam:"yes"`
	Network     string `valid:"required" URIParam:"yes"`
	IpAddresses string `URIParam:"yes" oneOf:"none,claimed,free,all"`
}

func (g *GetReq) Validate() error {
	if g.IpAddresses == "" {
		return nil
	}
	gip := g.IpAddresses
	if gip != "none" && gip != "claimed" && gip != "free" && gip != "all" {
		return fmt.Errorf("ip-addresses value must be one of the none, claimed, free, all.")
	}
	return nil
}
