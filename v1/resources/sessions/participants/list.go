package participants

import (
	"encoding/json"
	"time"
)

// SessionsParticipantsListOutputItemsData - Extra details the participant connected with.
type SessionsParticipantsListOutputItemsData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsParticipantsListOutputItems represents the sessions participants list output items type.
type SessionsParticipantsListOutputItems struct {
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
	Data SessionsParticipantsListOutputItemsData `json:"data"`
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

// SessionsParticipantsListOutputPagination represents the sessions participants list output pagination type.
type SessionsParticipantsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsParticipantsListOutput represents the sessions participants list output type.
type SessionsParticipantsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                   `json:"object"`
	Items      []SessionsParticipantsListOutputItems    `json:"items"`
	Pagination SessionsParticipantsListOutputPagination `json:"pagination"`
}

// MapSessionsParticipantsListOutputFromJSON deserializes JSON data into a SessionsParticipantsListOutput.
func MapSessionsParticipantsListOutputFromJSON(data []byte) (*SessionsParticipantsListOutput, error) {
	var v SessionsParticipantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsParticipantsListOutputToJSON serializes a SessionsParticipantsListOutput to JSON.
func MapSessionsParticipantsListOutputToJSON(v *SessionsParticipantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsParticipantsListQueryCreatedAt - Only return results whose session participant creation time is in this range. Set `gt`, `lt`, or both.
type SessionsParticipantsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsParticipantsListQueryUpdatedAt - Only return results whose session participant last update time is in this range. Set `gt`, `lt`, or both.
type SessionsParticipantsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsParticipantsListQuery represents the sessions participants list query type.
type SessionsParticipantsListQuery struct {
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
	// Type - Filter by participant type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by participant ID(s)
	Id *any `json:"id,omitempty"`
	// AgentId - Filter by connected agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by connected identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by connected consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by connected identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// AgentInstanceId - Filter by connected agent instance ID(s)
	AgentInstanceId *any `json:"agent_instance_id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// CreatedAt - Only return results whose session participant creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsParticipantsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session participant last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsParticipantsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionsParticipantsListQueryFromJSON deserializes JSON data into a SessionsParticipantsListQuery.
func MapSessionsParticipantsListQueryFromJSON(data []byte) (*SessionsParticipantsListQuery, error) {
	var v SessionsParticipantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsParticipantsListQueryToJSON serializes a SessionsParticipantsListQuery to JSON.
func MapSessionsParticipantsListQueryToJSON(v *SessionsParticipantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
