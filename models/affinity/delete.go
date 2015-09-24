package affinity

type DeleteReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
}
