package mcpavailability

import (
	"encoding/json"
)

// SkillsMcpAvailabilityGetOutput represents the skills mcp availability get output type.
type SkillsMcpAvailabilityGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// SkillId - The ID of the skill.
	SkillId string `json:"skill_id"`
	// Available - Whether apps connected over MCP can see the skill. Skills over a limit are left out of MCP connections but keep working everywhere else.
	Available bool `json:"available"`
	// ExceededLimits - The limits the skill is over. Empty when the skill is available.
	ExceededLimits []string `json:"exceeded_limits"`
	// FileCount - Number of files the skill would share over MCP.
	FileCount float64 `json:"file_count"`
	// Size - Total size in bytes of the files the skill would share over MCP.
	Size float64 `json:"size"`
	// MaxFileCount - Most files a skill can have and still be shared over MCP.
	MaxFileCount float64 `json:"max_file_count"`
	// MaxSize - Largest total size in bytes a skill can have and still be shared over MCP.
	MaxSize float64 `json:"max_size"`
}

// MapSkillsMcpAvailabilityGetOutputFromJSON deserializes JSON data into a SkillsMcpAvailabilityGetOutput.
func MapSkillsMcpAvailabilityGetOutputFromJSON(data []byte) (*SkillsMcpAvailabilityGetOutput, error) {
	var v SkillsMcpAvailabilityGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMcpAvailabilityGetOutputToJSON serializes a SkillsMcpAvailabilityGetOutput to JSON.
func MapSkillsMcpAvailabilityGetOutputToJSON(v *SkillsMcpAvailabilityGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
