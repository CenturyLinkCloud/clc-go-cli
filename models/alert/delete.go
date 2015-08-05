package alert

type DeleteReq struct {
	PolicyId string `valid:"required" URIParam:"yes"`
}
