package server

type AddNetwork struct {
	Server    `argument:"composed" URIParam:"ServerId"`
	NetworkId string `valid:"required"`
	IPAddress string
}

type RemoveNetwork struct {
	Server    `argument:"composed" URIParam:"ServerId"`
	NetworkId string `valid:"required" URIParam:"yes"`
}
