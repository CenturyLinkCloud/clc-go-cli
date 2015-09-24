package network

type UpdateReq struct {
	Network     `argument:"composed" URIParam:"NetworkId,DataCenter" json:"-"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
}
