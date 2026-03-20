package endpoint

// Request represents an API request.
type Request struct {
	// Path segments that form the URL path (joined with "/").
	Path []string
	// Query contains URL query parameters.
	Query map[string]any
	// Body is the JSON request body. It can be a struct, map, or nil.
	Body any
	// Headers contains additional request-specific headers.
	Headers map[string]string
}
