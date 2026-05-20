package portals

import (
	"encoding/json"
	"time"
)

// PortalsCreateOutputSkillConfiguration represents the portals create output skill configuration type.
type PortalsCreateOutputSkillConfiguration struct {
	Object                      string   `json:"object"`
	Id                          string   `json:"id"`
	IsDefault                   bool     `json:"is_default"`
	AllowScripts                bool     `json:"allow_scripts"`
	AllowedFileExtensions       []string `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool     `json:"allow_non_standard_directories"`
}

// PortalsCreateOutputAuth represents the portals create output auth type.
type PortalsCreateOutputAuth struct {
	Object                     string  `json:"object"`
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
}

// PortalsCreateOutputUrls represents the portals create output urls type.
type PortalsCreateOutputUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsCreateOutput represents the portals create output type.
type PortalsCreateOutput struct {
	Object                       string                                `json:"object"`
	Id                           string                                `json:"id"`
	Status                       string                                `json:"status"`
	Name                         string                                `json:"name"`
	Slug                         string                                `json:"slug"`
	Description                  *string                               `json:"description,omitempty"`
	AllowConsumerSkillAuthoring  bool                                  `json:"allow_consumer_skill_authoring"`
	AllowConsumerSkillPublishing bool                                  `json:"allow_consumer_skill_publishing"`
	SkillConfiguration           PortalsCreateOutputSkillConfiguration `json:"skill_configuration"`
	Auth                         PortalsCreateOutputAuth               `json:"auth"`
	Urls                         []PortalsCreateOutputUrls             `json:"urls"`
	CreatedAt                    time.Time                             `json:"created_at"`
	UpdatedAt                    time.Time                             `json:"updated_at"`
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
	Name                         string                                        `json:"name"`
	Description                  *string                                       `json:"description,omitempty"`
	AllowedRedirectUrlFilters    *[]PortalsCreateBodyAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters,omitempty"`
	SessionExpiryTimeInSeconds   *float64                                      `json:"session_expiry_time_in_seconds,omitempty"`
	AllowConsumerSkillAuthoring  *bool                                         `json:"allow_consumer_skill_authoring,omitempty"`
	AllowConsumerSkillPublishing *bool                                         `json:"allow_consumer_skill_publishing,omitempty"`
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
