package network

type CreateReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}
