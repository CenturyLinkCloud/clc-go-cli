package model_loader

type (
	ParseObjWrongTypeError struct{}
)

func (p ParseObjWrongTypeError) Error() string {
	return ""
}
