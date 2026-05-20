package conversations

import (
	"encoding/json"
	"time"
)

// ConversationsUpdateOutputCreatedByActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsUpdateOutputCreatedByActorOrganizationActorTeams struct {
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

// ConversationsUpdateOutputCreatedByActorOrganizationActor represents the conversations update output created by actor organization actor type.
type ConversationsUpdateOutputCreatedByActorOrganizationActor struct {
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
	Teams    []ConversationsUpdateOutputCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsUpdateOutputCreatedByActorConsumer represents the conversations update output created by actor consumer type.
type ConversationsUpdateOutputCreatedByActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsUpdateOutputCreatedByActor represents the conversations update output created by actor type.
type ConversationsUpdateOutputCreatedByActor struct {
	Type              string                                                    `json:"type"`
	Name              string                                                    `json:"name"`
	ImageUrl          *string                                                   `json:"image_url,omitempty"`
	Email             *string                                                   `json:"email,omitempty"`
	OrganizationActor *ConversationsUpdateOutputCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsUpdateOutputCreatedByActorConsumer          `json:"consumer,omitempty"`
}

// ConversationsUpdateOutputAssistantDefaultModelProvider represents the conversations update output assistant default model provider type.
type ConversationsUpdateOutputAssistantDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsUpdateOutputAssistantDefaultModel represents the conversations update output assistant default model type.
type ConversationsUpdateOutputAssistantDefaultModel struct {
	Object        string                                                 `json:"object"`
	Id            string                                                 `json:"id"`
	Slug          string                                                 `json:"slug"`
	Name          string                                                 `json:"name"`
	ContextWindow float64                                                `json:"context_window"`
	Provider      ConversationsUpdateOutputAssistantDefaultModelProvider `json:"provider"`
}

// ConversationsUpdateOutputAssistantAvailableModelsProvider represents the conversations update output assistant available models provider type.
type ConversationsUpdateOutputAssistantAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsUpdateOutputAssistantAvailableModels represents the conversations update output assistant available models type.
type ConversationsUpdateOutputAssistantAvailableModels struct {
	Object        string                                                    `json:"object"`
	Id            string                                                    `json:"id"`
	Slug          string                                                    `json:"slug"`
	Name          string                                                    `json:"name"`
	ContextWindow float64                                                   `json:"context_window"`
	Provider      ConversationsUpdateOutputAssistantAvailableModelsProvider `json:"provider"`
}

// ConversationsUpdateOutputAssistant represents the conversations update output assistant type.
type ConversationsUpdateOutputAssistant struct {
	Object          string                                              `json:"object"`
	Id              string                                              `json:"id"`
	Slug            string                                              `json:"slug"`
	Name            string                                              `json:"name"`
	OwnerType       string                                              `json:"owner_type"`
	OrganizationId  *string                                             `json:"organization_id,omitempty"`
	DefaultModel    *ConversationsUpdateOutputAssistantDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []ConversationsUpdateOutputAssistantAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                           `json:"created_at"`
	UpdatedAt       time.Time                                           `json:"updated_at"`
}

// ConversationsUpdateOutput represents the conversations update output type.
type ConversationsUpdateOutput struct {
	Object         string                                  `json:"object"`
	Id             string                                  `json:"id"`
	Title          *string                                 `json:"title,omitempty"`
	AssistantId    string                                  `json:"assistant_id"`
	InstanceId     string                                  `json:"instance_id"`
	OrganizationId string                                  `json:"organization_id"`
	CreatedByActor ConversationsUpdateOutputCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                                  `json:"root_message_id"`
	Assistant      ConversationsUpdateOutputAssistant      `json:"assistant"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
}

// MapConversationsUpdateOutputFromJSON deserializes JSON data into a ConversationsUpdateOutput.
func MapConversationsUpdateOutputFromJSON(data []byte) (*ConversationsUpdateOutput, error) {
	var v ConversationsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsUpdateOutputToJSON serializes a ConversationsUpdateOutput to JSON.
func MapConversationsUpdateOutputToJSON(v *ConversationsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsUpdateBody represents the conversations update body type.
type ConversationsUpdateBody struct {
	Title *string `json:"title,omitempty"`
}

// MapConversationsUpdateBodyFromJSON deserializes JSON data into a ConversationsUpdateBody.
func MapConversationsUpdateBodyFromJSON(data []byte) (*ConversationsUpdateBody, error) {
	var v ConversationsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsUpdateBodyToJSON serializes a ConversationsUpdateBody to JSON.
func MapConversationsUpdateBodyToJSON(v *ConversationsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
