package events

import (
	"encoding/json"
	"time"
)

// SessionsEventsListOutputItemsConnectionUsage - Usage statistics
type SessionsEventsListOutputItemsConnectionUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsEventsListOutputItemsConnectionMcp - MCP connection details
type SessionsEventsListOutputItemsConnectionMcp struct {
	// Capabilities - Capabilities the client reported.
	Capabilities map[string]any `json:"capabilities"`
	// ProtocolVersion - Negotiated MCP version.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsEventsListOutputItemsConnectionParticipantData - Extra details the participant connected with.
type SessionsEventsListOutputItemsConnectionParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsListOutputItemsConnectionParticipant represents the sessions events list output items connection participant type.
type SessionsEventsListOutputItemsConnectionParticipant struct {
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
	Data SessionsEventsListOutputItemsConnectionParticipantData `json:"data"`
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

// SessionsEventsListOutputItemsConnection represents the sessions events list output items connection type.
type SessionsEventsListOutputItemsConnection struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ConnectionState - Whether the connection is currently open.
	ConnectionState string `json:"connection_state"`
	// Transport - Transport protocol used
	Transport string `json:"transport"`
	// Usage - Usage statistics
	Usage SessionsEventsListOutputItemsConnectionUsage `json:"usage"`
	// Mcp - MCP connection details
	Mcp *SessionsEventsListOutputItemsConnectionMcp `json:"mcp,omitempty"`
	// SessionId - Session this connection belongs to.
	SessionId   string                                              `json:"session_id"`
	Participant *SessionsEventsListOutputItemsConnectionParticipant `json:"participant,omitempty"`
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

// SessionsEventsListOutputItemsProviderRun represents the sessions events list output items provider run type.
type SessionsEventsListOutputItemsProviderRun struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `running` while the provider handles messages. `stopped` once it finishes.
	Status string `json:"status"`
	// SessionId - Session the run belongs to.
	SessionId string `json:"session_id"`
	// SessionProviderId - Session provider that handled the run.
	SessionProviderId string `json:"session_provider_id"`
	// ProviderId - Provider that ran.
	ProviderId string `json:"provider_id"`
	// ConnectionId - Connection the run used.
	ConnectionId string `json:"connection_id"`
	// CompletedAt - When the run finished. `null` while it is still running.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsEventsListOutputItemsMessageHierarchy - Where this message sits in the reply chain.
type SessionsEventsListOutputItemsMessageHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - `child` when the message replies to another message, `parent` otherwise.
	Type string `json:"type"`
	// ParentMessageId - Message this replies to. `null` for parent messages.
	ParentMessageId *string `json:"parent_message_id,omitempty"`
	// ChildMessageIds - Messages replying to this one.
	ChildMessageIds []string `json:"child_message_ids"`
}

// SessionsEventsListOutputItemsMessageTransportMcp represents the sessions events list output items message transport mcp type.
type SessionsEventsListOutputItemsMessageTransportMcp struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id any `json:"id"`
	// ProtocolVersion - Negotiated MCP version, or `unknown`.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsEventsListOutputItemsMessageTransportToolCall represents the sessions events list output items message transport tool call type.
type SessionsEventsListOutputItemsMessageTransportToolCall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsEventsListOutputItemsMessageTransport - How this message was delivered.
type SessionsEventsListOutputItemsMessageTransport struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Which channel carried the message.
	Type     string                                                 `json:"type"`
	Mcp      *SessionsEventsListOutputItemsMessageTransportMcp      `json:"mcp,omitempty"`
	ToolCall *SessionsEventsListOutputItemsMessageTransportToolCall `json:"tool_call,omitempty"`
}

// SessionsEventsListOutputItemsMessageToolCallSenderParticipantData - Extra details the participant connected with.
type SessionsEventsListOutputItemsMessageToolCallSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsListOutputItemsMessageToolCallSenderParticipant represents the sessions events list output items message tool call sender participant type.
type SessionsEventsListOutputItemsMessageToolCallSenderParticipant struct {
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
	Data SessionsEventsListOutputItemsMessageToolCallSenderParticipantData `json:"data"`
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

// SessionsEventsListOutputItemsMessageToolCallResponderParticipantData - Extra details the participant connected with.
type SessionsEventsListOutputItemsMessageToolCallResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsListOutputItemsMessageToolCallResponderParticipant represents the sessions events list output items message tool call responder participant type.
type SessionsEventsListOutputItemsMessageToolCallResponderParticipant struct {
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
	Data SessionsEventsListOutputItemsMessageToolCallResponderParticipantData `json:"data"`
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

// SessionsEventsListOutputItemsMessageToolCallToolInputSchema represents the sessions events list output items message tool call tool input schema type.
type SessionsEventsListOutputItemsMessageToolCallToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// SessionsEventsListOutputItemsMessageToolCallToolOutputSchema represents the sessions events list output items message tool call tool output schema type.
type SessionsEventsListOutputItemsMessageToolCallToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// SessionsEventsListOutputItemsMessageToolCallToolTags represents the sessions events list output items message tool call tool tags type.
type SessionsEventsListOutputItemsMessageToolCallToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionsEventsListOutputItemsMessageToolCallTool represents the sessions events list output items message tool call tool type.
type SessionsEventsListOutputItemsMessageToolCallTool struct {
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
	Instructions []string                                                      `json:"instructions"`
	InputSchema  *SessionsEventsListOutputItemsMessageToolCallToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionsEventsListOutputItemsMessageToolCallToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionsEventsListOutputItemsMessageToolCallToolTags         `json:"tags,omitempty"`
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

// SessionsEventsListOutputItemsMessageToolCallError represents the sessions events list output items message tool call error type.
type SessionsEventsListOutputItemsMessageToolCallError struct {
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

// SessionsEventsListOutputItemsMessageToolCall represents the sessions events list output items message tool call type.
type SessionsEventsListOutputItemsMessageToolCall struct {
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
	ProviderRunId        *string                                                           `json:"provider_run_id,omitempty"`
	SenderParticipant    *SessionsEventsListOutputItemsMessageToolCallSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *SessionsEventsListOutputItemsMessageToolCallResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 SessionsEventsListOutputItemsMessageToolCallTool                  `json:"tool"`
	Error                *SessionsEventsListOutputItemsMessageToolCallError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsListOutputItemsMessageSenderParticipantData - Extra details the participant connected with.
type SessionsEventsListOutputItemsMessageSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsListOutputItemsMessageSenderParticipant represents the sessions events list output items message sender participant type.
type SessionsEventsListOutputItemsMessageSenderParticipant struct {
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
	Data SessionsEventsListOutputItemsMessageSenderParticipantData `json:"data"`
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

// SessionsEventsListOutputItemsMessageResponderParticipantData - Extra details the participant connected with.
type SessionsEventsListOutputItemsMessageResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsListOutputItemsMessageResponderParticipant represents the sessions events list output items message responder participant type.
type SessionsEventsListOutputItemsMessageResponderParticipant struct {
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
	Data SessionsEventsListOutputItemsMessageResponderParticipantData `json:"data"`
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

// SessionsEventsListOutputItemsMessageError represents the sessions events list output items message error type.
type SessionsEventsListOutputItemsMessageError struct {
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

// SessionsEventsListOutputItemsMessage represents the sessions events list output items message type.
type SessionsEventsListOutputItemsMessage struct {
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
	Hierarchy SessionsEventsListOutputItemsMessageHierarchy `json:"hierarchy"`
	// Transport - How this message was delivered.
	Transport SessionsEventsListOutputItemsMessageTransport `json:"transport"`
	// Input - Input message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output               *map[string]any                                           `json:"output,omitempty"`
	ToolCall             *SessionsEventsListOutputItemsMessageToolCall             `json:"tool_call,omitempty"`
	SenderParticipant    SessionsEventsListOutputItemsMessageSenderParticipant     `json:"sender_participant"`
	ResponderParticipant *SessionsEventsListOutputItemsMessageResponderParticipant `json:"responder_participant,omitempty"`
	Error                *SessionsEventsListOutputItemsMessageError                `json:"error,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsListOutputItemsError represents the sessions events list output items error type.
type SessionsEventsListOutputItemsError struct {
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

// SessionsEventsListOutputItemsWarning represents the sessions events list output items warning type.
type SessionsEventsListOutputItemsWarning struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable warning code.
	Code string `json:"code"`
	// Message - Human-readable warning message.
	Message string `json:"message"`
	// Data - Additional warning details.
	Data map[string]any `json:"data"`
	// SessionId - Session the warning occurred in.
	SessionId string `json:"session_id"`
	// ConnectionId - Connection the warning belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsListOutputItems represents the sessions events list output items type.
type SessionsEventsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What happened in the session.
	Type string `json:"type"`
	// SessionId - Session the event belongs to.
	SessionId   string                                    `json:"session_id"`
	Connection  *SessionsEventsListOutputItemsConnection  `json:"connection,omitempty"`
	ProviderRun *SessionsEventsListOutputItemsProviderRun `json:"provider_run,omitempty"`
	Message     *SessionsEventsListOutputItemsMessage     `json:"message,omitempty"`
	Error       *SessionsEventsListOutputItemsError       `json:"error,omitempty"`
	Warning     *SessionsEventsListOutputItemsWarning     `json:"warning,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsListOutputPagination represents the sessions events list output pagination type.
type SessionsEventsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsEventsListOutput represents the sessions events list output type.
type SessionsEventsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []SessionsEventsListOutputItems    `json:"items"`
	Pagination SessionsEventsListOutputPagination `json:"pagination"`
}

// MapSessionsEventsListOutputFromJSON deserializes JSON data into a SessionsEventsListOutput.
func MapSessionsEventsListOutputFromJSON(data []byte) (*SessionsEventsListOutput, error) {
	var v SessionsEventsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsEventsListOutputToJSON serializes a SessionsEventsListOutput to JSON.
func MapSessionsEventsListOutputToJSON(v *SessionsEventsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsEventsListQueryCreatedAt - Only return results whose session event creation time is in this range. Set `gt`, `lt`, or both.
type SessionsEventsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsEventsListQueryUpdatedAt - Only return results whose session event last update time is in this range. Set `gt`, `lt`, or both.
type SessionsEventsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsEventsListQuery represents the sessions events list query type.
type SessionsEventsListQuery struct {
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
	// Type - Filter by event type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by session event ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Filter by provider run ID(s)
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionErrorId - Filter by session error ID(s)
	SessionErrorId *any `json:"session_error_id,omitempty"`
	// CreatedAt - Only return results whose session event creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsEventsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session event last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsEventsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionsEventsListQueryFromJSON deserializes JSON data into a SessionsEventsListQuery.
func MapSessionsEventsListQueryFromJSON(data []byte) (*SessionsEventsListQuery, error) {
	var v SessionsEventsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsEventsListQueryToJSON serializes a SessionsEventsListQuery to JSON.
func MapSessionsEventsListQueryToJSON(v *SessionsEventsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
