package affinity

type DeleteReq struct {
	PolicyId string `valid:"required" URIParam:"yes"`
}
