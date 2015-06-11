package authentication

type LoginReq struct {
	Username string
	Password string
}

type LoginRes struct {
	UserName      string
	AccountAlias  string
	LocationAlias string
	Roles         []string
	BearerToken   string
}
