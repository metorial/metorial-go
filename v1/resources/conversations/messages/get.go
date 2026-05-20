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

// ConversationsMessagesGetOutputRequestActorOrganizationActor represents the conversations messages get output request actor organization actor type.
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
	Teams    []ConversationsMessagesGetOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesGetOutputRequestActorConsumer represents the conversations messages get output request actor consumer type.
type ConversationsMessagesGetOutputRequestActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesGetOutputRequestActor represents the conversations messages get output request actor type.
type ConversationsMessagesGetOutputRequestActor struct {
	Type              string                                                       `json:"type"`
	Name              string                                                       `json:"name"`
	ImageUrl          *string                                                      `json:"image_url,omitempty"`
	Email             *string                                                      `json:"email,omitempty"`
	OrganizationActor *ConversationsMessagesGetOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsMessagesGetOutputRequestActorConsumer          `json:"consumer,omitempty"`
}

// ConversationsMessagesGetOutputRequest represents the conversations messages get output request type.
type ConversationsMessagesGetOutputRequest struct {
	Object    string                                      `json:"object"`
	Id        string                                      `json:"id"`
	Status    string                                      `json:"status"`
	Actor     *ConversationsMessagesGetOutputRequestActor `json:"actor,omitempty"`
	CreatedAt time.Time                                   `json:"created_at"`
	UpdatedAt time.Time                                   `json:"updated_at"`
}

// ConversationsMessagesGetOutput represents the conversations messages get output type.
type ConversationsMessagesGetOutput struct {
	Object             string                                `json:"object"`
	Id                 string                                `json:"id"`
	ConversationItemId string                                `json:"conversation_item_id"`
	Type               string                                `json:"type"`
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
