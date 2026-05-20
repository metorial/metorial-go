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

// ConversationsMessagesCreateOutputRequestActorOrganizationActor represents the conversations messages create output request actor organization actor type.
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
	Teams    []ConversationsMessagesCreateOutputRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesCreateOutputRequestActorConsumer represents the conversations messages create output request actor consumer type.
type ConversationsMessagesCreateOutputRequestActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesCreateOutputRequestActor represents the conversations messages create output request actor type.
type ConversationsMessagesCreateOutputRequestActor struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *ConversationsMessagesCreateOutputRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsMessagesCreateOutputRequestActorConsumer          `json:"consumer,omitempty"`
}

// ConversationsMessagesCreateOutputRequest represents the conversations messages create output request type.
type ConversationsMessagesCreateOutputRequest struct {
	Object    string                                         `json:"object"`
	Id        string                                         `json:"id"`
	Status    string                                         `json:"status"`
	Actor     *ConversationsMessagesCreateOutputRequestActor `json:"actor,omitempty"`
	CreatedAt time.Time                                      `json:"created_at"`
	UpdatedAt time.Time                                      `json:"updated_at"`
}

// ConversationsMessagesCreateOutput represents the conversations messages create output type.
type ConversationsMessagesCreateOutput struct {
	Object             string                                   `json:"object"`
	Id                 string                                   `json:"id"`
	ConversationItemId string                                   `json:"conversation_item_id"`
	Type               string                                   `json:"type"`
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
