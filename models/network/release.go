package network

type ReleaseReq struct {
	Network `argument:"composed" URIParam:"NetworkId,DataCenter" json:"-"`
}
