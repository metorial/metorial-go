package devices

import (
	"encoding/json"
	"time"
)

// ConsumersDevicesGetOutputClients represents the consumers devices get output clients type.
type ConsumersDevicesGetOutputClients struct {
	Object string `json:"object"`
	// Key - The AI tool identifier.
	Key string `json:"key"`
	// Name - The AI tool name.
	Name string `json:"name"`
	// Detected - Whether the device found this AI tool at its last check-in.
	Detected bool `json:"detected"`
	// Installable - Whether Metorial Desktop can configure this AI tool on the device.
	Installable bool `json:"installable"`
	// SupportsSkills - Whether Metorial Desktop can link skills into this AI tool.
	SupportsSkills bool `json:"supports_skills"`
	// Enabled - Whether Metorial Desktop keeps this AI tool configured. Detected tools are enabled until turned off.
	Enabled bool `json:"enabled"`
	// ExcludedMagicMcpServerIds - Magic MCP servers not installed into this AI tool. Servers not listed here are installed, including ones granted later.
	ExcludedMagicMcpServerIds []string `json:"excluded_magic_mcp_server_ids"`
	// ExcludedSkillIds - Skills not linked into this AI tool. Skills not listed here are linked, including ones granted later.
	ExcludedSkillIds []string `json:"excluded_skill_ids"`
}

// ConsumersDevicesGetOutput represents the consumers devices get output type.
type ConsumersDevicesGetOutput struct {
	Object string `json:"object"`
	// Id - The device ID.
	Id string `json:"id"`
	// Status - Whether the device is signed in (`active`), was signed out from the app (`signed_out`), or was revoked from the portal or dashboard (`revoked`).
	Status string `json:"status"`
	// ConsumerProfileId - The ID of the consumer profile that enrolled the device.
	ConsumerProfileId string `json:"consumer_profile_id"`
	// Name - The computer name reported by the device.
	Name string `json:"name"`
	// Os - The operating system reported by the device.
	Os string `json:"os"`
	// Arch - The CPU architecture reported by the device.
	Arch string `json:"arch"`
	// AppVersion - The Metorial Desktop version running on the device.
	AppVersion string `json:"app_version"`
	// Clients - The AI tools the device reported, with what each one receives.
	Clients []ConsumersDevicesGetOutputClients `json:"clients"`
	// LastSeenAt - When the device last checked in.
	LastSeenAt time.Time `json:"last_seen_at"`
	// RevokedAt - When the device was revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// CreatedAt - When the device was enrolled.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the device was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapConsumersDevicesGetOutputFromJSON deserializes JSON data into a ConsumersDevicesGetOutput.
func MapConsumersDevicesGetOutputFromJSON(data []byte) (*ConsumersDevicesGetOutput, error) {
	var v ConsumersDevicesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersDevicesGetOutputToJSON serializes a ConsumersDevicesGetOutput to JSON.
func MapConsumersDevicesGetOutputToJSON(v *ConsumersDevicesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
