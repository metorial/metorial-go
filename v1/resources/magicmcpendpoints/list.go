package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsListOutputItems represents the magic mcp endpoints list output items type.
type MagicMcpEndpointsListOutputItems struct {
	Object            string           `json:"object"`
	Id                string           `json:"id"`
	Status            string           `json:"status"`
	Slug              string           `json:"slug"`
	Url               string           `json:"url"`
	ConsumerProfileId *string          `json:"consumer_profile_id,omitempty"`
	SessionTemplateId *string          `json:"session_template_id,omitempty"`
	SessionId         *string          `json:"session_id,omitempty"`
	Servers           []map[string]any `json:"servers"`
	Name              *string          `json:"name,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Metadata          map[string]any   `json:"metadata"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

// MagicMcpEndpointsListOutputPagination represents the magic mcp endpoints list output pagination type.
type MagicMcpEndpointsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpEndpointsListOutput represents the magic mcp endpoints list output type.
type MagicMcpEndpointsListOutput struct {
	Items      []MagicMcpEndpointsListOutputItems    `json:"items"`
	Pagination MagicMcpEndpointsListOutputPagination `json:"pagination"`
}

// MapMagicMcpEndpointsListOutputFromJSON deserializes JSON data into a MagicMcpEndpointsListOutput.
func MapMagicMcpEndpointsListOutputFromJSON(data []byte) (*MagicMcpEndpointsListOutput, error) {
	var v MagicMcpEndpointsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsListOutputToJSON serializes a MagicMcpEndpointsListOutput to JSON.
func MapMagicMcpEndpointsListOutputToJSON(v *MagicMcpEndpointsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsListQuery represents the magic mcp endpoints list query type.
type MagicMcpEndpointsListQuery struct {
	Limit            *float64 `json:"limit,omitempty"`
	After            *string  `json:"after,omitempty"`
	Before           *string  `json:"before,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	Order            *string  `json:"order,omitempty"`
	Status           *any     `json:"status,omitempty"`
	MagicMcpServerId *any     `json:"magic_mcp_server_id,omitempty"`
	Search           *string  `json:"search,omitempty"`
}

// MapMagicMcpEndpointsListQueryFromJSON deserializes JSON data into a MagicMcpEndpointsListQuery.
func MapMagicMcpEndpointsListQueryFromJSON(data []byte) (*MagicMcpEndpointsListQuery, error) {
	var v MagicMcpEndpointsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsListQueryToJSON serializes a MagicMcpEndpointsListQuery to JSON.
func MapMagicMcpEndpointsListQueryToJSON(v *MagicMcpEndpointsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
