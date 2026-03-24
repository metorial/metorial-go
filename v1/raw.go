package metorial

import (
	"fmt"
	"net/http"
	"strings"
)

// RawRequest represents a low-level HTTP request sent through the SDK.
type RawRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
}

// RawResponse contains the full response for a raw SDK request.
type RawResponse struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
}

// Fetch performs a raw HTTP request using the SDK's configured auth, host, and retry behavior.
func (sdk *MetorialSdk) Fetch(request *RawRequest) (*RawResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("metorial: request is required")
	}

	method := strings.ToUpper(strings.TrimSpace(request.Method))
	if method == "" {
		method = http.MethodGet
	}

	response, err := sdk.client.DoRaw(method, request.URL, request.Body, request.Headers)
	if response == nil {
		return nil, err
	}

	return &RawResponse{
		StatusCode: response.StatusCode,
		Status:     response.Status,
		Headers:    response.Headers,
		Body:       response.Body,
	}, err
}
