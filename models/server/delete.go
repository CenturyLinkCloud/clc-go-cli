package server

type DeleteReq struct {
	ServerId string `valid:"required" URIParam:"yes"`
}
