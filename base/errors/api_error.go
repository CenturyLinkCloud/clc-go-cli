package errors

import (
	"fmt"
)

type ApiError struct {
	StatusCode  int
	ApiResponse interface{}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error occured while sending request to API. Status code: %d", e.StatusCode)
}
