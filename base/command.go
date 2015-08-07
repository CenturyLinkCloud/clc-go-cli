package base

type Command interface {
	Execute(cn Connection) error
	Resource() string
	Command() string
	ShowBrief() string
	ShowHelp() string
	InputModel() interface{}
	OutputModel() interface{}
}
