package messages

import (
	"encoding/json"
	"time"
)

// ConversationsMessagesHandoffResponsesOutputModelProvider represents the conversations messages handoff responses output model provider type.
type ConversationsMessagesHandoffResponsesOutputModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsMessagesHandoffResponsesOutputModel represents the conversations messages handoff responses output model type.
type ConversationsMessagesHandoffResponsesOutputModel struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	ContextWindow float64                                                  `json:"context_window"`
	Provider      ConversationsMessagesHandoffResponsesOutputModelProvider `json:"provider"`
}

// ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorMember represents the conversations messages handoff responses output request actor organization actor member type.
type ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorTeams struct {
	// Id - The team ID
	Id string `json:"id"`
	// Name - The team name
	Name string `json:"name"`
	// Slug - The team slug
	Slug string `json:"slug"`
	// AssignmentId - The team assignment ID
	AssignmentId string `json:"assignment_id"`
	// CreatedAt - The team assignment creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The team assignment last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActor represents the conversations messages handoff responses output request actor organization actor type.
type ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActor struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Type - The organization member's type
	Type string `json:"type"`
	// OrganizationId - The organization member's organization ID
	OrganizationId string `json:"organization_id"`
	// Name - The organization member's name
	Name string `json:"name"`
	// Email - The organization member's email
	Email *string `json:"email,omitempty"`
	// ImageUrl - The organization member's image URL
	ImageUrl string                                                                          `json:"image_url"`
	Member   *ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesHandoffResponsesOutputRequestActorConsumer represents the conversations messages handoff responses output request actor consumer type.
type ConversationsMessagesHandoffResponsesOutputRequestActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesHandoffResponsesOutputRequestActor represents the conversations messages handoff responses output request actor type.
type ConversationsMessagesHandoffResponsesOutputRequestActor struct {
	Type              string                                                                    `json:"type"`
	Name              string                                                                    `json:"name"`
	ImageUrl          *string                                                                   `json:"image_url,omitempty"`
	Email             *string                                                                   `json:"email,omitempty"`
	OrganizationActor *ConversationsMessagesHandoffResponsesOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsMessagesHandoffResponsesOutputRequestActorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                           `json:"consumer_profile,omitempty"`
}

// ConversationsMessagesHandoffResponsesOutputRequest represents the conversations messages handoff responses output request type.
type ConversationsMessagesHandoffResponsesOutputRequest struct {
	Object    string                                                   `json:"object"`
	Id        string                                                   `json:"id"`
	Status    string                                                   `json:"status"`
	Actor     *ConversationsMessagesHandoffResponsesOutputRequestActor `json:"actor,omitempty"`
	CreatedAt time.Time                                                `json:"created_at"`
	UpdatedAt time.Time                                                `json:"updated_at"`
}

// ConversationsMessagesHandoffResponsesOutput represents the conversations messages handoff responses output type.
type ConversationsMessagesHandoffResponsesOutput struct {
	Object             string                                             `json:"object"`
	Id                 string                                             `json:"id"`
	ConversationItemId string                                             `json:"conversation_item_id"`
	Type               string                                             `json:"type"`
	Status             string                                             `json:"status"`
	AssistantId        *string                                            `json:"assistant_id,omitempty"`
	ParentMessageId    *string                                            `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesHandoffResponsesOutputModel  `json:"model,omitempty"`
	Request            ConversationsMessagesHandoffResponsesOutputRequest `json:"request"`
	Items              []map[string]any                                   `json:"items"`
	CreatedAt          time.Time                                          `json:"created_at"`
}

// MapConversationsMessagesHandoffResponsesOutputFromJSON deserializes JSON data into a ConversationsMessagesHandoffResponsesOutput.
func MapConversationsMessagesHandoffResponsesOutputFromJSON(data []byte) (*ConversationsMessagesHandoffResponsesOutput, error) {
	var v ConversationsMessagesHandoffResponsesOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesHandoffResponsesOutputToJSON serializes a ConversationsMessagesHandoffResponsesOutput to JSON.
func MapConversationsMessagesHandoffResponsesOutputToJSON(v *ConversationsMessagesHandoffResponsesOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsMessagesHandoffResponsesBodyResponses represents the conversations messages handoff responses body responses type.
type ConversationsMessagesHandoffResponsesBodyResponses struct {
	ToolCallId string `json:"tool_call_id"`
	Output     any    `json:"output"`
}

// ConversationsMessagesHandoffResponsesBody represents the conversations messages handoff responses body type.
type ConversationsMessagesHandoffResponsesBody struct {
	Responses []ConversationsMessagesHandoffResponsesBodyResponses `json:"responses"`
}

// MapConversationsMessagesHandoffResponsesBodyFromJSON deserializes JSON data into a ConversationsMessagesHandoffResponsesBody.
func MapConversationsMessagesHandoffResponsesBodyFromJSON(data []byte) (*ConversationsMessagesHandoffResponsesBody, error) {
	var v ConversationsMessagesHandoffResponsesBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesHandoffResponsesBodyToJSON serializes a ConversationsMessagesHandoffResponsesBody to JSON.
func MapConversationsMessagesHandoffResponsesBodyToJSON(v *ConversationsMessagesHandoffResponsesBody) ([]byte, error) {
	return json.Marshal(v)
}
