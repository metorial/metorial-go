package messages

import (
	"encoding/json"
	"time"
)

// SessionsMessagesListOutputItemsHierarchy - Where this message sits in the reply chain.
type SessionsMessagesListOutputItemsHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - `child` when the message replies to another message, `parent` otherwise.
	Type string `json:"type"`
	// ParentMessageId - Message this replies to. `null` for parent messages.
	ParentMessageId *string `json:"parent_message_id,omitempty"`
	// ChildMessageIds - Messages replying to this one.
	ChildMessageIds []string `json:"child_message_ids"`
}

// SessionsMessagesListOutputItemsTransportMcp represents the sessions messages list output items transport mcp type.
type SessionsMessagesListOutputItemsTransportMcp struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id any `json:"id"`
	// ProtocolVersion - Negotiated MCP version, or `unknown`.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsMessagesListOutputItemsTransportToolCall represents the sessions messages list output items transport tool call type.
type SessionsMessagesListOutputItemsTransportToolCall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsMessagesListOutputItemsTransport - How this message was delivered.
type SessionsMessagesListOutputItemsTransport struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Which channel carried the message.
	Type     string                                            `json:"type"`
	Mcp      *SessionsMessagesListOutputItemsTransportMcp      `json:"mcp,omitempty"`
	ToolCall *SessionsMessagesListOutputItemsTransportToolCall `json:"tool_call,omitempty"`
}

// SessionsMessagesListOutputItemsToolCallSenderParticipantData - Extra details the participant connected with.
type SessionsMessagesListOutputItemsToolCallSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesListOutputItemsToolCallSenderParticipant represents the sessions messages list output items tool call sender participant type.
type SessionsMessagesListOutputItemsToolCallSenderParticipant struct {
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
	Data SessionsMessagesListOutputItemsToolCallSenderParticipantData `json:"data"`
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

// SessionsMessagesListOutputItemsToolCallResponderParticipantData - Extra details the participant connected with.
type SessionsMessagesListOutputItemsToolCallResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesListOutputItemsToolCallResponderParticipant represents the sessions messages list output items tool call responder participant type.
type SessionsMessagesListOutputItemsToolCallResponderParticipant struct {
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
	Data SessionsMessagesListOutputItemsToolCallResponderParticipantData `json:"data"`
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

// SessionsMessagesListOutputItemsToolCallToolInputSchema represents the sessions messages list output items tool call tool input schema type.
type SessionsMessagesListOutputItemsToolCallToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// SessionsMessagesListOutputItemsToolCallToolOutputSchema represents the sessions messages list output items tool call tool output schema type.
type SessionsMessagesListOutputItemsToolCallToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// SessionsMessagesListOutputItemsToolCallToolTags represents the sessions messages list output items tool call tool tags type.
type SessionsMessagesListOutputItemsToolCallToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionsMessagesListOutputItemsToolCallTool represents the sessions messages list output items tool call tool type.
type SessionsMessagesListOutputItemsToolCallTool struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - Key used to call the tool over MCP.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What the tool can do.
	Capabilities map[string]any `json:"capabilities"`
	// Constraints - Rules for using the tool.
	Constraints []string `json:"constraints"`
	// Instructions - Guidance for calling the tool correctly.
	Instructions []string                                                 `json:"instructions"`
	InputSchema  *SessionsMessagesListOutputItemsToolCallToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionsMessagesListOutputItemsToolCallToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionsMessagesListOutputItemsToolCallToolTags         `json:"tags,omitempty"`
	// ToolCategoryKeys - Keys of the tool categories this tool belongs to
	ToolCategoryKeys []string `json:"tool_category_keys"`
	// SpecificationId - Specification this tool belongs to.
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider this tool belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsMessagesListOutputItemsToolCallError represents the sessions messages list output items tool call error type.
type SessionsMessagesListOutputItemsToolCallError struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Additional error details.
	Data map[string]any `json:"data"`
	// Status - Indicates whether the error is still being processed or has been fully processed and grouped.
	Status string `json:"status"`
	// SessionId - Session the error occurred in.
	SessionId string `json:"session_id"`
	// ProviderRunId - Provider run the error belongs to. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// ConnectionId - Connection the error belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// GroupId - Group this error belongs to. `null` while the error is still processing.
	GroupId *string `json:"group_id,omitempty"`
	// SimilarErrorCount - How many errors are in this group.
	SimilarErrorCount float64 `json:"similar_error_count"`
	// IssueId - Provider issue this error is aggregated into. `null` when there is none.
	IssueId *string `json:"issue_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsMessagesListOutputItemsToolCall represents the sessions messages list output items tool call type.
type SessionsMessagesListOutputItemsToolCall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ToolKey - The key identifying the tool that was called.
	ToolKey string `json:"tool_key"`
	// Type - Always `tool_call`.
	Type string `json:"type"`
	// Status - `waiting_for_response` until the tool replies.
	Status string `json:"status"`
	// Source - Whether the client or the provider made the call.
	Source string `json:"source"`
	// Transport - Transport protocol used
	Transport string `json:"transport"`
	// SessionId - Session the tool was called in.
	SessionId string `json:"session_id"`
	// MessageId - Message that records this call.
	MessageId string `json:"message_id"`
	// SessionProviderId - Session provider that handled the call. `null` when there is none.
	SessionProviderId *string `json:"session_provider_id,omitempty"`
	// ConnectionId - Connection the call belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// ProviderRunId - Provider run that produced the call. `null` when there is none.
	ProviderRunId        *string                                                      `json:"provider_run_id,omitempty"`
	SenderParticipant    *SessionsMessagesListOutputItemsToolCallSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *SessionsMessagesListOutputItemsToolCallResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 SessionsMessagesListOutputItemsToolCallTool                  `json:"tool"`
	Error                *SessionsMessagesListOutputItemsToolCallError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsMessagesListOutputItemsSenderParticipantData - Extra details the participant connected with.
type SessionsMessagesListOutputItemsSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesListOutputItemsSenderParticipant represents the sessions messages list output items sender participant type.
type SessionsMessagesListOutputItemsSenderParticipant struct {
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
	Data SessionsMessagesListOutputItemsSenderParticipantData `json:"data"`
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

// SessionsMessagesListOutputItemsResponderParticipantData - Extra details the participant connected with.
type SessionsMessagesListOutputItemsResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesListOutputItemsResponderParticipant represents the sessions messages list output items responder participant type.
type SessionsMessagesListOutputItemsResponderParticipant struct {
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
	Data SessionsMessagesListOutputItemsResponderParticipantData `json:"data"`
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

// SessionsMessagesListOutputItemsError represents the sessions messages list output items error type.
type SessionsMessagesListOutputItemsError struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Additional error details.
	Data map[string]any `json:"data"`
	// Status - Indicates whether the error is still being processed or has been fully processed and grouped.
	Status string `json:"status"`
	// SessionId - Session the error occurred in.
	SessionId string `json:"session_id"`
	// ProviderRunId - Provider run the error belongs to. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// ConnectionId - Connection the error belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// GroupId - Group this error belongs to. `null` while the error is still processing.
	GroupId *string `json:"group_id,omitempty"`
	// SimilarErrorCount - How many errors are in this group.
	SimilarErrorCount float64 `json:"similar_error_count"`
	// IssueId - Provider issue this error is aggregated into. `null` when there is none.
	IssueId *string `json:"issue_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsMessagesListOutputItems represents the sessions messages list output items type.
type SessionsMessagesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What the message carries: a tool call, connection control, or MCP protocol data.
	Type string `json:"type"`
	// Status - `waiting_for_response` until a reply arrives.
	Status string `json:"status"`
	// Source - Whether the client or the provider sent the message.
	Source string `json:"source"`
	// SessionId - Session the message was sent in.
	SessionId string `json:"session_id"`
	// SessionProviderId - Session provider that handled the message. `null` when there is none.
	SessionProviderId *string `json:"session_provider_id,omitempty"`
	// ConnectionId - Connection the message belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// ProviderRunId - Provider run that produced the message. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// Hierarchy - Where this message sits in the reply chain.
	Hierarchy SessionsMessagesListOutputItemsHierarchy `json:"hierarchy"`
	// Transport - How this message was delivered.
	Transport SessionsMessagesListOutputItemsTransport `json:"transport"`
	// Input - Input message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output               *map[string]any                                      `json:"output,omitempty"`
	ToolCall             *SessionsMessagesListOutputItemsToolCall             `json:"tool_call,omitempty"`
	SenderParticipant    SessionsMessagesListOutputItemsSenderParticipant     `json:"sender_participant"`
	ResponderParticipant *SessionsMessagesListOutputItemsResponderParticipant `json:"responder_participant,omitempty"`
	Error                *SessionsMessagesListOutputItemsError                `json:"error,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsMessagesListOutputPagination represents the sessions messages list output pagination type.
type SessionsMessagesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsMessagesListOutput represents the sessions messages list output type.
type SessionsMessagesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []SessionsMessagesListOutputItems    `json:"items"`
	Pagination SessionsMessagesListOutputPagination `json:"pagination"`
}

// MapSessionsMessagesListOutputFromJSON deserializes JSON data into a SessionsMessagesListOutput.
func MapSessionsMessagesListOutputFromJSON(data []byte) (*SessionsMessagesListOutput, error) {
	var v SessionsMessagesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsMessagesListOutputToJSON serializes a SessionsMessagesListOutput to JSON.
func MapSessionsMessagesListOutputToJSON(v *SessionsMessagesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsMessagesListQueryCreatedAt - Only return results whose session message creation time is in this range. Set `gt`, `lt`, or both.
type SessionsMessagesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsMessagesListQueryUpdatedAt - Only return results whose session message last update time is in this range. Set `gt`, `lt`, or both.
type SessionsMessagesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsMessagesListQuery represents the sessions messages list query type.
type SessionsMessagesListQuery struct {
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
	// Type - Filter by message type(s)
	Type *any `json:"type,omitempty"`
	// Source - Filter by message source(s)
	Source *any `json:"source,omitempty"`
	// Hierarchy - Filter by message hierarchy
	Hierarchy *any `json:"hierarchy,omitempty"`
	// Status - Filter by message status(es)
	Status *any `json:"status,omitempty"`
	// Id - Filter by message ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Filter by provider run ID(s)
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// ErrorId - Filter by error ID(s)
	ErrorId *any `json:"error_id,omitempty"`
	// ParticipantId - Filter by participant ID(s)
	ParticipantId *any `json:"participant_id,omitempty"`
	// ParentMessageId - Filter by parent message ID(s)
	ParentMessageId *any `json:"parent_message_id,omitempty"`
	// CreatedAt - Only return results whose session message creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsMessagesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session message last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsMessagesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionsMessagesListQueryFromJSON deserializes JSON data into a SessionsMessagesListQuery.
func MapSessionsMessagesListQueryFromJSON(data []byte) (*SessionsMessagesListQuery, error) {
	var v SessionsMessagesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsMessagesListQueryToJSON serializes a SessionsMessagesListQuery to JSON.
func MapSessionsMessagesListQueryToJSON(v *SessionsMessagesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
