package network

import (
	"fmt"
)

type GetReq struct {
	Network     `argument:"composed" URIParam:"NetworkId,DataCenter" json:"-"`
	IpAddresses string `URIParam:"yes" oneOf:"none,claimed,free,all"`
}

func (g *GetReq) Validate() error {
	err := g.Network.Validate()
	if err != nil {
		return err
	}

	if g.IpAddresses == "" {
		return nil
	}
	gip := g.IpAddresses
	if gip != "none" && gip != "claimed" && gip != "free" && gip != "all" {
		return fmt.Errorf("ip-addresses value must be one of the none, claimed, free, all.")
	}
	return nil
}
