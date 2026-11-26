package messages

import (
	"encoding/json"
	"time"
)

// SessionsMessagesGetOutputHierarchy - Where this message sits in the reply chain.
type SessionsMessagesGetOutputHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - `child` when the message replies to another message, `parent` otherwise.
	Type string `json:"type"`
	// ParentMessageId - Message this replies to. `null` for parent messages.
	ParentMessageId *string `json:"parent_message_id,omitempty"`
	// ChildMessageIds - Messages replying to this one.
	ChildMessageIds []string `json:"child_message_ids"`
}

// SessionsMessagesGetOutputTransportMcp represents the sessions messages get output transport mcp type.
type SessionsMessagesGetOutputTransportMcp struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id any `json:"id"`
	// ProtocolVersion - Negotiated MCP version, or `unknown`.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsMessagesGetOutputTransportToolCall represents the sessions messages get output transport tool call type.
type SessionsMessagesGetOutputTransportToolCall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsMessagesGetOutputTransport - How this message was delivered.
type SessionsMessagesGetOutputTransport struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Which channel carried the message.
	Type     string                                      `json:"type"`
	Mcp      *SessionsMessagesGetOutputTransportMcp      `json:"mcp,omitempty"`
	ToolCall *SessionsMessagesGetOutputTransportToolCall `json:"tool_call,omitempty"`
}

// SessionsMessagesGetOutputToolCallSenderParticipantData - Extra details the participant connected with.
type SessionsMessagesGetOutputToolCallSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesGetOutputToolCallSenderParticipant represents the sessions messages get output tool call sender participant type.
type SessionsMessagesGetOutputToolCallSenderParticipant struct {
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
	Data SessionsMessagesGetOutputToolCallSenderParticipantData `json:"data"`
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

// SessionsMessagesGetOutputToolCallResponderParticipantData - Extra details the participant connected with.
type SessionsMessagesGetOutputToolCallResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesGetOutputToolCallResponderParticipant represents the sessions messages get output tool call responder participant type.
type SessionsMessagesGetOutputToolCallResponderParticipant struct {
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
	Data SessionsMessagesGetOutputToolCallResponderParticipantData `json:"data"`
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

// SessionsMessagesGetOutputToolCallToolInputSchema represents the sessions messages get output tool call tool input schema type.
type SessionsMessagesGetOutputToolCallToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// SessionsMessagesGetOutputToolCallToolOutputSchema represents the sessions messages get output tool call tool output schema type.
type SessionsMessagesGetOutputToolCallToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// SessionsMessagesGetOutputToolCallToolTags represents the sessions messages get output tool call tool tags type.
type SessionsMessagesGetOutputToolCallToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionsMessagesGetOutputToolCallTool represents the sessions messages get output tool call tool type.
type SessionsMessagesGetOutputToolCallTool struct {
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
	Instructions []string                                           `json:"instructions"`
	InputSchema  *SessionsMessagesGetOutputToolCallToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionsMessagesGetOutputToolCallToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionsMessagesGetOutputToolCallToolTags         `json:"tags,omitempty"`
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

// SessionsMessagesGetOutputToolCallError represents the sessions messages get output tool call error type.
type SessionsMessagesGetOutputToolCallError struct {
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

// SessionsMessagesGetOutputToolCall represents the sessions messages get output tool call type.
type SessionsMessagesGetOutputToolCall struct {
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
	ProviderRunId        *string                                                `json:"provider_run_id,omitempty"`
	SenderParticipant    *SessionsMessagesGetOutputToolCallSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *SessionsMessagesGetOutputToolCallResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 SessionsMessagesGetOutputToolCallTool                  `json:"tool"`
	Error                *SessionsMessagesGetOutputToolCallError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsMessagesGetOutputSenderParticipantData - Extra details the participant connected with.
type SessionsMessagesGetOutputSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesGetOutputSenderParticipant represents the sessions messages get output sender participant type.
type SessionsMessagesGetOutputSenderParticipant struct {
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
	Data SessionsMessagesGetOutputSenderParticipantData `json:"data"`
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

// SessionsMessagesGetOutputResponderParticipantData - Extra details the participant connected with.
type SessionsMessagesGetOutputResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsMessagesGetOutputResponderParticipant represents the sessions messages get output responder participant type.
type SessionsMessagesGetOutputResponderParticipant struct {
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
	Data SessionsMessagesGetOutputResponderParticipantData `json:"data"`
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

// SessionsMessagesGetOutputError represents the sessions messages get output error type.
type SessionsMessagesGetOutputError struct {
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

// SessionsMessagesGetOutput represents the sessions messages get output type.
type SessionsMessagesGetOutput struct {
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
	Hierarchy SessionsMessagesGetOutputHierarchy `json:"hierarchy"`
	// Transport - How this message was delivered.
	Transport SessionsMessagesGetOutputTransport `json:"transport"`
	// Input - Input message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output               *map[string]any                                `json:"output,omitempty"`
	ToolCall             *SessionsMessagesGetOutputToolCall             `json:"tool_call,omitempty"`
	SenderParticipant    SessionsMessagesGetOutputSenderParticipant     `json:"sender_participant"`
	ResponderParticipant *SessionsMessagesGetOutputResponderParticipant `json:"responder_participant,omitempty"`
	Error                *SessionsMessagesGetOutputError                `json:"error,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapSessionsMessagesGetOutputFromJSON deserializes JSON data into a SessionsMessagesGetOutput.
func MapSessionsMessagesGetOutputFromJSON(data []byte) (*SessionsMessagesGetOutput, error) {
	var v SessionsMessagesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsMessagesGetOutputToJSON serializes a SessionsMessagesGetOutput to JSON.
func MapSessionsMessagesGetOutputToJSON(v *SessionsMessagesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
