package base

type Command interface {
	Execute() error
	Resource() string
	Command() string
	InputModel() interface{}
	OutputModel() interface{}
}
