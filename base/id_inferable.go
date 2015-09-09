package base

// IDInferable represents models that have IDs that can be infered from names
// through communicating the server.
type IDInferable interface {
	// After InferID has been completed, all the IDs possible have to have
	// been infered from names.
	InferID(cn Connection) error
	// GetNames exists for the purpose of getting the list of names of entities.
	// The name argument is a field that points to an entity name.
	GetNames(cn Connection, name string) ([]string, error)
}
