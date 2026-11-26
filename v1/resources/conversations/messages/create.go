package messages

import (
	"encoding/json"
	"time"
)

// ConversationsMessagesCreateOutputModelProvider represents the conversations messages create output model provider type.
type ConversationsMessagesCreateOutputModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsMessagesCreateOutputModel represents the conversations messages create output model type.
type ConversationsMessagesCreateOutputModel struct {
	Object        string                                         `json:"object"`
	Id            string                                         `json:"id"`
	Slug          string                                         `json:"slug"`
	Name          string                                         `json:"name"`
	ContextWindow float64                                        `json:"context_window"`
	Provider      ConversationsMessagesCreateOutputModelProvider `json:"provider"`
}

// ConversationsMessagesCreateOutputRequestActorOrganizationActorMember represents the conversations messages create output request actor organization actor member type.
type ConversationsMessagesCreateOutputRequestActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsMessagesCreateOutputRequestActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsMessagesCreateOutputRequestActorOrganizationActorTeams struct {
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

// ConversationsMessagesCreateOutputRequestActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsMessagesCreateOutputRequestActorOrganizationActor struct {
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
	ImageUrl string                                                                `json:"image_url"`
	Member   *ConversationsMessagesCreateOutputRequestActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsMessagesCreateOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesCreateOutputRequestActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsMessagesCreateOutputRequestActorConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesCreateOutputRequestActor - Actor behind a participant.
type ConversationsMessagesCreateOutputRequestActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsMessagesCreateOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsMessagesCreateOutputRequestActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsMessagesCreateOutputRequestContext represents the conversations messages create output request context type.
type ConversationsMessagesCreateOutputRequestContext struct {
	// Mode - The mode the assistant ran in after permission checks
	Mode string `json:"mode"`
	// RequestedMode - The mode requested by the client
	RequestedMode string `json:"requested_mode"`
	// CanWrite - Whether the assistant could change the underlying resource
	CanWrite bool `json:"can_write"`
}

// ConversationsMessagesCreateOutputRequest represents the conversations messages create output request type.
type ConversationsMessagesCreateOutputRequest struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Actor - Actor behind a participant.
	Actor     *ConversationsMessagesCreateOutputRequestActor   `json:"actor,omitempty"`
	Context   *ConversationsMessagesCreateOutputRequestContext `json:"context,omitempty"`
	CreatedAt time.Time                                        `json:"created_at"`
	UpdatedAt time.Time                                        `json:"updated_at"`
}

// ConversationsMessagesCreateOutput represents the conversations messages create output type.
type ConversationsMessagesCreateOutput struct {
	Object             string                                   `json:"object"`
	Id                 string                                   `json:"id"`
	ConversationItemId string                                   `json:"conversation_item_id"`
	Type               string                                   `json:"type"`
	Status             string                                   `json:"status"`
	AssistantId        *string                                  `json:"assistant_id,omitempty"`
	ParentMessageId    *string                                  `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesCreateOutputModel  `json:"model,omitempty"`
	Request            ConversationsMessagesCreateOutputRequest `json:"request"`
	Items              []map[string]any                         `json:"items"`
	CreatedAt          time.Time                                `json:"created_at"`
}

// MapConversationsMessagesCreateOutputFromJSON deserializes JSON data into a ConversationsMessagesCreateOutput.
func MapConversationsMessagesCreateOutputFromJSON(data []byte) (*ConversationsMessagesCreateOutput, error) {
	var v ConversationsMessagesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesCreateOutputToJSON serializes a ConversationsMessagesCreateOutput to JSON.
func MapConversationsMessagesCreateOutputToJSON(v *ConversationsMessagesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsMessagesCreateBodyMessage represents the conversations messages create body message type.
type ConversationsMessagesCreateBodyMessage struct {
	Parts []any `json:"parts"`
}

// ConversationsMessagesCreateBody represents the conversations messages create body type.
type ConversationsMessagesCreateBody struct {
	Message         ConversationsMessagesCreateBodyMessage `json:"message"`
	ParentMessageId *string                                `json:"parent_message_id,omitempty"`
	ModelId         *string                                `json:"model_id,omitempty"`
	// Context - Assistant-specific request context, such as the editing mode, the frozen selection, or the target block. Validated by the assistant implementation.
	Context *map[string]any `json:"context,omitempty"`
}

// MapConversationsMessagesCreateBodyFromJSON deserializes JSON data into a ConversationsMessagesCreateBody.
func MapConversationsMessagesCreateBodyFromJSON(data []byte) (*ConversationsMessagesCreateBody, error) {
	var v ConversationsMessagesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesCreateBodyToJSON serializes a ConversationsMessagesCreateBody to JSON.
func MapConversationsMessagesCreateBodyToJSON(v *ConversationsMessagesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
