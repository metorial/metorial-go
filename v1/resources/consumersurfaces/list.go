package consumersurfaces

import (
	"encoding/json"
	"time"
)

// ConsumerSurfacesListOutputItemsSkillConfiguration - Skill rules applied to this surface.
type ConsumerSurfacesListOutputItemsSkillConfiguration struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when the surface has no skill configuration of its own.
	IsDefault bool `json:"is_default"`
	// AllowScripts - Skills on this surface can run scripts.
	AllowScripts bool `json:"allow_scripts"`
	// AllowedFileExtensions - File extensions skills on this surface can use.
	AllowedFileExtensions []string `json:"allowed_file_extensions"`
	// AllowNonStandardDirectories - Skills on this surface can use directories outside the standard ones.
	AllowNonStandardDirectories bool `json:"allow_non_standard_directories"`
}

// ConsumerSurfacesListOutputItemsAuth - How consumers sign in to this surface.
type ConsumerSurfacesListOutputItemsAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - Consumer session lifetime in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// EmailWhitelist - Email addresses and domains allowed to sign in on this surface.
	EmailWhitelist []string `json:"email_whitelist"`
}

// ConsumerSurfacesListOutputItems represents the consumer surfaces list output items type.
type ConsumerSurfacesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the surface is archived.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AllowConsumerSkillAuthoring - Consumers can author skills on this surface.
	AllowConsumerSkillAuthoring bool `json:"allow_consumer_skill_authoring"`
	// AllowConsumerSkillPublishing - Consumers can publish skills from this surface.
	AllowConsumerSkillPublishing bool `json:"allow_consumer_skill_publishing"`
	// SkillConfiguration - Skill rules applied to this surface.
	SkillConfiguration ConsumerSurfacesListOutputItemsSkillConfiguration `json:"skill_configuration"`
	// Auth - How consumers sign in to this surface.
	Auth ConsumerSurfacesListOutputItemsAuth `json:"auth"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConsumerSurfacesListOutputPagination represents the consumer surfaces list output pagination type.
type ConsumerSurfacesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConsumerSurfacesListOutput represents the consumer surfaces list output type.
type ConsumerSurfacesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []ConsumerSurfacesListOutputItems    `json:"items"`
	Pagination ConsumerSurfacesListOutputPagination `json:"pagination"`
}

// MapConsumerSurfacesListOutputFromJSON deserializes JSON data into a ConsumerSurfacesListOutput.
func MapConsumerSurfacesListOutputFromJSON(data []byte) (*ConsumerSurfacesListOutput, error) {
	var v ConsumerSurfacesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesListOutputToJSON serializes a ConsumerSurfacesListOutput to JSON.
func MapConsumerSurfacesListOutputToJSON(v *ConsumerSurfacesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumerSurfacesListQuery represents the consumer surfaces list query type.
type ConsumerSurfacesListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
}

// MapConsumerSurfacesListQueryFromJSON deserializes JSON data into a ConsumerSurfacesListQuery.
func MapConsumerSurfacesListQueryFromJSON(data []byte) (*ConsumerSurfacesListQuery, error) {
	var v ConsumerSurfacesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesListQueryToJSON serializes a ConsumerSurfacesListQuery to JSON.
func MapConsumerSurfacesListQueryToJSON(v *ConsumerSurfacesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
