package messages

import (
	"encoding/json"
	"time"
)

// ConversationsMessagesGetOutputModelProvider represents the conversations messages get output model provider type.
type ConversationsMessagesGetOutputModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsMessagesGetOutputModel represents the conversations messages get output model type.
type ConversationsMessagesGetOutputModel struct {
	Object        string                                      `json:"object"`
	Id            string                                      `json:"id"`
	Slug          string                                      `json:"slug"`
	Name          string                                      `json:"name"`
	ContextWindow float64                                     `json:"context_window"`
	Provider      ConversationsMessagesGetOutputModelProvider `json:"provider"`
}

// ConversationsMessagesGetOutputRequestActorOrganizationActorMember represents the conversations messages get output request actor organization actor member type.
type ConversationsMessagesGetOutputRequestActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsMessagesGetOutputRequestActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsMessagesGetOutputRequestActorOrganizationActorTeams struct {
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

// ConversationsMessagesGetOutputRequestActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsMessagesGetOutputRequestActorOrganizationActor struct {
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
	ImageUrl string                                                             `json:"image_url"`
	Member   *ConversationsMessagesGetOutputRequestActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsMessagesGetOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesGetOutputRequestActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsMessagesGetOutputRequestActorConsumer struct {
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

// ConversationsMessagesGetOutputRequestActor - Actor behind a participant.
type ConversationsMessagesGetOutputRequestActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsMessagesGetOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsMessagesGetOutputRequestActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsMessagesGetOutputRequestContext represents the conversations messages get output request context type.
type ConversationsMessagesGetOutputRequestContext struct {
	// Mode - The mode the assistant ran in after permission checks
	Mode string `json:"mode"`
	// RequestedMode - The mode requested by the client
	RequestedMode string `json:"requested_mode"`
	// CanWrite - Whether the assistant could change the underlying resource
	CanWrite bool `json:"can_write"`
}

// ConversationsMessagesGetOutputRequest represents the conversations messages get output request type.
type ConversationsMessagesGetOutputRequest struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Actor - Actor behind a participant.
	Actor     *ConversationsMessagesGetOutputRequestActor   `json:"actor,omitempty"`
	Context   *ConversationsMessagesGetOutputRequestContext `json:"context,omitempty"`
	CreatedAt time.Time                                     `json:"created_at"`
	UpdatedAt time.Time                                     `json:"updated_at"`
}

// ConversationsMessagesGetOutput represents the conversations messages get output type.
type ConversationsMessagesGetOutput struct {
	Object             string                                `json:"object"`
	Id                 string                                `json:"id"`
	ConversationItemId string                                `json:"conversation_item_id"`
	Type               string                                `json:"type"`
	Status             string                                `json:"status"`
	AssistantId        *string                               `json:"assistant_id,omitempty"`
	ParentMessageId    *string                               `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesGetOutputModel  `json:"model,omitempty"`
	Request            ConversationsMessagesGetOutputRequest `json:"request"`
	Items              []map[string]any                      `json:"items"`
	CreatedAt          time.Time                             `json:"created_at"`
}

// MapConversationsMessagesGetOutputFromJSON deserializes JSON data into a ConversationsMessagesGetOutput.
func MapConversationsMessagesGetOutputFromJSON(data []byte) (*ConversationsMessagesGetOutput, error) {
	var v ConversationsMessagesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesGetOutputToJSON serializes a ConversationsMessagesGetOutput to JSON.
func MapConversationsMessagesGetOutputToJSON(v *ConversationsMessagesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
