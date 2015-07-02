package base

type Formatter interface {
	FormatOutput(model interface{}) (string, error)
}
