package endpoint

import "fmt"

// Error represents an API error response from Metorial.
type Error struct {
	// StatusCode is the HTTP status code.
	StatusCode int `json:"-"`
	// Message is the error message from the API.
	Message string `json:"message"`
	// Code is the error code from the API.
	Code string `json:"code,omitempty"`
	// RequestID is the unique request identifier for debugging.
	RequestID string `json:"-"`
}

func (e *Error) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("metorial: %s (status %d, request_id: %s)", e.Message, e.StatusCode, e.RequestID)
	}
	return fmt.Sprintf("metorial: %s (status %d)", e.Message, e.StatusCode)
}

// IsNotFound reports whether the error is a 404 Not Found.
func (e *Error) IsNotFound() bool {
	return e.StatusCode == 404
}

// IsRateLimited reports whether the error is a 429 Too Many Requests.
func (e *Error) IsRateLimited() bool {
	return e.StatusCode == 429
}

// IsUnauthorized reports whether the error is a 401 Unauthorized.
func (e *Error) IsUnauthorized() bool {
	return e.StatusCode == 401
}
