package devices

import (
	"encoding/json"
	"time"
)

// ConsumersDevicesListOutputItemsClients represents the consumers devices list output items clients type.
type ConsumersDevicesListOutputItemsClients struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Key - Which AI tool this entry is for.
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

// ConsumersDevicesListOutputItems represents the consumers devices list output items type.
type ConsumersDevicesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
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
	Clients []ConsumersDevicesListOutputItemsClients `json:"clients"`
	// LastSeenAt - When the device last checked in.
	LastSeenAt time.Time `json:"last_seen_at"`
	// RevokedAt - When the device was revoked. Revoked devices stay listable.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConsumersDevicesListOutputPagination represents the consumers devices list output pagination type.
type ConsumersDevicesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConsumersDevicesListOutput represents the consumers devices list output type.
type ConsumersDevicesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []ConsumersDevicesListOutputItems    `json:"items"`
	Pagination ConsumersDevicesListOutputPagination `json:"pagination"`
}

// MapConsumersDevicesListOutputFromJSON deserializes JSON data into a ConsumersDevicesListOutput.
func MapConsumersDevicesListOutputFromJSON(data []byte) (*ConsumersDevicesListOutput, error) {
	var v ConsumersDevicesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersDevicesListOutputToJSON serializes a ConsumersDevicesListOutput to JSON.
func MapConsumersDevicesListOutputToJSON(v *ConsumersDevicesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersDevicesListQuery represents the consumers devices list query type.
type ConsumersDevicesListQuery struct {
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
	// ConsumerProfileId - Only return devices enrolled by these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Only return devices enrolled by any profile of these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// Status - Only return devices with these statuses. Defaults to active.
	Status *any `json:"status,omitempty"`
}

// MapConsumersDevicesListQueryFromJSON deserializes JSON data into a ConsumersDevicesListQuery.
func MapConsumersDevicesListQueryFromJSON(data []byte) (*ConsumersDevicesListQuery, error) {
	var v ConsumersDevicesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersDevicesListQueryToJSON serializes a ConsumersDevicesListQuery to JSON.
func MapConsumersDevicesListQueryToJSON(v *ConsumersDevicesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
