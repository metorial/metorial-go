package metorial

import "github.com/metorial/metorial-go/v1/internal/endpoint"

// Error represents an API error response from Metorial.
// Use errors.As to check if an error is an API error:
//
//	var apiErr *metorial.Error
//	if errors.As(err, &apiErr) {
//		fmt.Println(apiErr.StatusCode, apiErr.Message)
//	}
type Error = endpoint.Error
