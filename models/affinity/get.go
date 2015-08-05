package affinity

type GetReq struct {
	PolicyId string `valid:"required" URIParam:"yes"`
}
