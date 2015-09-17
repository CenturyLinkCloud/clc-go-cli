package server

type GetCredentialsReq struct {
	Server `argument:"composed" URIParam:"ServerId"`
}

type GetCredentialsRes struct {
	UserName string
	Password string
}
