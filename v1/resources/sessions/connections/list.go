package connections

import (
	"encoding/json"
	"time"
)

// SessionsConnectionsListOutputItemsUsage - Usage statistics
type SessionsConnectionsListOutputItemsUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsConnectionsListOutputItemsMcp - MCP connection details
type SessionsConnectionsListOutputItemsMcp struct {
	// Capabilities - Capabilities the client reported.
	Capabilities map[string]any `json:"capabilities"`
	// ProtocolVersion - Negotiated MCP version.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsConnectionsListOutputItemsParticipantData - Extra details the participant connected with.
type SessionsConnectionsListOutputItemsParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsConnectionsListOutputItemsParticipant represents the sessions connections list output items participant type.
type SessionsConnectionsListOutputItemsParticipant struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What kind of actor the participant is. Older client types report as `agent`.
	Type string `json:"type"`
	// Identifier - How the participant identifies itself.
	Identifier string `json:"identifier"`
	// Name - Display name.
	Name string `json:"name"`
	// Data - Extra details the participant connected with.
	Data SessionsConnectionsListOutputItemsParticipantData `json:"data"`
	// ProviderId - Provider the participant belongs to. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// ConnectionType - How the participant connected.
	ConnectionType *string `json:"connection_type,omitempty"`
	// AgentId - Agent the participant belongs to. `null` when there is none.
	AgentId *string `json:"agent_id,omitempty"`
	// AgentInstanceId - Agent instance the participant runs as. `null` when there is none.
	AgentInstanceId *string `json:"agent_instance_id,omitempty"`
	// IdentityActorId - Identity actor the participant acts as. `null` when there is none.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the participant acts as. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// AgentActorId - Actor record for the agent. `null` when there is none.
	AgentActorId *string `json:"agent_actor_id,omitempty"`
	// AgentClientId - Client record for the agent. `null` when there is none.
	AgentClientId *string `json:"agent_client_id,omitempty"`
	// ConsumerId - Consumer the participant acts for. `null` when there is none.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsConnectionsListOutputItems represents the sessions connections list output items type.
type SessionsConnectionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ConnectionState - Whether the connection is currently open.
	ConnectionState string `json:"connection_state"`
	// Transport - Transport protocol used
	Transport string `json:"transport"`
	// Usage - Usage statistics
	Usage SessionsConnectionsListOutputItemsUsage `json:"usage"`
	// Mcp - MCP connection details
	Mcp *SessionsConnectionsListOutputItemsMcp `json:"mcp,omitempty"`
	// SessionId - Session this connection belongs to.
	SessionId   string                                         `json:"session_id"`
	Participant *SessionsConnectionsListOutputItemsParticipant `json:"participant,omitempty"`
	// HasErrors - `true` when the connection has errors.
	HasErrors bool `json:"has_errors"`
	// HasWarnings - `true` when the connection has warnings.
	HasWarnings bool `json:"has_warnings"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// LastMessageAt - When the connection last exchanged a message.
	LastMessageAt time.Time `json:"last_message_at"`
	// LastActiveAt - Last time the connection showed activity.
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
}

// SessionsConnectionsListOutputPagination represents the sessions connections list output pagination type.
type SessionsConnectionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsConnectionsListOutput represents the sessions connections list output type.
type SessionsConnectionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []SessionsConnectionsListOutputItems    `json:"items"`
	Pagination SessionsConnectionsListOutputPagination `json:"pagination"`
}

// MapSessionsConnectionsListOutputFromJSON deserializes JSON data into a SessionsConnectionsListOutput.
func MapSessionsConnectionsListOutputFromJSON(data []byte) (*SessionsConnectionsListOutput, error) {
	var v SessionsConnectionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsConnectionsListOutputToJSON serializes a SessionsConnectionsListOutput to JSON.
func MapSessionsConnectionsListOutputToJSON(v *SessionsConnectionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsConnectionsListQueryCreatedAt - Only return results whose session connection creation time is in this range. Set `gt`, `lt`, or both.
type SessionsConnectionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsConnectionsListQueryUpdatedAt - Only return results whose session connection last update time is in this range. Set `gt`, `lt`, or both.
type SessionsConnectionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsConnectionsListQuery represents the sessions connections list query type.
type SessionsConnectionsListQuery struct {
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
	// Status - Filter by connection status
	Status *any `json:"status,omitempty"`
	// ConnectionState - Filter by connection state
	ConnectionState *any `json:"connection_state,omitempty"`
	// Id - Filter by session connection ID(s)
	Id *any `json:"id,omitempty"`
	// AgentId - Filter by connected agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by connected identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerProfileId - Filter by the selected consumer profile ID or IDs.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Filter by connected consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by connected identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// AgentInstanceId - Filter by connected agent instance ID(s)
	AgentInstanceId *any `json:"agent_instance_id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ParticipantId - Filter by participant ID(s)
	ParticipantId *any `json:"participant_id,omitempty"`
	// CreatedAt - Only return results whose session connection creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsConnectionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session connection last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsConnectionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionsConnectionsListQueryFromJSON deserializes JSON data into a SessionsConnectionsListQuery.
func MapSessionsConnectionsListQueryFromJSON(data []byte) (*SessionsConnectionsListQuery, error) {
	var v SessionsConnectionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsConnectionsListQueryToJSON serializes a SessionsConnectionsListQuery to JSON.
func MapSessionsConnectionsListQueryToJSON(v *SessionsConnectionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
