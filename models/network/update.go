package network

type UpdateReq struct {
	DataCenter  string `valid:"required" URIParam:"yes"`
	Network     string `valid:"required" URIParam:"yes"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
}
