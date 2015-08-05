package affinity

type UpdateReq struct {
	PolicyId string `valid:"required" URIParam:"yes"`
	Name     string `valid:"required"`
}
