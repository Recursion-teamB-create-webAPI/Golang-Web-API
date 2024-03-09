package utilError

import "fmt"

type NoKeywordError struct {}

func (e *NoKeywordError) Error() string {
	return fmt.Sprintln("No key word when searching.")
}

func NewNoKeywordError() *NoKeywordError {
	return &NoKeywordError{}
}


type HTTPMethodNotAllowedError struct {
	HTTPMethod string
}

func (e *HTTPMethodNotAllowedError) Error() string {
	return fmt.Sprintf("HTTPMethod name: %v\nThis http method is not allowed in this endpoint.", e.HTTPMethod)
}

func NewHTTPMethodNotAllowedError(httpMethod string) *HTTPMethodNotAllowedError {
	return &HTTPMethodNotAllowedError{
		HTTPMethod: httpMethod,
	}
}