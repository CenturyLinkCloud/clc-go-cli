package server

type RestoreReq struct {
	ServerId string `valid:"required" URIParam:"yes"`

	TargetGroupId string `valid:"required"`
}
