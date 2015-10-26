package group

type GetHAPolicy struct {
	Group `URIParam:"GroupId" argument:"composed" json:"-"`
}
