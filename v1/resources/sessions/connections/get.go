package connections

import (
	"encoding/json"
	"time"
)

// SessionsConnectionsGetOutputUsage - Usage statistics
type SessionsConnectionsGetOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsConnectionsGetOutputMcp - MCP connection details
type SessionsConnectionsGetOutputMcp struct {
	// Capabilities - Capabilities the client reported.
	Capabilities map[string]any `json:"capabilities"`
	// ProtocolVersion - Negotiated MCP version.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsConnectionsGetOutputParticipantData - Extra details the participant connected with.
type SessionsConnectionsGetOutputParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsConnectionsGetOutputParticipant represents the sessions connections get output participant type.
type SessionsConnectionsGetOutputParticipant struct {
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
	Data SessionsConnectionsGetOutputParticipantData `json:"data"`
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

// SessionsConnectionsGetOutput represents the sessions connections get output type.
type SessionsConnectionsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ConnectionState - Whether the connection is currently open.
	ConnectionState string `json:"connection_state"`
	// Transport - Transport protocol used
	Transport string `json:"transport"`
	// Usage - Usage statistics
	Usage SessionsConnectionsGetOutputUsage `json:"usage"`
	// Mcp - MCP connection details
	Mcp *SessionsConnectionsGetOutputMcp `json:"mcp,omitempty"`
	// SessionId - Session this connection belongs to.
	SessionId   string                                   `json:"session_id"`
	Participant *SessionsConnectionsGetOutputParticipant `json:"participant,omitempty"`
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

// MapSessionsConnectionsGetOutputFromJSON deserializes JSON data into a SessionsConnectionsGetOutput.
func MapSessionsConnectionsGetOutputFromJSON(data []byte) (*SessionsConnectionsGetOutput, error) {
	var v SessionsConnectionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsConnectionsGetOutputToJSON serializes a SessionsConnectionsGetOutput to JSON.
func MapSessionsConnectionsGetOutputToJSON(v *SessionsConnectionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
