package errors

import "fmt"

func EmptyField(field string) error {
	return fmt.Errorf("The %s field must be set and non-empty", field)
}
