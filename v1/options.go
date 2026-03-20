package metorial

// Option configures the Metorial SDK.
type Option func(*options)

type options struct {
	apiKey  string
	apiHost string
	headers map[string]string
}

const defaultAPIHost = "https://api.metorial.com"

// WithAPIKey sets the API key for authentication.
func WithAPIKey(key string) Option {
	return func(o *options) {
		o.apiKey = key
	}
}

// WithAPIHost overrides the default API host.
func WithAPIHost(host string) Option {
	return func(o *options) {
		o.apiHost = host
	}
}

// WithHeaders sets additional HTTP headers sent with every request.
func WithHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.headers = headers
	}
}
