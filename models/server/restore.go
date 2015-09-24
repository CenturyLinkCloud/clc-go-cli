package server

type RestoreReq struct {
	ServerId string `json:"-" valid:"required" URIParam:"yes"`

	TargetGroupId string `valid:"required"`
}
