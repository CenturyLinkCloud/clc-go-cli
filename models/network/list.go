package network

type ListReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}
