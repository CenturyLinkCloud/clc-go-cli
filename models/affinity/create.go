package affinity

type CreateReq struct {
	Name     string `valid:"required"`
	Location string `valid:"required"`
}
