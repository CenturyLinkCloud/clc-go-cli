package network

type ReleaseReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
	Network    string `valid:"required" URIParam:"yes"`
}
