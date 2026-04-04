//go:build js && wasm

package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
)

type Client struct {
	apiHost    string
	apiKey     string
	headers    map[string]string
	apiVersion string
}

func NewClient(apiHost, apiKey, apiVersion string, headers map[string]string) *Client {
	return &Client{
		apiHost:    strings.TrimRight(apiHost, "/"),
		apiKey:     apiKey,
		headers:    headers,
		apiVersion: apiVersion,
	}
}

func (c *Client) Get(req *Request, result any) error {
	return c.do("GET", req, result)
}

func (c *Client) Post(req *Request, result any) error {
	return c.do("POST", req, result)
}

func (c *Client) Patch(req *Request, result any) error {
	return c.do("PATCH", req, result)
}

func (c *Client) Put(req *Request, result any) error {
	return c.do("PUT", req, result)
}

func (c *Client) Delete(req *Request, result any) error {
	return c.do("DELETE", req, result)
}

type RawResponse struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
}

func (c *Client) DoRaw(method string, fullURL string, body []byte, headers map[string]string) (*RawResponse, error) {
	return c.doRaw(method, fullURL, body, headers)
}

func (c *Client) do(method string, req *Request, result any) error {
	path := strings.Join(req.Path, "/")
	fullURL := fmt.Sprintf("%s/%s", c.apiHost, path)

	if len(req.Query) > 0 {
		q := url.Values{}
		for k, v := range req.Query {
			q.Set(k, queryValue(v))
		}
		fullURL += "?" + q.Encode()
	}

	var payload []byte
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return fmt.Errorf("metorial: failed to marshal request body: %w", err)
		}
		payload = jsonBody
	}

	response, err := c.doRaw(method, fullURL, payload, req.Headers)
	if err != nil {
		return err
	}

	if response.StatusCode == 204 || len(response.Body) == 0 || result == nil {
		return nil
	}

	if err := json.Unmarshal(response.Body, result); err != nil {
		return fmt.Errorf("metorial: failed to unmarshal response: %w", err)
	}

	return nil
}

func (c *Client) doRaw(method string, fullURL string, body []byte, headers map[string]string) (*RawResponse, error) {
	headerMap := map[string]string{
		"Accept": "application/json",
	}
	if len(body) > 0 {
		headerMap["Content-Type"] = "application/json"
	}
	if strings.TrimSpace(c.apiKey) != "" {
		headerMap["Authorization"] = "Bearer " + c.apiKey
	}
	for key, value := range c.headers {
		headerMap[key] = value
	}
	for key, value := range headers {
		headerMap[key] = value
	}

	options := js.Global().Get("Object").New()
	options.Set("method", method)
	options.Set("credentials", "include")

	jsHeaders := js.Global().Get("Headers").New()
	for key, value := range headerMap {
		jsHeaders.Call("set", key, value)
	}
	options.Set("headers", jsHeaders)

	if len(body) > 0 {
		uint8Array := js.Global().Get("Uint8Array").New(len(body))
		js.CopyBytesToJS(uint8Array, body)
		options.Set("body", uint8Array)
	}

	responseValue, err := await(js.Global().Call("fetch", fullURL, options))
	if err != nil {
		return nil, fmt.Errorf("metorial: request failed: %w", err)
	}

	arrayBufferValue, err := await(responseValue.Call("arrayBuffer"))
	if err != nil {
		return nil, fmt.Errorf("metorial: failed to read response body: %w", err)
	}

	uint8Array := js.Global().Get("Uint8Array").New(arrayBufferValue)
	respBody := make([]byte, uint8Array.Get("length").Int())
	js.CopyBytesToGo(respBody, uint8Array)

	statusCode := responseValue.Get("status").Int()
	statusText := responseValue.Get("statusText").String()
	result := &RawResponse{
		StatusCode: statusCode,
		Status:     strconv.Itoa(statusCode) + " " + statusText,
		Headers:    readHeaders(responseValue.Get("headers")),
		Body:       respBody,
	}

	if statusCode < 200 || statusCode >= 300 {
		apiErr := &Error{
			StatusCode: statusCode,
			RequestID:  result.Headers.Get("X-Request-ID"),
		}
		_ = json.Unmarshal(respBody, apiErr)
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(statusCode)
		}
		return result, apiErr
	}

	return result, nil
}

func await(promise js.Value) (js.Value, error) {
	resultCh := make(chan js.Value, 1)
	errorCh := make(chan error, 1)

	resolve := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := js.Undefined()
		if len(args) > 0 {
			value = args[0]
		}
		resultCh <- value
		return nil
	})
	defer resolve.Release()

	reject := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) > 0 {
			errorCh <- fmt.Errorf("%s", args[0].String())
		} else {
			errorCh <- fmt.Errorf("promise rejected")
		}
		return nil
	})
	defer reject.Release()

	promise.Call("then", resolve).Call("catch", reject)

	select {
	case value := <-resultCh:
		return value, nil
	case err := <-errorCh:
		return js.Undefined(), err
	}
}

func readHeaders(value js.Value) http.Header {
	headers := http.Header{}
	callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 2 {
			return nil
		}

		headerValue := args[0].String()
		headerName := args[1].String()
		headers.Add(headerName, headerValue)
		return nil
	})
	defer callback.Release()

	value.Call("forEach", callback)
	return headers
}

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
