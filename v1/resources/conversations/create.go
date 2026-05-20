package conversations

import (
	"encoding/json"
	"time"
)

// ConversationsCreateOutputCreatedByActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsCreateOutputCreatedByActorOrganizationActorTeams struct {
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

// ConversationsCreateOutputCreatedByActorOrganizationActor represents the conversations create output created by actor organization actor type.
type ConversationsCreateOutputCreatedByActorOrganizationActor struct {
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
	ImageUrl string                                                          `json:"image_url"`
	Teams    []ConversationsCreateOutputCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsCreateOutputCreatedByActorConsumer represents the conversations create output created by actor consumer type.
type ConversationsCreateOutputCreatedByActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsCreateOutputCreatedByActor represents the conversations create output created by actor type.
type ConversationsCreateOutputCreatedByActor struct {
	Type              string                                                    `json:"type"`
	Name              string                                                    `json:"name"`
	ImageUrl          *string                                                   `json:"image_url,omitempty"`
	Email             *string                                                   `json:"email,omitempty"`
	OrganizationActor *ConversationsCreateOutputCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsCreateOutputCreatedByActorConsumer          `json:"consumer,omitempty"`
}

// ConversationsCreateOutputAssistantDefaultModelProvider represents the conversations create output assistant default model provider type.
type ConversationsCreateOutputAssistantDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsCreateOutputAssistantDefaultModel represents the conversations create output assistant default model type.
type ConversationsCreateOutputAssistantDefaultModel struct {
	Object        string                                                 `json:"object"`
	Id            string                                                 `json:"id"`
	Slug          string                                                 `json:"slug"`
	Name          string                                                 `json:"name"`
	ContextWindow float64                                                `json:"context_window"`
	Provider      ConversationsCreateOutputAssistantDefaultModelProvider `json:"provider"`
}

// ConversationsCreateOutputAssistantAvailableModelsProvider represents the conversations create output assistant available models provider type.
type ConversationsCreateOutputAssistantAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsCreateOutputAssistantAvailableModels represents the conversations create output assistant available models type.
type ConversationsCreateOutputAssistantAvailableModels struct {
	Object        string                                                    `json:"object"`
	Id            string                                                    `json:"id"`
	Slug          string                                                    `json:"slug"`
	Name          string                                                    `json:"name"`
	ContextWindow float64                                                   `json:"context_window"`
	Provider      ConversationsCreateOutputAssistantAvailableModelsProvider `json:"provider"`
}

// ConversationsCreateOutputAssistant represents the conversations create output assistant type.
type ConversationsCreateOutputAssistant struct {
	Object          string                                              `json:"object"`
	Id              string                                              `json:"id"`
	Slug            string                                              `json:"slug"`
	Name            string                                              `json:"name"`
	OwnerType       string                                              `json:"owner_type"`
	OrganizationId  *string                                             `json:"organization_id,omitempty"`
	DefaultModel    *ConversationsCreateOutputAssistantDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []ConversationsCreateOutputAssistantAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                           `json:"created_at"`
	UpdatedAt       time.Time                                           `json:"updated_at"`
}

// ConversationsCreateOutput represents the conversations create output type.
type ConversationsCreateOutput struct {
	Object         string                                  `json:"object"`
	Id             string                                  `json:"id"`
	Title          *string                                 `json:"title,omitempty"`
	AssistantId    string                                  `json:"assistant_id"`
	InstanceId     string                                  `json:"instance_id"`
	OrganizationId string                                  `json:"organization_id"`
	CreatedByActor ConversationsCreateOutputCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                                  `json:"root_message_id"`
	Assistant      ConversationsCreateOutputAssistant      `json:"assistant"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
}

// MapConversationsCreateOutputFromJSON deserializes JSON data into a ConversationsCreateOutput.
func MapConversationsCreateOutputFromJSON(data []byte) (*ConversationsCreateOutput, error) {
	var v ConversationsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsCreateOutputToJSON serializes a ConversationsCreateOutput to JSON.
func MapConversationsCreateOutputToJSON(v *ConversationsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsCreateBody represents the conversations create body type.
type ConversationsCreateBody struct {
	AssistantId string  `json:"assistant_id"`
	Title       *string `json:"title,omitempty"`
}

// MapConversationsCreateBodyFromJSON deserializes JSON data into a ConversationsCreateBody.
func MapConversationsCreateBodyFromJSON(data []byte) (*ConversationsCreateBody, error) {
	var v ConversationsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsCreateBodyToJSON serializes a ConversationsCreateBody to JSON.
func MapConversationsCreateBodyToJSON(v *ConversationsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
