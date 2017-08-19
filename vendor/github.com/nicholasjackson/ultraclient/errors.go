package ultraclient

import (
	"fmt"
	"net/url"
)

const (
	ErrorTimeout                 = "timeout"
	ErrorCircuitOpen             = "circuit open"
	ErrorGeneral                 = "general error"
	ErrorUnableToCompleteRequest = "unable to complete request"
)

// ClientError implements the Error interface and is a generic client error
type ClientError struct {
	Message string
	URL     url.URL
}

// Error implements the error interface
func (s ClientError) Error() string {
	return fmt.Sprintf("%v for url: %v", s.Message, s.URL.String())
}
