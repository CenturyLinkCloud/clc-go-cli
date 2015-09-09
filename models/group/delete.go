package group

type DeleteReq struct {
	Group `argument:"composed" URIParam:"GroupId"`
}
