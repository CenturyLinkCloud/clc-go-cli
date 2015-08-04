package network

type GetReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}
