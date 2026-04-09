package portals

import (
	"encoding/json"
	"time"
)

// PortalsCreateOutputAuthAllowedRedirectUrlFilters represents the portals create output auth allowed redirect url filters type.
type PortalsCreateOutputAuthAllowedRedirectUrlFilters struct {
	Url string `json:"url"`
}

// PortalsCreateOutputAuth represents the portals create output auth type.
type PortalsCreateOutputAuth struct {
	Object                     string                                             `json:"object"`
	SessionExpiryTimeInSeconds float64                                            `json:"session_expiry_time_in_seconds"`
	AllowedRedirectUrlFilters  []PortalsCreateOutputAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsCreateOutputUrls represents the portals create output urls type.
type PortalsCreateOutputUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsCreateOutputBrand represents the portals create output brand type.
type PortalsCreateOutputBrand struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}

// PortalsCreateOutput represents the portals create output type.
type PortalsCreateOutput struct {
	Object      string                    `json:"object"`
	Id          string                    `json:"id"`
	Status      string                    `json:"status"`
	Name        string                    `json:"name"`
	Slug        string                    `json:"slug"`
	Description *string                   `json:"description,omitempty"`
	Auth        PortalsCreateOutputAuth   `json:"auth"`
	Urls        []PortalsCreateOutputUrls `json:"urls"`
	Brand       PortalsCreateOutputBrand  `json:"brand"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

// MapPortalsCreateOutputFromJSON deserializes JSON data into a PortalsCreateOutput.
func MapPortalsCreateOutputFromJSON(data []byte) (*PortalsCreateOutput, error) {
	var v PortalsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsCreateOutputToJSON serializes a PortalsCreateOutput to JSON.
func MapPortalsCreateOutputToJSON(v *PortalsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsCreateBodyAllowedRedirectUrlFilters represents the portals create body allowed redirect url filters type.
type PortalsCreateBodyAllowedRedirectUrlFilters struct {
	Url string `json:"url"`
}

// PortalsCreateBody represents the portals create body type.
type PortalsCreateBody struct {
	Name                       string                                        `json:"name"`
	Description                *string                                       `json:"description,omitempty"`
	AllowedRedirectUrlFilters  *[]PortalsCreateBodyAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters,omitempty"`
	SessionExpiryTimeInSeconds *float64                                      `json:"session_expiry_time_in_seconds,omitempty"`
}

// MapPortalsCreateBodyFromJSON deserializes JSON data into a PortalsCreateBody.
func MapPortalsCreateBodyFromJSON(data []byte) (*PortalsCreateBody, error) {
	var v PortalsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsCreateBodyToJSON serializes a PortalsCreateBody to JSON.
func MapPortalsCreateBodyToJSON(v *PortalsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
