package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client handles HTTP communication with the Metorial API.
type Client struct {
	httpClient *http.Client
	apiHost    string
	apiKey     string
	headers    map[string]string
	apiVersion string
}

// NewClient creates a new API client.
func NewClient(apiHost, apiKey, apiVersion string, headers map[string]string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		apiHost:    strings.TrimRight(apiHost, "/"),
		apiKey:     apiKey,
		headers:    headers,
		apiVersion: apiVersion,
	}
}

// Get performs a GET request.
func (c *Client) Get(req *Request, result any) error {
	return c.do("GET", req, result)
}

// Post performs a POST request.
func (c *Client) Post(req *Request, result any) error {
	return c.do("POST", req, result)
}

// Patch performs a PATCH request.
func (c *Client) Patch(req *Request, result any) error {
	return c.do("PATCH", req, result)
}

// Put performs a PUT request.
func (c *Client) Put(req *Request, result any) error {
	return c.do("PUT", req, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(req *Request, result any) error {
	return c.do("DELETE", req, result)
}

// RawResponse contains the full HTTP response for raw requests.
type RawResponse struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
}

// DoRaw performs a raw HTTP request against the configured API host.
func (c *Client) DoRaw(method string, fullURL string, body []byte, headers map[string]string) (*RawResponse, error) {
	return c.doRawWithRetry(method, fullURL, body, headers, 0)
}

func (c *Client) do(method string, req *Request, result any) error {
	return c.doWithRetry(method, req, result, 0)
}

func (c *Client) doWithRetry(method string, req *Request, result any, tryCount int) error {
	// Build URL
	path := strings.Join(req.Path, "/")
	fullURL := fmt.Sprintf("%s/%s", c.apiHost, path)

	// Add query parameters
	if len(req.Query) > 0 {
		q := url.Values{}
		for k, v := range req.Query {
			q.Set(k, queryValue(v))
		}
		fullURL += "?" + q.Encode()
	}

	// Build request body
	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return fmt.Errorf("metorial: failed to marshal request body: %w", err)
		}
		body = bytes.NewReader(jsonBody)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return fmt.Errorf("metorial: failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	if body != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	for k, v := range c.headers {
		httpReq.Header.Set(k, v)
	}
	if req.Headers != nil {
		for k, v := range req.Headers {
			httpReq.Header.Set(k, v)
		}
	}

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("metorial: request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("metorial: failed to read response body: %w", err)
	}

	// Handle rate limiting with retry
	if resp.StatusCode == 429 && tryCount < 3 {
		retryAfter := resp.Header.Get("Retry-After")
		delay := 3 * time.Second
		if retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil {
				delay = time.Duration(seconds+3) * time.Second
			}
		}
		time.Sleep(delay)
		return c.doWithRetry(method, req, result, tryCount+1)
	}

	// Handle errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &Error{
			StatusCode: resp.StatusCode,
			RequestID:  resp.Header.Get("X-Request-ID"),
		}
		_ = json.Unmarshal(respBody, apiErr)
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(resp.StatusCode)
		}
		return apiErr
	}

	// Handle 204 No Content
	if resp.StatusCode == 204 || len(respBody) == 0 {
		return nil
	}

	// Unmarshal response
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("metorial: failed to unmarshal response: %w", err)
		}
	}

	return nil
}

func (c *Client) doRawWithRetry(method string, fullURL string, body []byte, headers map[string]string, tryCount int) (*RawResponse, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	httpReq, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("metorial: failed to create request: %w", err)
	}

	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	for k, v := range c.headers {
		httpReq.Header.Set(k, v)
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("metorial: request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("metorial: failed to read response body: %w", err)
	}

	if resp.StatusCode == 429 && tryCount < 3 {
		retryAfter := resp.Header.Get("Retry-After")
		delay := 3 * time.Second
		if retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil {
				delay = time.Duration(seconds+3) * time.Second
			}
		}
		time.Sleep(delay)
		return c.doRawWithRetry(method, fullURL, body, headers, tryCount+1)
	}

	result := &RawResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header.Clone(),
		Body:       respBody,
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &Error{
			StatusCode: resp.StatusCode,
			RequestID:  resp.Header.Get("X-Request-ID"),
		}
		_ = json.Unmarshal(respBody, apiErr)
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(resp.StatusCode)
		}
		return result, apiErr
	}

	return result, nil
}

// queryValue converts a value to a query string value.
func queryValue(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		if val == float64(int64(val)) {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// StructToQuery converts a struct with json tags to a query parameter map.
// It uses json marshaling to respect json tags and omitempty.
func StructToQuery(v any) map[string]any {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	return m
}

// StructToBody converts a struct to a request body map.
func StructToBody(v any) map[string]any {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	return m
}
