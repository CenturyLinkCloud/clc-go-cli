package server

type RemoveIPAddressReq struct {
	Server   `argument:"composed" URIParam:"ServerId"`
	PublicIp string `valid:"required" URIParam:"yes"`
}
