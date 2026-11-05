package nullable

import (
	"encoding/json"
	"testing"
)

type body struct {
	Path    string            `json:"path"`
	IfMatch *Nullable[string] `json:"if_match,omitempty"`
}

func TestNullableEncoding(t *testing.T) {
	cases := []struct {
		name string
		in   body
		want string
	}{
		{"omitted", body{Path: "a"}, `{"path":"a"}`},
		{"null", body{Path: "a", IfMatch: Null[string]()}, `{"path":"a","if_match":null}`},
		{"value", body{Path: "a", IfMatch: Of("etag")}, `{"path":"a","if_match":"etag"}`},
	}
	for _, c := range cases {
		got, err := json.Marshal(c.in)
		if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if string(got) != c.want {
			t.Errorf("%s: got %s, want %s", c.name, got, c.want)
		}
	}
}

func TestNullableDecoding(t *testing.T) {
	var n Nullable[string]
	if err := json.Unmarshal([]byte(`null`), &n); err != nil {
		t.Fatal(err)
	}
	if !n.IsNull() {
		t.Fatal("expected null")
	}
	if err := json.Unmarshal([]byte(`"x"`), &n); err != nil {
		t.Fatal(err)
	}
	if v, ok := n.Value(); !ok || v != "x" {
		t.Fatalf("got %q %v", v, ok)
	}
}
