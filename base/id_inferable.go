package base

// IDInferable represents models that have IDs that can be infered by names.
type IDInferable interface {
	InferID(cn Connection) error
	GetNames(cn Connection, name string) ([]string, error)
}
