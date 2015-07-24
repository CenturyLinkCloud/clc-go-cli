package server

type RemoveIPAddressReq struct {
	ServerId string `valid:"required" URIParam:"yes"`
	PublicIp string `valid:"required" URIParam:"yes"`
}
