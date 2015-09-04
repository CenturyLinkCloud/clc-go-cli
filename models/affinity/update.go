package affinity

type UpdateReq struct {
	Policy
	Name string `valid:"required"`
}
