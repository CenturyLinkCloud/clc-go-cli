package autoscale

type GetReq struct {
	Policy `argument:"composed" URIParam:"PolicyId" json:"-"`
}
