package oneandone

import (
	"fmt"
)

type errorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ApiError struct {
	httpStatusCode int
	message        string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%d - %s", e.httpStatusCode, e.message)
}

func (e *ApiError) HttpStatusCode() int {
	return e.httpStatusCode
}

func (e *ApiError) Message() string {
	return e.message
}
