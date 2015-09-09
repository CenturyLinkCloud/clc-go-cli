package group

type GetReq struct {
	Group `URIParam:"GroupId" argument:"composed"`
}
