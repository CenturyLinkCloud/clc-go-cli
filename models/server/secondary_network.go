package server

type AddNetwork struct {
	ServerId  string `valid:"required" URIParam:"yes"`
	NetworkId string `valid:"required"`
	IPAddress string
}

type RemoveNetwork struct {
	ServerId  string `valid:"required" URIParam:"yes"`
	NetworkId string `valid:"required" URIParam:"yes"`
}
