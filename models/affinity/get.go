package affinity

type GetReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
}
