package commands

type Login struct {
	CommandBase
}

type inputStub struct{}

func NewLogin(info CommandExcInfo) *Login {
	l := Login{}
	l.ExcInfo = info
	return &l
}

func (l *Login) InputModel() interface{} {
	return &inputStub{}
}
