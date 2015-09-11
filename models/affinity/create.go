package affinity

type CreateReq struct {
	Name       string `valid:"required"`
	DataCenter string `valid:"required" json:"Location"`
}
