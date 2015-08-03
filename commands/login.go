package commands

type Login struct {
	CommandBase
}

func NewLogin(info CommandExcInfo) *Login {
	l := Login{}
	l.ExcInfo = info
	return &l
}
