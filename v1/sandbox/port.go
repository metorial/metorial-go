package sandbox

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
)

const tokenHeader = "Metorial-Sandbox-Token"

// Port sends HTTP requests to a port inside the machine. Paths are resolved against
// the port's URL and every request carries a fresh access token.
type Port struct {
	Port int

	access *accessCache
	client *http.Client
	closed atomic.Bool
	url    atomic.Value
}

func connectPort(ctx context.Context, port int, access *accessCache, client *http.Client) (*Port, error) {
	if client == nil {
		client = http.DefaultClient
	}

	grant, err := access.get(ctx)
	if err != nil {
		return nil, err
	}

	connection := &Port{Port: port, access: access, client: client}
	connection.url.Store(grant.URL)
	return connection, nil
}

// URL is the port's base URL.
func (p *Port) URL() string {
	value, _ := p.url.Load().(string)
	return value
}

func (p *Port) Token(ctx context.Context) (string, error) {
	grant, err := p.access.get(ctx)
	if err != nil {
		return "", err
	}
	return grant.Token, nil
}

// Do sends req to the port. req.URL may be a path such as "/api/items?limit=1".
// A 401 response refreshes the token and retries once when the body can be replayed.
func (p *Port) Do(req *http.Request) (*http.Response, error) {
	if p.closed.Load() {
		return nil, closedError()
	}

	send := func(attempt *http.Request) (*http.Response, error) {
		grant, err := p.access.get(attempt.Context())
		if err != nil {
			return nil, err
		}
		p.url.Store(grant.URL)

		target, err := p.resolve(grant.URL, attempt.URL)
		if err != nil {
			return nil, err
		}

		outgoing := attempt.Clone(attempt.Context())
		outgoing.URL = target
		outgoing.Host = target.Host
		outgoing.RequestURI = ""
		outgoing.Header.Set(tokenHeader, grant.Token)

		return p.client.Do(outgoing)
	}

	resp, err := send(req)
	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		return resp, err
	}

	retry := req
	if req.Body != nil && req.Body != http.NoBody {
		if req.GetBody == nil {
			return resp, nil
		}
		body, err := req.GetBody()
		if err != nil {
			return resp, nil
		}
		retry = req.Clone(req.Context())
		retry.Body = body
	}

	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	p.access.invalidate()
	return send(retry)
}

func (p *Port) Request(ctx context.Context, method, path string, body io.Reader, header http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}
	for key, values := range header {
		req.Header[key] = values
	}
	return p.Do(req)
}

func (p *Port) Get(ctx context.Context, path string, header http.Header) (*http.Response, error) {
	return p.Request(ctx, http.MethodGet, path, nil, header)
}

func (p *Port) Head(ctx context.Context, path string, header http.Header) (*http.Response, error) {
	return p.Request(ctx, http.MethodHead, path, nil, header)
}

func (p *Port) Delete(ctx context.Context, path string, header http.Header) (*http.Response, error) {
	return p.Request(ctx, http.MethodDelete, path, nil, header)
}

func (p *Port) Post(ctx context.Context, path, contentType string, body io.Reader) (*http.Response, error) {
	return p.Request(ctx, http.MethodPost, path, body, contentTypeHeader(contentType))
}

func (p *Port) Put(ctx context.Context, path, contentType string, body io.Reader) (*http.Response, error) {
	return p.Request(ctx, http.MethodPut, path, body, contentTypeHeader(contentType))
}

func (p *Port) Patch(ctx context.Context, path, contentType string, body io.Reader) (*http.Response, error) {
	return p.Request(ctx, http.MethodPatch, path, body, contentTypeHeader(contentType))
}

// HTTPClient returns a client that sends every request to this port, so it can be
// handed to libraries that expect an *http.Client. Request URLs may be paths.
func (p *Port) HTTPClient() *http.Client {
	return &http.Client{Transport: portTransport{port: p}}
}

// Close stops further requests through this connection.
func (p *Port) Close() error {
	p.closed.Store(true)
	return nil
}

func (p *Port) resolve(base string, target *url.URL) (*url.URL, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	if target.IsAbs() {
		if target.Scheme != baseURL.Scheme || target.Host != baseURL.Host {
			return nil, fmt.Errorf("sandbox: %s is not on port %d", target, p.Port)
		}
		return target, nil
	}

	resolved := *baseURL
	resolved.Path = strings.TrimRight(baseURL.Path, "/") + "/" + strings.TrimLeft(target.Path, "/")
	resolved.RawPath = ""
	resolved.RawQuery = target.RawQuery
	resolved.Fragment = ""
	return &resolved, nil
}

func contentTypeHeader(contentType string) http.Header {
	if contentType == "" {
		return nil
	}
	return http.Header{"Content-Type": {contentType}}
}

type portTransport struct {
	port *Port
}

func (t portTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.port.Do(req)
}
