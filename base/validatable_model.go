package base

type ValidatableModel interface {
	Validate() error
}
