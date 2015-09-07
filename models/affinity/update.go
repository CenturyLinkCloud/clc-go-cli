package affinity

type UpdateReq struct {
	Policy `argument:"composed"`
	Name   string `valid:"required"`
}
