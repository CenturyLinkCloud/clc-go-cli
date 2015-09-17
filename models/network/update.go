package network

type UpdateReq struct {
	Network     `argument:"composed" URIParam:"NetworkId,DataCenter"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
}
