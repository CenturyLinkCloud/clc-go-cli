package affinity

type UpdateReq struct {
	Policy `argument:"composed" URIParam:"PolicyId"`
	Name   string `valid:"required"`
}
