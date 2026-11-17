package metorial

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAPIHostVectors(t *testing.T) {
	data, err := os.ReadFile("../testdata/api-key-hosts.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixtures []struct {
		Name  string
		Token string
		Host  *string
	}
	if err := json.Unmarshal(data, &fixtures); err != nil {
		t.Fatal(err)
	}

	for _, fixture := range fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			host, expected := "", defaultAPIHost
			if fixture.Host != nil {
				host, expected = *fixture.Host, *fixture.Host
			}
			if got := apiHostFromToken(fixture.Token); got != host {
				t.Fatalf("decoded host = %q, want %q", got, host)
			}
			if fixture.Token == "" {
				return
			}

			sdk, err := New(WithAPIKey(fixture.Token))
			if err != nil {
				t.Fatal(err)
			}
			if got := sdk.client.URL("instance"); got != strings.TrimRight(expected, "/")+"/instance" {
				t.Fatalf("request URL = %q, want host %q", got, expected)
			}

			for _, override := range []string{"https://override.test/custom", defaultAPIHost, ""} {
				for _, opts := range [][]Option{
					{WithAPIKey(fixture.Token), WithAPIHost(override)},
					{WithAPIHost(override), WithAPIKey(fixture.Token)},
				} {
					sdk, err := New(opts...)
					if err != nil {
						t.Fatal(err)
					}
					if got := sdk.client.URL("instance"); got != override+"/instance" {
						t.Fatalf("override request URL = %q, want %q", got, override+"/instance")
					}
				}
			}
		})
	}
}

func tokenForHost(host string) string {
	data, _ := json.Marshal([]string{host, "v2-ext"})
	value := new(big.Int).SetBytes(data)
	base := big.NewInt(62)
	encoded := ""
	for value.Sign() > 0 {
		remainder := new(big.Int)
		value.QuoRem(value, base, remainder)
		encoded = string(apiKeyBase62Alphabet[remainder.Int64()]) + encoded
	}

	return "metorial_sk_" + strings.Repeat("A", 60) + encoded + "v2"
}

func TestTokenHostRequest(t *testing.T) {
	var token string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/custom/v1/instance" {
			t.Errorf("request path = %q", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer "+token {
			t.Error("request did not preserve bearer token")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()
	token = tokenForHost(server.URL + "/custom/v1/")

	sdk, err := New(WithAPIKey(token))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := sdk.Instance.Get(); err != nil {
		t.Fatal(err)
	}
}

func TestOversizedTokenHost(t *testing.T) {
	if got := apiHostFromToken("metorial_sk_" + strings.Repeat("A", 8192) + "v2"); got != "" {
		t.Fatalf("oversized token host = %q", got)
	}
}
