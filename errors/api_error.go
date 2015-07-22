package errors

import (
	"fmt"
)

type ApiError struct {
	StatusCode  int
	ApiResponse interface{}
	Reason      string
}

func (e *ApiError) Error() string {
	msg := fmt.Sprintf("Error occured while sending request to API. Status code: %d.", e.StatusCode)
	if e.Reason != "" {
		msg += fmt.Sprintf(" Server says: %s.", e.Reason)
	}
	return msg
}
