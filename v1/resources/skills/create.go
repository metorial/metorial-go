package skills

import (
	"encoding/json"
	"time"
)

// SkillsCreateOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsCreateOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsCreateOutputHierarchyCreatorOrganizationActor represents the skills create output hierarchy creator organization actor type.
type SkillsCreateOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                     `json:"image_url"`
	Teams    []SkillsCreateOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyCreatorConsumer represents the skills create output hierarchy creator consumer type.
type SkillsCreateOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyCreator represents the skills create output hierarchy creator type.
type SkillsCreateOutputHierarchyCreator struct {
	Type              string                                               `json:"type"`
	Name              string                                               `json:"name"`
	ImageUrl          *string                                              `json:"image_url,omitempty"`
	Email             *string                                              `json:"email,omitempty"`
	OrganizationActor *SkillsCreateOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsCreateOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsCreateOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsCreateOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsCreateOutputHierarchyForkCreatorOrganizationActor represents the skills create output hierarchy fork creator organization actor type.
type SkillsCreateOutputHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                         `json:"image_url"`
	Teams    []SkillsCreateOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyForkCreatorConsumer represents the skills create output hierarchy fork creator consumer type.
type SkillsCreateOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyForkCreator represents the skills create output hierarchy fork creator type.
type SkillsCreateOutputHierarchyForkCreator struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *SkillsCreateOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsCreateOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills create output hierarchy fork original creator organization actor type.
type SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                 `json:"image_url"`
	Teams    []SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyForkOriginalCreatorConsumer represents the skills create output hierarchy fork original creator consumer type.
type SkillsCreateOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchyForkOriginalCreator represents the skills create output hierarchy fork original creator type.
type SkillsCreateOutputHierarchyForkOriginalCreator struct {
	Type              string                                                           `json:"type"`
	Name              string                                                           `json:"name"`
	ImageUrl          *string                                                          `json:"image_url,omitempty"`
	Email             *string                                                          `json:"email,omitempty"`
	OrganizationActor *SkillsCreateOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsCreateOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsCreateOutputHierarchyFork represents the skills create output hierarchy fork type.
type SkillsCreateOutputHierarchyFork struct {
	Id              string                                          `json:"id"`
	ParentSkillId   string                                          `json:"parent_skill_id"`
	Creator         *SkillsCreateOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsCreateOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                       `json:"created_at"`
}

// SkillsCreateOutputHierarchyEntity represents the skills create output hierarchy entity type.
type SkillsCreateOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsCreateOutputHierarchy represents the skills create output hierarchy type.
type SkillsCreateOutputHierarchy struct {
	Object        string                              `json:"object"`
	Type          string                              `json:"type"`
	ParentSkillId *string                             `json:"parent_skill_id,omitempty"`
	Creator       *SkillsCreateOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsCreateOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsCreateOutputHierarchyEntity   `json:"entity"`
}

// SkillsCreateOutputIntegrationsConfiguration represents the skills create output integrations configuration type.
type SkillsCreateOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsCreateOutputIntegrations represents the skills create output integrations type.
type SkillsCreateOutputIntegrations struct {
	Object        string                                      `json:"object"`
	Id            string                                      `json:"id"`
	Slug          string                                      `json:"slug"`
	Name          string                                      `json:"name"`
	Description   *string                                     `json:"description,omitempty"`
	Metadata      *map[string]any                             `json:"metadata,omitempty"`
	Configuration SkillsCreateOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                   `json:"created_at"`
	UpdatedAt     time.Time                                   `json:"updated_at"`
	ArchivedAt    *time.Time                                  `json:"archived_at,omitempty"`
}

// SkillsCreateOutputProviders represents the skills create output providers type.
type SkillsCreateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsCreateOutput represents the skills create output type.
type SkillsCreateOutput struct {
	Object            string                           `json:"object"`
	Id                string                           `json:"id"`
	Status            string                           `json:"status"`
	Slug              string                           `json:"slug"`
	Name              string                           `json:"name"`
	Description       *string                          `json:"description,omitempty"`
	ImageUrl          string                           `json:"image_url"`
	ClientName        string                           `json:"client_name"`
	ClientDescription *string                          `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                  `json:"client_metadata,omitempty"`
	License           *string                          `json:"license,omitempty"`
	Compatibility     *string                          `json:"compatibility,omitempty"`
	Metadata          map[string]any                   `json:"metadata"`
	StoreId           string                           `json:"store_id"`
	Hierarchy         SkillsCreateOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsCreateOutputIntegrations `json:"integrations"`
	Providers         []SkillsCreateOutputProviders    `json:"providers"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}

// MapSkillsCreateOutputFromJSON deserializes JSON data into a SkillsCreateOutput.
func MapSkillsCreateOutputFromJSON(data []byte) (*SkillsCreateOutput, error) {
	var v SkillsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsCreateOutputToJSON serializes a SkillsCreateOutput to JSON.
func MapSkillsCreateOutputToJSON(v *SkillsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsCreateBody represents the skills create body type.
type SkillsCreateBody struct {
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	ImageFileId       *string         `json:"image_file_id,omitempty"`
	TemplateId        *string         `json:"template_id,omitempty"`
}

// MapSkillsCreateBodyFromJSON deserializes JSON data into a SkillsCreateBody.
func MapSkillsCreateBodyFromJSON(data []byte) (*SkillsCreateBody, error) {
	var v SkillsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsCreateBodyToJSON serializes a SkillsCreateBody to JSON.
func MapSkillsCreateBodyToJSON(v *SkillsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
