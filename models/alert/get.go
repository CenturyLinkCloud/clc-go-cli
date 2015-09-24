package alert

type GetReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
}
