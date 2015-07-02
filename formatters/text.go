package formatters

type TextFormatter struct{}

func (f *TextFormatter) FormatOutput(model interface{}) (res string, err error) {
	return "", err
}
