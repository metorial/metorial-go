package consumersurfaces

import (
	"encoding/json"
	"time"
)

// ConsumerSurfacesGetOutputSkillConfiguration - Skill rules applied to this surface.
type ConsumerSurfacesGetOutputSkillConfiguration struct {
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

// ConsumerSurfacesGetOutputAuth - How consumers sign in to this surface.
type ConsumerSurfacesGetOutputAuth struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SessionExpiryTimeInSeconds - Consumer session lifetime in seconds.
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
	// EmailWhitelist - Email addresses and domains allowed to sign in on this surface.
	EmailWhitelist []string `json:"email_whitelist"`
}

// ConsumerSurfacesGetOutput represents the consumer surfaces get output type.
type ConsumerSurfacesGetOutput struct {
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
	SkillConfiguration ConsumerSurfacesGetOutputSkillConfiguration `json:"skill_configuration"`
	// Auth - How consumers sign in to this surface.
	Auth ConsumerSurfacesGetOutputAuth `json:"auth"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapConsumerSurfacesGetOutputFromJSON deserializes JSON data into a ConsumerSurfacesGetOutput.
func MapConsumerSurfacesGetOutputFromJSON(data []byte) (*ConsumerSurfacesGetOutput, error) {
	var v ConsumerSurfacesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesGetOutputToJSON serializes a ConsumerSurfacesGetOutput to JSON.
func MapConsumerSurfacesGetOutputToJSON(v *ConsumerSurfacesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
