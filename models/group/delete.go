package group

type DeleteReq struct {
	GroupId string `valid:"required" URIParam:"yes"`
}
