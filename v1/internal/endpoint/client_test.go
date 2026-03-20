package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/providers/prov_123" {
			t.Errorf("expected /providers/prov_123, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("expected Bearer test-key, got %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept: application/json, got %s", r.Header.Get("Accept"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"id":   "prov_123",
			"name": "Test Provider",
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := client.Get(&Request{
		Path: []string{"providers", "prov_123"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "prov_123" {
		t.Errorf("expected id prov_123, got %s", result.ID)
	}
	if result.Name != "Test Provider" {
		t.Errorf("expected name Test Provider, got %s", result.Name)
	}
}

func TestClientPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}

		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "New Session" {
			t.Errorf("expected body name 'New Session', got %v", body["name"])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"id":   "sess_456",
			"name": "New Session",
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := client.Post(&Request{
		Path: []string{"sessions"},
		Body: map[string]any{"name": "New Session"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "sess_456" {
		t.Errorf("expected id sess_456, got %s", result.ID)
	}
}

func TestClientPatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		if r.URL.Path != "/sessions/sess_1" {
			t.Errorf("expected /sessions/sess_1, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"id": "sess_1", "name": "Updated"})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := client.Patch(&Request{
		Path: []string{"sessions", "sess_1"},
		Body: map[string]any{"name": "Updated"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "Updated" {
		t.Errorf("expected name Updated, got %s", result.Name)
	}
}

func TestClientDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	err := client.Delete(&Request{
		Path: []string{"sessions", "sess_1"},
	}, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("limit") != "10" {
			t.Errorf("expected limit=10, got %s", r.URL.Query().Get("limit"))
		}
		if r.URL.Query().Get("order") != "asc" {
			t.Errorf("expected order=asc, got %s", r.URL.Query().Get("order"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"items": []any{}})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		Items []any `json:"items"`
	}
	err := client.Get(&Request{
		Path:  []string{"providers"},
		Query: map[string]any{"limit": float64(10), "order": "asc"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientCustomHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "value" {
			t.Errorf("expected X-Custom: value, got %s", r.Header.Get("X-Custom"))
		}
		if r.Header.Get("X-Request") != "per-request" {
			t.Errorf("expected X-Request: per-request, got %s", r.Header.Get("X-Request"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", map[string]string{
		"X-Custom": "value",
	})

	err := client.Get(&Request{
		Path:    []string{"test"},
		Headers: map[string]string{"X-Request": "per-request"},
	}, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-ID", "req_abc")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Provider not found",
			"code":    "not_found",
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	err := client.Get(&Request{
		Path: []string{"providers", "invalid"},
	}, nil)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("expected *Error, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Provider not found" {
		t.Errorf("expected message 'Provider not found', got %s", apiErr.Message)
	}
	if apiErr.Code != "not_found" {
		t.Errorf("expected code 'not_found', got %s", apiErr.Code)
	}
	if apiErr.RequestID != "req_abc" {
		t.Errorf("expected request ID 'req_abc', got %s", apiErr.RequestID)
	}
	if !apiErr.IsNotFound() {
		t.Error("expected IsNotFound() to be true")
	}
	if apiErr.IsRateLimited() {
		t.Error("expected IsRateLimited() to be false")
	}
}

func TestClientRateLimitRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]any{"message": "rate limited"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"id": "prov_1"})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID string `json:"id"`
	}
	err := client.Get(&Request{
		Path: []string{"providers"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("expected 2 attempts, got %d", attempts)
	}
	if result.ID != "prov_1" {
		t.Errorf("expected id prov_1, got %s", result.ID)
	}
}

func TestClientUnauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"message": "Invalid API key"})
	}))
	defer server.Close()

	client := NewClient(server.URL, "bad-key", "2026-01-01-magnetar", nil)

	err := client.Get(&Request{Path: []string{"test"}}, nil)

	apiErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("expected *Error, got %T", err)
	}
	if !apiErr.IsUnauthorized() {
		t.Error("expected IsUnauthorized() to be true")
	}
}

func TestClientPostWithStructBody(t *testing.T) {
	type CreateBody struct {
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "My Deployment" {
			t.Errorf("expected name 'My Deployment', got %v", body["name"])
		}
		if _, ok := body["description"]; ok {
			t.Error("expected description to be omitted")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"id": "deploy_1", "name": "My Deployment"})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := client.Post(&Request{
		Path: []string{"provider-deployments"},
		Body: &CreateBody{Name: "My Deployment"},
	}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "deploy_1" {
		t.Errorf("expected id deploy_1, got %s", result.ID)
	}
}

func TestClientNoContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID string `json:"id"`
	}
	err := client.Delete(&Request{Path: []string{"sessions", "sess_1"}}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientTimeDeserialization(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":"p1","created_at":"2025-06-15T10:30:00Z","updated_at":"2025-06-15T10:30:00Z"}`)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	var result struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	err := client.Get(&Request{Path: []string{"providers", "p1"}}, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.CreatedAt.IsZero() {
		t.Error("expected created_at to be parsed")
	}
	expected := time.Date(2025, 6, 15, 10, 30, 0, 0, time.UTC)
	if !result.CreatedAt.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, result.CreatedAt)
	}
}

func TestStructToQuery(t *testing.T) {
	type Params struct {
		Limit *float64 `json:"limit,omitempty"`
		After *string  `json:"after,omitempty"`
		Order *string  `json:"order,omitempty"`
	}

	limit := float64(25)
	order := "desc"

	q := StructToQuery(&Params{
		Limit: &limit,
		Order: &order,
	})

	if q == nil {
		t.Fatal("expected non-nil query map")
	}
	if q["limit"] != float64(25) {
		t.Errorf("expected limit=25, got %v", q["limit"])
	}
	if q["order"] != "desc" {
		t.Errorf("expected order=desc, got %v", q["order"])
	}
	if _, ok := q["after"]; ok {
		t.Error("expected after to be omitted")
	}
}

func TestStructToQueryNil(t *testing.T) {
	q := StructToQuery(nil)
	if q != nil {
		t.Errorf("expected nil for nil input, got %v", q)
	}
}

func TestStructToBody(t *testing.T) {
	type Body struct {
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	b := StructToBody(&Body{Name: "test"})

	if b == nil {
		t.Fatal("expected non-nil body map")
	}
	if b["name"] != "test" {
		t.Errorf("expected name=test, got %v", b["name"])
	}
	if _, ok := b["description"]; ok {
		t.Error("expected description to be omitted")
	}
}

func TestQueryValue(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{"hello", "hello"},
		{float64(10), "10"},
		{float64(3.14), "3.14"},
		{42, "42"},
		{int64(99), "99"},
		{true, "true"},
		{false, "false"},
	}

	for _, tt := range tests {
		result := queryValue(tt.input)
		if result != tt.expected {
			t.Errorf("queryValue(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestErrorString(t *testing.T) {
	err := &Error{
		StatusCode: 404,
		Message:    "Not found",
	}
	expected := "metorial: Not found (status 404)"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	errWithReqID := &Error{
		StatusCode: 500,
		Message:    "Internal error",
		RequestID:  "req_xyz",
	}
	expected = "metorial: Internal error (status 500, request_id: req_xyz)"
	if errWithReqID.Error() != expected {
		t.Errorf("expected %q, got %q", expected, errWithReqID.Error())
	}
}

func TestClientEmptyResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	err := client.Get(&Request{Path: []string{"test"}}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-key", "2026-01-01-magnetar", nil)

	err := client.Get(&Request{Path: []string{"test"}}, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("expected *Error, got %T", err)
	}
	if apiErr.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Internal Server Error" {
		t.Errorf("expected default status text, got %s", apiErr.Message)
	}
}

func TestClientTrailingSlashTrimmed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/providers" {
			t.Errorf("expected /providers, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{})
	}))
	defer server.Close()

	client := NewClient(server.URL+"/", "test-key", "2026-01-01-magnetar", nil)

	err := client.Get(&Request{Path: []string{"providers"}}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
