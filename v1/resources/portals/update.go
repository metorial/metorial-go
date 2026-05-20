package portals

import (
	"encoding/json"
	"time"
)

// PortalsUpdateOutputSkillConfiguration represents the portals update output skill configuration type.
type PortalsUpdateOutputSkillConfiguration struct {
	Object                      string   `json:"object"`
	Id                          string   `json:"id"`
	IsDefault                   bool     `json:"is_default"`
	AllowScripts                bool     `json:"allow_scripts"`
	AllowedFileExtensions       []string `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool     `json:"allow_non_standard_directories"`
}

// PortalsUpdateOutputAuth represents the portals update output auth type.
type PortalsUpdateOutputAuth struct {
	Object                     string  `json:"object"`
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
}

// PortalsUpdateOutputUrls represents the portals update output urls type.
type PortalsUpdateOutputUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsUpdateOutput represents the portals update output type.
type PortalsUpdateOutput struct {
	Object                       string                                `json:"object"`
	Id                           string                                `json:"id"`
	Status                       string                                `json:"status"`
	Name                         string                                `json:"name"`
	Slug                         string                                `json:"slug"`
	Description                  *string                               `json:"description,omitempty"`
	AllowConsumerSkillAuthoring  bool                                  `json:"allow_consumer_skill_authoring"`
	AllowConsumerSkillPublishing bool                                  `json:"allow_consumer_skill_publishing"`
	SkillConfiguration           PortalsUpdateOutputSkillConfiguration `json:"skill_configuration"`
	Auth                         PortalsUpdateOutputAuth               `json:"auth"`
	Urls                         []PortalsUpdateOutputUrls             `json:"urls"`
	CreatedAt                    time.Time                             `json:"created_at"`
	UpdatedAt                    time.Time                             `json:"updated_at"`
}

// MapPortalsUpdateOutputFromJSON deserializes JSON data into a PortalsUpdateOutput.
func MapPortalsUpdateOutputFromJSON(data []byte) (*PortalsUpdateOutput, error) {
	var v PortalsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsUpdateOutputToJSON serializes a PortalsUpdateOutput to JSON.
func MapPortalsUpdateOutputToJSON(v *PortalsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsUpdateBodyAllowedRedirectUrlFilters represents the portals update body allowed redirect url filters type.
type PortalsUpdateBodyAllowedRedirectUrlFilters struct {
	Url string `json:"url"`
}

// PortalsUpdateBodySkillConfiguration represents the portals update body skill configuration type.
type PortalsUpdateBodySkillConfiguration struct {
	AllowScripts                *bool     `json:"allow_scripts,omitempty"`
	AllowedFileExtensions       *[]string `json:"allowed_file_extensions,omitempty"`
	AllowNonStandardDirectories *bool     `json:"allow_non_standard_directories,omitempty"`
}

// PortalsUpdateBody represents the portals update body type.
type PortalsUpdateBody struct {
	Name                         *string                                       `json:"name,omitempty"`
	Description                  *string                                       `json:"description,omitempty"`
	AllowedRedirectUrlFilters    *[]PortalsUpdateBodyAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters,omitempty"`
	SessionExpiryTimeInSeconds   *float64                                      `json:"session_expiry_time_in_seconds,omitempty"`
	AllowConsumerSkillAuthoring  *bool                                         `json:"allow_consumer_skill_authoring,omitempty"`
	AllowConsumerSkillPublishing *bool                                         `json:"allow_consumer_skill_publishing,omitempty"`
	SkillConfiguration           *PortalsUpdateBodySkillConfiguration          `json:"skill_configuration,omitempty"`
}

// MapPortalsUpdateBodyFromJSON deserializes JSON data into a PortalsUpdateBody.
func MapPortalsUpdateBodyFromJSON(data []byte) (*PortalsUpdateBody, error) {
	var v PortalsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsUpdateBodyToJSON serializes a PortalsUpdateBody to JSON.
func MapPortalsUpdateBodyToJSON(v *PortalsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
