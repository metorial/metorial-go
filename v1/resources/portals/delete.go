package portals

import (
	"encoding/json"
	"time"
)

// PortalsDeleteOutputSkillConfiguration represents the portals delete output skill configuration type.
type PortalsDeleteOutputSkillConfiguration struct {
	Object                      string   `json:"object"`
	Id                          string   `json:"id"`
	IsDefault                   bool     `json:"is_default"`
	AllowScripts                bool     `json:"allow_scripts"`
	AllowedFileExtensions       []string `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool     `json:"allow_non_standard_directories"`
}

// PortalsDeleteOutputAuth represents the portals delete output auth type.
type PortalsDeleteOutputAuth struct {
	Object                     string  `json:"object"`
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
}

// PortalsDeleteOutputUrls represents the portals delete output urls type.
type PortalsDeleteOutputUrls struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PortalsDeleteOutput represents the portals delete output type.
type PortalsDeleteOutput struct {
	Object                       string                                `json:"object"`
	Id                           string                                `json:"id"`
	Status                       string                                `json:"status"`
	Name                         string                                `json:"name"`
	Slug                         string                                `json:"slug"`
	Description                  *string                               `json:"description,omitempty"`
	AllowConsumerSkillAuthoring  bool                                  `json:"allow_consumer_skill_authoring"`
	AllowConsumerSkillPublishing bool                                  `json:"allow_consumer_skill_publishing"`
	SkillConfiguration           PortalsDeleteOutputSkillConfiguration `json:"skill_configuration"`
	Auth                         PortalsDeleteOutputAuth               `json:"auth"`
	Urls                         []PortalsDeleteOutputUrls             `json:"urls"`
	CreatedAt                    time.Time                             `json:"created_at"`
	UpdatedAt                    time.Time                             `json:"updated_at"`
}

// MapPortalsDeleteOutputFromJSON deserializes JSON data into a PortalsDeleteOutput.
func MapPortalsDeleteOutputFromJSON(data []byte) (*PortalsDeleteOutput, error) {
	var v PortalsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsDeleteOutputToJSON serializes a PortalsDeleteOutput to JSON.
func MapPortalsDeleteOutputToJSON(v *PortalsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
