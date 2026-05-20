package skills

import (
	"encoding/json"
	"time"
)

// SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyCreatorOrganizationActor represents the skills delete output hierarchy creator organization actor type.
type SkillsDeleteOutputHierarchyCreatorOrganizationActor struct {
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
	Teams    []SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyCreatorConsumer represents the skills delete output hierarchy creator consumer type.
type SkillsDeleteOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyCreator represents the skills delete output hierarchy creator type.
type SkillsDeleteOutputHierarchyCreator struct {
	Type              string                                               `json:"type"`
	Name              string                                               `json:"name"`
	ImageUrl          *string                                              `json:"image_url,omitempty"`
	Email             *string                                              `json:"email,omitempty"`
	OrganizationActor *SkillsDeleteOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDeleteOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyForkCreatorOrganizationActor represents the skills delete output hierarchy fork creator organization actor type.
type SkillsDeleteOutputHierarchyForkCreatorOrganizationActor struct {
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
	Teams    []SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkCreatorConsumer represents the skills delete output hierarchy fork creator consumer type.
type SkillsDeleteOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkCreator represents the skills delete output hierarchy fork creator type.
type SkillsDeleteOutputHierarchyForkCreator struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *SkillsDeleteOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDeleteOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills delete output hierarchy fork original creator organization actor type.
type SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	Teams    []SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer represents the skills delete output hierarchy fork original creator consumer type.
type SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreator represents the skills delete output hierarchy fork original creator type.
type SkillsDeleteOutputHierarchyForkOriginalCreator struct {
	Type              string                                                           `json:"type"`
	Name              string                                                           `json:"name"`
	ImageUrl          *string                                                          `json:"image_url,omitempty"`
	Email             *string                                                          `json:"email,omitempty"`
	OrganizationActor *SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDeleteOutputHierarchyFork represents the skills delete output hierarchy fork type.
type SkillsDeleteOutputHierarchyFork struct {
	Id              string                                          `json:"id"`
	ParentSkillId   string                                          `json:"parent_skill_id"`
	Creator         *SkillsDeleteOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsDeleteOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                       `json:"created_at"`
}

// SkillsDeleteOutputHierarchyEntity represents the skills delete output hierarchy entity type.
type SkillsDeleteOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchy represents the skills delete output hierarchy type.
type SkillsDeleteOutputHierarchy struct {
	Object        string                              `json:"object"`
	Type          string                              `json:"type"`
	ParentSkillId *string                             `json:"parent_skill_id,omitempty"`
	Creator       *SkillsDeleteOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsDeleteOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsDeleteOutputHierarchyEntity   `json:"entity"`
}

// SkillsDeleteOutputIntegrationsConfiguration represents the skills delete output integrations configuration type.
type SkillsDeleteOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsDeleteOutputIntegrations represents the skills delete output integrations type.
type SkillsDeleteOutputIntegrations struct {
	Object        string                                      `json:"object"`
	Id            string                                      `json:"id"`
	Slug          string                                      `json:"slug"`
	Name          string                                      `json:"name"`
	Description   *string                                     `json:"description,omitempty"`
	Metadata      *map[string]any                             `json:"metadata,omitempty"`
	Configuration SkillsDeleteOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                   `json:"created_at"`
	UpdatedAt     time.Time                                   `json:"updated_at"`
	ArchivedAt    *time.Time                                  `json:"archived_at,omitempty"`
}

// SkillsDeleteOutputProviders represents the skills delete output providers type.
type SkillsDeleteOutputProviders struct {
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

// SkillsDeleteOutput represents the skills delete output type.
type SkillsDeleteOutput struct {
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
	Hierarchy         SkillsDeleteOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsDeleteOutputIntegrations `json:"integrations"`
	Providers         []SkillsDeleteOutputProviders    `json:"providers"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}

// MapSkillsDeleteOutputFromJSON deserializes JSON data into a SkillsDeleteOutput.
func MapSkillsDeleteOutputFromJSON(data []byte) (*SkillsDeleteOutput, error) {
	var v SkillsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDeleteOutputToJSON serializes a SkillsDeleteOutput to JSON.
func MapSkillsDeleteOutputToJSON(v *SkillsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
