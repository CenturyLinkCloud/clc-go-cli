package group

type GetScheduledActivities struct {
	Group `URIParam:"GroupId" argument:"composed" json:"-"`
}
