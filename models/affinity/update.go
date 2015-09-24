package affinity

type UpdateReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
	Name   string `valid:"required"`
}
