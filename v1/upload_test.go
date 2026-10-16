package metorial

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestUploadFileSendsSmallFilesAsMultipart(t *testing.T) {
	local := filepath.Join(t.TempDir(), "notes.txt")
	writeTestFile(t, local, "hello")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/files" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer metorial_sk_test" {
			t.Errorf("expected API credentials, got %q", got)
		}
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			t.Fatal(err)
		}

		expected := map[string]string{
			"purpose":       "generic",
			"title":         "Notes",
			"instance_id":   "inst_1",
			"store_id":      "store_1",
			"path":          "docs/notes.txt",
			"store_replace": "true",
		}
		for key, value := range expected {
			if got := r.FormValue(key); got != value {
				t.Errorf("expected %s=%q, got %q", key, value, got)
			}
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		data, _ := io.ReadAll(file)
		if header.Filename != "notes.txt" || string(data) != "hello" {
			t.Errorf("unexpected file %q with contents %q", header.Filename, data)
		}
		if got := header.Header.Get("Content-Type"); !strings.HasPrefix(got, "text/plain") {
			t.Errorf("unexpected file content type %q", got)
		}

		writeTestJSON(w, http.StatusOK, map[string]any{"object": "file", "id": "file_1"})
	}))
	defer server.Close()

	result, err := newTestSdk(t, server.URL).UploadFile(context.Background(), local, &UploadFileOptions{
		Title:        "Notes",
		InstanceId:   "inst_1",
		StoreId:      "store_1",
		StorePath:    "docs/notes.txt",
		StoreReplace: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != "file_1" {
		t.Fatalf("expected file_1, got %q", result.Id)
	}
}

func TestUploadFileRejectsDirectories(t *testing.T) {
	directory := t.TempDir()
	writeTestFile(t, filepath.Join(directory, "notes.txt"), "hello")

	sdk := newTestSdk(t, "http://127.0.0.1:1")
	if _, err := sdk.UploadFile(context.Background(), directory, nil); err == nil || !strings.Contains(err.Error(), "is not a file") {
		t.Fatalf("expected a directory to be rejected, got %v", err)
	}
}

func TestUploadFileReaderSendsLargeFilesThroughAnUploadURL(t *testing.T) {
	previous := directUploadLimit
	directUploadLimit = 4
	t.Cleanup(func() { directUploadLimit = previous })

	contents := "large contents"
	steps := []string{}
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/storage":
			steps = append(steps, "put")
			if r.Header.Get("Authorization") != "" {
				t.Error("expected no Metorial credentials on the storage request")
			}
			if got := r.Header.Get("Content-Type"); got != "application/octet-stream" {
				t.Errorf("expected the declared content type, got %q", got)
			}
			data, _ := io.ReadAll(r.Body)
			if string(data) != contents {
				t.Errorf("unexpected uploaded contents %q", data)
			}
			w.WriteHeader(http.StatusOK)

		case r.Method == http.MethodPost && r.URL.Path == "/files":
			body := map[string]any{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}

			switch body["mode"] {
			case "get_upload_url":
				steps = append(steps, "get_upload_url")
				if body["file_name"] != "archive.bin" || body["file_size"] != float64(len(contents)) || body["instance_id"] != "inst_1" {
					t.Errorf("unexpected upload URL request %v", body)
				}
				writeTestJSON(w, http.StatusOK, map[string]any{
					"object": "file_upload",
					"id":     "fup_1",
					"upload": map[string]any{"url": server.URL + "/storage", "method": "PUT"},
				})
			case "complete":
				steps = append(steps, "complete")
				if body["file_upload_id"] != "fup_1" || body["instance_id"] != "inst_1" {
					t.Errorf("unexpected complete request %v", body)
				}
				writeTestJSON(w, http.StatusOK, map[string]any{"object": "file", "id": "file_2"})
			default:
				t.Errorf("unexpected mode %v", body["mode"])
			}

		default:
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	result, err := newTestSdk(t, server.URL).UploadFileReader(context.Background(), "archive.bin", strings.NewReader(contents), int64(len(contents)), &UploadFileOptions{InstanceId: "inst_1"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != "file_2" {
		t.Fatalf("expected file_2, got %q", result.Id)
	}
	if expected := []string{"get_upload_url", "put", "complete"}; !slices.Equal(steps, expected) {
		t.Fatalf("expected steps %v, got %v", expected, steps)
	}
}

func TestUploadFileReportsStorageFailures(t *testing.T) {
	previous := directUploadLimit
	directUploadLimit = 4
	t.Cleanup(func() { directUploadLimit = previous })

	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			http.Error(w, "SignatureDoesNotMatch", http.StatusForbidden)
			return
		}
		writeTestJSON(w, http.StatusOK, map[string]any{
			"id":     "fup_1",
			"upload": map[string]any{"url": server.URL + "/storage", "method": "PUT"},
		})
	}))
	defer server.Close()

	_, err := newTestSdk(t, server.URL).UploadFileReader(context.Background(), "archive.bin", strings.NewReader("large contents"), 14, nil)
	if err == nil || !strings.Contains(err.Error(), "SignatureDoesNotMatch") {
		t.Fatalf("expected the storage error to be reported, got %v", err)
	}
}

func TestUploadFileReturnsAPIErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeTestJSON(w, http.StatusBadRequest, map[string]any{"status": "error", "message": "Unknown file purpose"})
	}))
	defer server.Close()

	_, err := newTestSdk(t, server.URL).UploadFileReader(context.Background(), "notes.txt", strings.NewReader("hello"), 5, &UploadFileOptions{Purpose: "nope"})
	if err == nil || !strings.Contains(err.Error(), "Unknown file purpose") {
		t.Fatalf("expected the API error, got %v", err)
	}
}

func TestUploadFileValidatesStoreOptions(t *testing.T) {
	sdk := newTestSdk(t, "http://127.0.0.1:1")
	if _, err := sdk.UploadFileReader(context.Background(), "x.txt", strings.NewReader("x"), 1, &UploadFileOptions{StoreId: "store_1"}); err == nil {
		t.Fatal("expected StoreId without StorePath to fail")
	}
	if _, err := sdk.UploadFileReader(context.Background(), "x.txt", strings.NewReader("x"), 1, &UploadFileOptions{StoreReplace: true}); err == nil {
		t.Fatal("expected StoreReplace without a store to fail")
	}
}

func newTestSdk(t *testing.T, host string) *MetorialSdk {
	t.Helper()

	sdk, err := New(WithAPIKey("metorial_sk_test"), WithAPIHost(host))
	if err != nil {
		t.Fatal(err)
	}
	return sdk
}

func writeTestFile(t *testing.T, path, contents string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeTestJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
