package toolcalls

import (
	"encoding/json"
	"time"
)

// ToolCallsCreateOutputSenderParticipantData - Extra details the participant connected with.
type ToolCallsCreateOutputSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// ToolCallsCreateOutputSenderParticipant represents the tool calls create output sender participant type.
type ToolCallsCreateOutputSenderParticipant struct {
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
	Data ToolCallsCreateOutputSenderParticipantData `json:"data"`
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

// ToolCallsCreateOutputResponderParticipantData - Extra details the participant connected with.
type ToolCallsCreateOutputResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// ToolCallsCreateOutputResponderParticipant represents the tool calls create output responder participant type.
type ToolCallsCreateOutputResponderParticipant struct {
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
	Data ToolCallsCreateOutputResponderParticipantData `json:"data"`
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

// ToolCallsCreateOutputToolInputSchema represents the tool calls create output tool input schema type.
type ToolCallsCreateOutputToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ToolCallsCreateOutputToolOutputSchema represents the tool calls create output tool output schema type.
type ToolCallsCreateOutputToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ToolCallsCreateOutputToolTags represents the tool calls create output tool tags type.
type ToolCallsCreateOutputToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ToolCallsCreateOutputTool represents the tool calls create output tool type.
type ToolCallsCreateOutputTool struct {
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
	Instructions []string                               `json:"instructions"`
	InputSchema  *ToolCallsCreateOutputToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ToolCallsCreateOutputToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *ToolCallsCreateOutputToolTags         `json:"tags,omitempty"`
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

// ToolCallsCreateOutputError represents the tool calls create output error type.
type ToolCallsCreateOutputError struct {
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

// ToolCallsCreateOutput represents the tool calls create output type.
type ToolCallsCreateOutput struct {
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
	ProviderRunId        *string                                    `json:"provider_run_id,omitempty"`
	SenderParticipant    *ToolCallsCreateOutputSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *ToolCallsCreateOutputResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 ToolCallsCreateOutputTool                  `json:"tool"`
	Error                *ToolCallsCreateOutputError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapToolCallsCreateOutputFromJSON deserializes JSON data into a ToolCallsCreateOutput.
func MapToolCallsCreateOutputFromJSON(data []byte) (*ToolCallsCreateOutput, error) {
	var v ToolCallsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapToolCallsCreateOutputToJSON serializes a ToolCallsCreateOutput to JSON.
func MapToolCallsCreateOutputToJSON(v *ToolCallsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ToolCallsCreateBody represents the tool calls create body type.
type ToolCallsCreateBody struct {
	// ToolId - Tool to call, by ID or key.
	ToolId string `json:"tool_id"`
	// Input - Input to pass to the tool.
	Input map[string]any `json:"input"`
	// Metadata - Custom data to identify or organize the tool call.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SessionId - Session to run the tool call in.
	SessionId string `json:"session_id"`
}

// MapToolCallsCreateBodyFromJSON deserializes JSON data into a ToolCallsCreateBody.
func MapToolCallsCreateBodyFromJSON(data []byte) (*ToolCallsCreateBody, error) {
	var v ToolCallsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapToolCallsCreateBodyToJSON serializes a ToolCallsCreateBody to JSON.
func MapToolCallsCreateBodyToJSON(v *ToolCallsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
