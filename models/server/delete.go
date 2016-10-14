package server

type DeleteReq struct {
	Server `argument:"composed" URIParam:"ServerId"  json:"-"`
}
