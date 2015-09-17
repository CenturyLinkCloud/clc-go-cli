package server

type AddNetwork struct {
	ServerId  string `json:"-" valid:"required" URIParam:"yes"`
	NetworkId string `valid:"required"`
	IpAddress string
}

type RemoveNetwork struct {
	ServerId  string `valid:"required" URIParam:"yes"`
	NetworkId string `valid:"required" URIParam:"yes"`
}
