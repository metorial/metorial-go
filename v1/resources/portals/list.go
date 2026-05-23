package portals

import (
	"encoding/json"
	"time"
)

// PortalsListOutputItemsSkillConfiguration represents the portals list output items skill configuration type.
type PortalsListOutputItemsSkillConfiguration struct {
	Object                      string   `json:"object"`
	Id                          string   `json:"id"`
	IsDefault                   bool     `json:"is_default"`
	AllowScripts                bool     `json:"allow_scripts"`
	AllowedFileExtensions       []string `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool     `json:"allow_non_standard_directories"`
}

// PortalsListOutputItemsAuthAllowedRedirectUrlFilters represents the portals list output items auth allowed redirect url filters type.
type PortalsListOutputItemsAuthAllowedRedirectUrlFilters struct {
	Url string `json:"url"`
}

// PortalsListOutputItemsAuth represents the portals list output items auth type.
type PortalsListOutputItemsAuth struct {
	Object                     string                                                `json:"object"`
	SessionExpiryTimeInSeconds float64                                               `json:"session_expiry_time_in_seconds"`
	AllowedRedirectUrlFilters  []PortalsListOutputItemsAuthAllowedRedirectUrlFilters `json:"allowed_redirect_url_filters"`
}

// PortalsListOutputItemsUrls represents the portals list output items urls type.
type PortalsListOutputItemsUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsListOutputItems represents the portals list output items type.
type PortalsListOutputItems struct {
	Object                       string                                   `json:"object"`
	Id                           string                                   `json:"id"`
	Status                       string                                   `json:"status"`
	Name                         string                                   `json:"name"`
	Slug                         string                                   `json:"slug"`
	Description                  *string                                  `json:"description,omitempty"`
	AllowConsumerSkillAuthoring  bool                                     `json:"allow_consumer_skill_authoring"`
	AllowConsumerSkillPublishing bool                                     `json:"allow_consumer_skill_publishing"`
	SkillConfiguration           PortalsListOutputItemsSkillConfiguration `json:"skill_configuration"`
	Auth                         PortalsListOutputItemsAuth               `json:"auth"`
	Urls                         []PortalsListOutputItemsUrls             `json:"urls"`
	CreatedAt                    time.Time                                `json:"created_at"`
	UpdatedAt                    time.Time                                `json:"updated_at"`
}

// PortalsListOutputPagination represents the portals list output pagination type.
type PortalsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsListOutput represents the portals list output type.
type PortalsListOutput struct {
	Items      []PortalsListOutputItems    `json:"items"`
	Pagination PortalsListOutputPagination `json:"pagination"`
}

// MapPortalsListOutputFromJSON deserializes JSON data into a PortalsListOutput.
func MapPortalsListOutputFromJSON(data []byte) (*PortalsListOutput, error) {
	var v PortalsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListOutputToJSON serializes a PortalsListOutput to JSON.
func MapPortalsListOutputToJSON(v *PortalsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListQuery represents the portals list query type.
type PortalsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapPortalsListQueryFromJSON deserializes JSON data into a PortalsListQuery.
func MapPortalsListQueryFromJSON(data []byte) (*PortalsListQuery, error) {
	var v PortalsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListQueryToJSON serializes a PortalsListQuery to JSON.
func MapPortalsListQueryToJSON(v *PortalsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
