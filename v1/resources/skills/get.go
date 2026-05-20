package skills

import (
	"encoding/json"
	"time"
)

// SkillsGetOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsGetOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsGetOutputHierarchyCreatorOrganizationActor represents the skills get output hierarchy creator organization actor type.
type SkillsGetOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                  `json:"image_url"`
	Teams    []SkillsGetOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyCreatorConsumer represents the skills get output hierarchy creator consumer type.
type SkillsGetOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyCreator represents the skills get output hierarchy creator type.
type SkillsGetOutputHierarchyCreator struct {
	Type              string                                            `json:"type"`
	Name              string                                            `json:"name"`
	ImageUrl          *string                                           `json:"image_url,omitempty"`
	Email             *string                                           `json:"email,omitempty"`
	OrganizationActor *SkillsGetOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsGetOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsGetOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsGetOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsGetOutputHierarchyForkCreatorOrganizationActor represents the skills get output hierarchy fork creator organization actor type.
type SkillsGetOutputHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                      `json:"image_url"`
	Teams    []SkillsGetOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyForkCreatorConsumer represents the skills get output hierarchy fork creator consumer type.
type SkillsGetOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyForkCreator represents the skills get output hierarchy fork creator type.
type SkillsGetOutputHierarchyForkCreator struct {
	Type              string                                                `json:"type"`
	Name              string                                                `json:"name"`
	ImageUrl          *string                                               `json:"image_url,omitempty"`
	Email             *string                                               `json:"email,omitempty"`
	OrganizationActor *SkillsGetOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsGetOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills get output hierarchy fork original creator organization actor type.
type SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                              `json:"image_url"`
	Teams    []SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyForkOriginalCreatorConsumer represents the skills get output hierarchy fork original creator consumer type.
type SkillsGetOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchyForkOriginalCreator represents the skills get output hierarchy fork original creator type.
type SkillsGetOutputHierarchyForkOriginalCreator struct {
	Type              string                                                        `json:"type"`
	Name              string                                                        `json:"name"`
	ImageUrl          *string                                                       `json:"image_url,omitempty"`
	Email             *string                                                       `json:"email,omitempty"`
	OrganizationActor *SkillsGetOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsGetOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsGetOutputHierarchyFork represents the skills get output hierarchy fork type.
type SkillsGetOutputHierarchyFork struct {
	Id              string                                       `json:"id"`
	ParentSkillId   string                                       `json:"parent_skill_id"`
	Creator         *SkillsGetOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsGetOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                    `json:"created_at"`
}

// SkillsGetOutputHierarchyEntity represents the skills get output hierarchy entity type.
type SkillsGetOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsGetOutputHierarchy represents the skills get output hierarchy type.
type SkillsGetOutputHierarchy struct {
	Object        string                           `json:"object"`
	Type          string                           `json:"type"`
	ParentSkillId *string                          `json:"parent_skill_id,omitempty"`
	Creator       *SkillsGetOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsGetOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsGetOutputHierarchyEntity   `json:"entity"`
}

// SkillsGetOutputIntegrationsConfiguration represents the skills get output integrations configuration type.
type SkillsGetOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsGetOutputIntegrations represents the skills get output integrations type.
type SkillsGetOutputIntegrations struct {
	Object        string                                   `json:"object"`
	Id            string                                   `json:"id"`
	Slug          string                                   `json:"slug"`
	Name          string                                   `json:"name"`
	Description   *string                                  `json:"description,omitempty"`
	Metadata      *map[string]any                          `json:"metadata,omitempty"`
	Configuration SkillsGetOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                `json:"created_at"`
	UpdatedAt     time.Time                                `json:"updated_at"`
	ArchivedAt    *time.Time                               `json:"archived_at,omitempty"`
}

// SkillsGetOutputProviders represents the skills get output providers type.
type SkillsGetOutputProviders struct {
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

// SkillsGetOutput represents the skills get output type.
type SkillsGetOutput struct {
	Object            string                        `json:"object"`
	Id                string                        `json:"id"`
	Status            string                        `json:"status"`
	Slug              string                        `json:"slug"`
	Name              string                        `json:"name"`
	Description       *string                       `json:"description,omitempty"`
	ImageUrl          string                        `json:"image_url"`
	ClientName        string                        `json:"client_name"`
	ClientDescription *string                       `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any               `json:"client_metadata,omitempty"`
	License           *string                       `json:"license,omitempty"`
	Compatibility     *string                       `json:"compatibility,omitempty"`
	Metadata          map[string]any                `json:"metadata"`
	StoreId           string                        `json:"store_id"`
	Hierarchy         SkillsGetOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsGetOutputIntegrations `json:"integrations"`
	Providers         []SkillsGetOutputProviders    `json:"providers"`
	CreatedAt         time.Time                     `json:"created_at"`
	UpdatedAt         time.Time                     `json:"updated_at"`
}

// MapSkillsGetOutputFromJSON deserializes JSON data into a SkillsGetOutput.
func MapSkillsGetOutputFromJSON(data []byte) (*SkillsGetOutput, error) {
	var v SkillsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGetOutputToJSON serializes a SkillsGetOutput to JSON.
func MapSkillsGetOutputToJSON(v *SkillsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
