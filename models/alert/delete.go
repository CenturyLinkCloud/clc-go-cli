package alert

type DeleteReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
}
