package searchError

import "fmt"

type NoGoogleCustomSearchApiResponseError struct {}

func (e *NoGoogleCustomSearchApiResponseError) Error() string {
	return fmt.Sprintln("Google Custom Search Api didn't get back response.")
} 

func NewNoGoogleCustomSearchApiResponseError() *NoGoogleCustomSearchApiResponseError {
	return &NoGoogleCustomSearchApiResponseError{}
}