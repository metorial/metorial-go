package messages

import (
	"encoding/json"
	"time"
)

// ConversationsMessagesCancelOutputModelProvider represents the conversations messages cancel output model provider type.
type ConversationsMessagesCancelOutputModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsMessagesCancelOutputModel represents the conversations messages cancel output model type.
type ConversationsMessagesCancelOutputModel struct {
	Object        string                                         `json:"object"`
	Id            string                                         `json:"id"`
	Slug          string                                         `json:"slug"`
	Name          string                                         `json:"name"`
	ContextWindow float64                                        `json:"context_window"`
	Provider      ConversationsMessagesCancelOutputModelProvider `json:"provider"`
}

// ConversationsMessagesCancelOutputRequestActorOrganizationActorMember represents the conversations messages cancel output request actor organization actor member type.
type ConversationsMessagesCancelOutputRequestActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsMessagesCancelOutputRequestActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsMessagesCancelOutputRequestActorOrganizationActorTeams struct {
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

// ConversationsMessagesCancelOutputRequestActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsMessagesCancelOutputRequestActorOrganizationActor struct {
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
	Member   *ConversationsMessagesCancelOutputRequestActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsMessagesCancelOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesCancelOutputRequestActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsMessagesCancelOutputRequestActorConsumer struct {
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

// ConversationsMessagesCancelOutputRequestActor - Actor behind a participant.
type ConversationsMessagesCancelOutputRequestActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsMessagesCancelOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsMessagesCancelOutputRequestActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsMessagesCancelOutputRequestContext represents the conversations messages cancel output request context type.
type ConversationsMessagesCancelOutputRequestContext struct {
	// Mode - The mode the assistant ran in after permission checks
	Mode string `json:"mode"`
	// RequestedMode - The mode requested by the client
	RequestedMode string `json:"requested_mode"`
	// CanWrite - Whether the assistant could change the underlying resource
	CanWrite bool `json:"can_write"`
}

// ConversationsMessagesCancelOutputRequest represents the conversations messages cancel output request type.
type ConversationsMessagesCancelOutputRequest struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Actor - Actor behind a participant.
	Actor     *ConversationsMessagesCancelOutputRequestActor   `json:"actor,omitempty"`
	Context   *ConversationsMessagesCancelOutputRequestContext `json:"context,omitempty"`
	CreatedAt time.Time                                        `json:"created_at"`
	UpdatedAt time.Time                                        `json:"updated_at"`
}

// ConversationsMessagesCancelOutput represents the conversations messages cancel output type.
type ConversationsMessagesCancelOutput struct {
	Object             string                                   `json:"object"`
	Id                 string                                   `json:"id"`
	ConversationItemId string                                   `json:"conversation_item_id"`
	Type               string                                   `json:"type"`
	Status             string                                   `json:"status"`
	AssistantId        *string                                  `json:"assistant_id,omitempty"`
	ParentMessageId    *string                                  `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesCancelOutputModel  `json:"model,omitempty"`
	Request            ConversationsMessagesCancelOutputRequest `json:"request"`
	Items              []map[string]any                         `json:"items"`
	CreatedAt          time.Time                                `json:"created_at"`
}

// MapConversationsMessagesCancelOutputFromJSON deserializes JSON data into a ConversationsMessagesCancelOutput.
func MapConversationsMessagesCancelOutputFromJSON(data []byte) (*ConversationsMessagesCancelOutput, error) {
	var v ConversationsMessagesCancelOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesCancelOutputToJSON serializes a ConversationsMessagesCancelOutput to JSON.
func MapConversationsMessagesCancelOutputToJSON(v *ConversationsMessagesCancelOutput) ([]byte, error) {
	return json.Marshal(v)
}
