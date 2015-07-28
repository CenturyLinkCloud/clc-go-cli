package group

type GetReq struct {
	GroupId string `valid:"required" URIParam:"yes"`
}
