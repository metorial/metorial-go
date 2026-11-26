package events

import (
	"encoding/json"
	"time"
)

// SessionsEventsGetOutputConnectionUsage - Usage statistics
type SessionsEventsGetOutputConnectionUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsEventsGetOutputConnectionMcp - MCP connection details
type SessionsEventsGetOutputConnectionMcp struct {
	// Capabilities - Capabilities the client reported.
	Capabilities map[string]any `json:"capabilities"`
	// ProtocolVersion - Negotiated MCP version.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsEventsGetOutputConnectionParticipantData - Extra details the participant connected with.
type SessionsEventsGetOutputConnectionParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsGetOutputConnectionParticipant represents the sessions events get output connection participant type.
type SessionsEventsGetOutputConnectionParticipant struct {
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
	Data SessionsEventsGetOutputConnectionParticipantData `json:"data"`
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

// SessionsEventsGetOutputConnection represents the sessions events get output connection type.
type SessionsEventsGetOutputConnection struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ConnectionState - Whether the connection is currently open.
	ConnectionState string `json:"connection_state"`
	// Transport - Transport protocol used
	Transport string `json:"transport"`
	// Usage - Usage statistics
	Usage SessionsEventsGetOutputConnectionUsage `json:"usage"`
	// Mcp - MCP connection details
	Mcp *SessionsEventsGetOutputConnectionMcp `json:"mcp,omitempty"`
	// SessionId - Session this connection belongs to.
	SessionId   string                                        `json:"session_id"`
	Participant *SessionsEventsGetOutputConnectionParticipant `json:"participant,omitempty"`
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

// SessionsEventsGetOutputProviderRun represents the sessions events get output provider run type.
type SessionsEventsGetOutputProviderRun struct {
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

// SessionsEventsGetOutputMessageHierarchy - Where this message sits in the reply chain.
type SessionsEventsGetOutputMessageHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - `child` when the message replies to another message, `parent` otherwise.
	Type string `json:"type"`
	// ParentMessageId - Message this replies to. `null` for parent messages.
	ParentMessageId *string `json:"parent_message_id,omitempty"`
	// ChildMessageIds - Messages replying to this one.
	ChildMessageIds []string `json:"child_message_ids"`
}

// SessionsEventsGetOutputMessageTransportMcp represents the sessions events get output message transport mcp type.
type SessionsEventsGetOutputMessageTransportMcp struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id any `json:"id"`
	// ProtocolVersion - Negotiated MCP version, or `unknown`.
	ProtocolVersion string `json:"protocol_version"`
	// Transport - How the MCP connection is carried.
	Transport string `json:"transport"`
}

// SessionsEventsGetOutputMessageTransportToolCall represents the sessions events get output message transport tool call type.
type SessionsEventsGetOutputMessageTransportToolCall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsEventsGetOutputMessageTransport - How this message was delivered.
type SessionsEventsGetOutputMessageTransport struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Which channel carried the message.
	Type     string                                           `json:"type"`
	Mcp      *SessionsEventsGetOutputMessageTransportMcp      `json:"mcp,omitempty"`
	ToolCall *SessionsEventsGetOutputMessageTransportToolCall `json:"tool_call,omitempty"`
}

// SessionsEventsGetOutputMessageToolCallSenderParticipantData - Extra details the participant connected with.
type SessionsEventsGetOutputMessageToolCallSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsGetOutputMessageToolCallSenderParticipant represents the sessions events get output message tool call sender participant type.
type SessionsEventsGetOutputMessageToolCallSenderParticipant struct {
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
	Data SessionsEventsGetOutputMessageToolCallSenderParticipantData `json:"data"`
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

// SessionsEventsGetOutputMessageToolCallResponderParticipantData - Extra details the participant connected with.
type SessionsEventsGetOutputMessageToolCallResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsGetOutputMessageToolCallResponderParticipant represents the sessions events get output message tool call responder participant type.
type SessionsEventsGetOutputMessageToolCallResponderParticipant struct {
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
	Data SessionsEventsGetOutputMessageToolCallResponderParticipantData `json:"data"`
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

// SessionsEventsGetOutputMessageToolCallToolInputSchema represents the sessions events get output message tool call tool input schema type.
type SessionsEventsGetOutputMessageToolCallToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// SessionsEventsGetOutputMessageToolCallToolOutputSchema represents the sessions events get output message tool call tool output schema type.
type SessionsEventsGetOutputMessageToolCallToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// SessionsEventsGetOutputMessageToolCallToolTags represents the sessions events get output message tool call tool tags type.
type SessionsEventsGetOutputMessageToolCallToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// SessionsEventsGetOutputMessageToolCallTool represents the sessions events get output message tool call tool type.
type SessionsEventsGetOutputMessageToolCallTool struct {
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
	Instructions []string                                                `json:"instructions"`
	InputSchema  *SessionsEventsGetOutputMessageToolCallToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *SessionsEventsGetOutputMessageToolCallToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *SessionsEventsGetOutputMessageToolCallToolTags         `json:"tags,omitempty"`
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

// SessionsEventsGetOutputMessageToolCallError represents the sessions events get output message tool call error type.
type SessionsEventsGetOutputMessageToolCallError struct {
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

// SessionsEventsGetOutputMessageToolCall represents the sessions events get output message tool call type.
type SessionsEventsGetOutputMessageToolCall struct {
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
	ProviderRunId        *string                                                     `json:"provider_run_id,omitempty"`
	SenderParticipant    *SessionsEventsGetOutputMessageToolCallSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *SessionsEventsGetOutputMessageToolCallResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 SessionsEventsGetOutputMessageToolCallTool                  `json:"tool"`
	Error                *SessionsEventsGetOutputMessageToolCallError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsGetOutputMessageSenderParticipantData - Extra details the participant connected with.
type SessionsEventsGetOutputMessageSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsGetOutputMessageSenderParticipant represents the sessions events get output message sender participant type.
type SessionsEventsGetOutputMessageSenderParticipant struct {
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
	Data SessionsEventsGetOutputMessageSenderParticipantData `json:"data"`
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

// SessionsEventsGetOutputMessageResponderParticipantData - Extra details the participant connected with.
type SessionsEventsGetOutputMessageResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// SessionsEventsGetOutputMessageResponderParticipant represents the sessions events get output message responder participant type.
type SessionsEventsGetOutputMessageResponderParticipant struct {
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
	Data SessionsEventsGetOutputMessageResponderParticipantData `json:"data"`
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

// SessionsEventsGetOutputMessageError represents the sessions events get output message error type.
type SessionsEventsGetOutputMessageError struct {
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

// SessionsEventsGetOutputMessage represents the sessions events get output message type.
type SessionsEventsGetOutputMessage struct {
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
	Hierarchy SessionsEventsGetOutputMessageHierarchy `json:"hierarchy"`
	// Transport - How this message was delivered.
	Transport SessionsEventsGetOutputMessageTransport `json:"transport"`
	// Input - Input message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output message data. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output               *map[string]any                                     `json:"output,omitempty"`
	ToolCall             *SessionsEventsGetOutputMessageToolCall             `json:"tool_call,omitempty"`
	SenderParticipant    SessionsEventsGetOutputMessageSenderParticipant     `json:"sender_participant"`
	ResponderParticipant *SessionsEventsGetOutputMessageResponderParticipant `json:"responder_participant,omitempty"`
	Error                *SessionsEventsGetOutputMessageError                `json:"error,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsEventsGetOutputError represents the sessions events get output error type.
type SessionsEventsGetOutputError struct {
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

// SessionsEventsGetOutputWarning represents the sessions events get output warning type.
type SessionsEventsGetOutputWarning struct {
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

// SessionsEventsGetOutput represents the sessions events get output type.
type SessionsEventsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What happened in the session.
	Type string `json:"type"`
	// SessionId - Session the event belongs to.
	SessionId   string                              `json:"session_id"`
	Connection  *SessionsEventsGetOutputConnection  `json:"connection,omitempty"`
	ProviderRun *SessionsEventsGetOutputProviderRun `json:"provider_run,omitempty"`
	Message     *SessionsEventsGetOutputMessage     `json:"message,omitempty"`
	Error       *SessionsEventsGetOutputError       `json:"error,omitempty"`
	Warning     *SessionsEventsGetOutputWarning     `json:"warning,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapSessionsEventsGetOutputFromJSON deserializes JSON data into a SessionsEventsGetOutput.
func MapSessionsEventsGetOutputFromJSON(data []byte) (*SessionsEventsGetOutput, error) {
	var v SessionsEventsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsEventsGetOutputToJSON serializes a SessionsEventsGetOutput to JSON.
func MapSessionsEventsGetOutputToJSON(v *SessionsEventsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
