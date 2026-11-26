package toolcalls

import (
	"encoding/json"
	"time"
)

// ToolCallsListOutputItemsSenderParticipantData - Extra details the participant connected with.
type ToolCallsListOutputItemsSenderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// ToolCallsListOutputItemsSenderParticipant represents the tool calls list output items sender participant type.
type ToolCallsListOutputItemsSenderParticipant struct {
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
	Data ToolCallsListOutputItemsSenderParticipantData `json:"data"`
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

// ToolCallsListOutputItemsResponderParticipantData - Extra details the participant connected with.
type ToolCallsListOutputItemsResponderParticipantData struct {
	// Identifier - Participant identifier from the connection payload.
	Identifier string `json:"identifier"`
	// Name - Display name from the connection payload.
	Name string `json:"name"`
}

// ToolCallsListOutputItemsResponderParticipant represents the tool calls list output items responder participant type.
type ToolCallsListOutputItemsResponderParticipant struct {
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
	Data ToolCallsListOutputItemsResponderParticipantData `json:"data"`
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

// ToolCallsListOutputItemsToolInputSchema represents the tool calls list output items tool input schema type.
type ToolCallsListOutputItemsToolInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ToolCallsListOutputItemsToolOutputSchema represents the tool calls list output items tool output schema type.
type ToolCallsListOutputItemsToolOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ToolCallsListOutputItemsToolTags represents the tool calls list output items tool tags type.
type ToolCallsListOutputItemsToolTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ToolCallsListOutputItemsTool represents the tool calls list output items tool type.
type ToolCallsListOutputItemsTool struct {
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
	Instructions []string                                  `json:"instructions"`
	InputSchema  *ToolCallsListOutputItemsToolInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ToolCallsListOutputItemsToolOutputSchema `json:"output_schema,omitempty"`
	Tags         *ToolCallsListOutputItemsToolTags         `json:"tags,omitempty"`
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

// ToolCallsListOutputItemsError represents the tool calls list output items error type.
type ToolCallsListOutputItemsError struct {
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

// ToolCallsListOutputItems represents the tool calls list output items type.
type ToolCallsListOutputItems struct {
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
	ProviderRunId        *string                                       `json:"provider_run_id,omitempty"`
	SenderParticipant    *ToolCallsListOutputItemsSenderParticipant    `json:"sender_participant,omitempty"`
	ResponderParticipant *ToolCallsListOutputItemsResponderParticipant `json:"responder_participant,omitempty"`
	Tool                 ToolCallsListOutputItemsTool                  `json:"tool"`
	Error                *ToolCallsListOutputItemsError                `json:"error,omitempty"`
	// Input - Input data passed to the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Input *map[string]any `json:"input,omitempty"`
	// Output - Output data returned from the tool call. Bodies larger than 16 KB in list responses, or larger than 1 MB when retrieving by ID, are replaced with a session.message.omitted_body object.
	Output *map[string]any `json:"output,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ToolCallsListOutputPagination represents the tool calls list output pagination type.
type ToolCallsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ToolCallsListOutput represents the tool calls list output type.
type ToolCallsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
	Items      []ToolCallsListOutputItems    `json:"items"`
	Pagination ToolCallsListOutputPagination `json:"pagination"`
}

// MapToolCallsListOutputFromJSON deserializes JSON data into a ToolCallsListOutput.
func MapToolCallsListOutputFromJSON(data []byte) (*ToolCallsListOutput, error) {
	var v ToolCallsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapToolCallsListOutputToJSON serializes a ToolCallsListOutput to JSON.
func MapToolCallsListOutputToJSON(v *ToolCallsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ToolCallsListQueryCreatedAt - Only return results whose tool call creation time is in this range. Set `gt`, `lt`, or both.
type ToolCallsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ToolCallsListQueryUpdatedAt - Only return results whose tool call last update time is in this range. Set `gt`, `lt`, or both.
type ToolCallsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ToolCallsListQuery represents the tool calls list query type.
type ToolCallsListQuery struct {
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
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionTemplateId - Filter by session template ID(s)
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// AgentId - Filter by agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerProfileId - Filter by the selected consumer profile ID or IDs.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// AgentInstanceId - Filter by agent instance ID(s)
	AgentInstanceId *any `json:"agent_instance_id,omitempty"`
	// ToolId - Filter by tool ID(s)
	ToolId *any `json:"tool_id,omitempty"`
	// Status - Filter by tool call status(es)
	Status *any `json:"status,omitempty"`
	// CreatedAt - Only return results whose tool call creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ToolCallsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose tool call last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ToolCallsListQueryUpdatedAt `json:"updated_at,omitempty"`
	// Search - Search by tool call ID or tool key, or by the name of a related tool, provider, deployment, config, auth config, agent, actor, identity, or consumer
	Search *string `json:"search,omitempty"`
}

// MapToolCallsListQueryFromJSON deserializes JSON data into a ToolCallsListQuery.
func MapToolCallsListQueryFromJSON(data []byte) (*ToolCallsListQuery, error) {
	var v ToolCallsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapToolCallsListQueryToJSON serializes a ToolCallsListQuery to JSON.
func MapToolCallsListQueryToJSON(v *ToolCallsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
