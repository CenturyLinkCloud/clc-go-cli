package server

type RemoveIPAddressReq struct {
	Server   `argument:"composed" URIParam:"ServerId" json:"-"`
	PublicIp string `valid:"required" URIParam:"yes" json:"-"`
}
