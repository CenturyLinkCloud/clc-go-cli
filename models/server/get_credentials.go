package server

type GetCredentialsReq struct {
	ServerId string `valid:"required" URIParam:"true"`
}

type GetCredentialsRes struct {
	UserName string
	Password string
}
