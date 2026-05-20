package skills

import (
	"encoding/json"
	"time"
)

// SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActor represents the skills publish consumer skill output hierarchy creator organization actor type.
type SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                                   `json:"image_url"`
	Teams    []SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyCreatorConsumer represents the skills publish consumer skill output hierarchy creator consumer type.
type SkillsPublishConsumerSkillOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyCreator represents the skills publish consumer skill output hierarchy creator type.
type SkillsPublishConsumerSkillOutputHierarchyCreator struct {
	Type              string                                                             `json:"type"`
	Name              string                                                             `json:"name"`
	ImageUrl          *string                                                            `json:"image_url,omitempty"`
	Email             *string                                                            `json:"email,omitempty"`
	OrganizationActor *SkillsPublishConsumerSkillOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsPublishConsumerSkillOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActor represents the skills publish consumer skill output hierarchy fork creator organization actor type.
type SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                                       `json:"image_url"`
	Teams    []SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkCreatorConsumer represents the skills publish consumer skill output hierarchy fork creator consumer type.
type SkillsPublishConsumerSkillOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkCreator represents the skills publish consumer skill output hierarchy fork creator type.
type SkillsPublishConsumerSkillOutputHierarchyForkCreator struct {
	Type              string                                                                 `json:"type"`
	Name              string                                                                 `json:"name"`
	ImageUrl          *string                                                                `json:"image_url,omitempty"`
	Email             *string                                                                `json:"email,omitempty"`
	OrganizationActor *SkillsPublishConsumerSkillOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsPublishConsumerSkillOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills publish consumer skill output hierarchy fork original creator organization actor type.
type SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                               `json:"image_url"`
	Teams    []SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorConsumer represents the skills publish consumer skill output hierarchy fork original creator consumer type.
type SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreator represents the skills publish consumer skill output hierarchy fork original creator type.
type SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreator struct {
	Type              string                                                                         `json:"type"`
	Name              string                                                                         `json:"name"`
	ImageUrl          *string                                                                        `json:"image_url,omitempty"`
	Email             *string                                                                        `json:"email,omitempty"`
	OrganizationActor *SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsPublishConsumerSkillOutputHierarchyFork represents the skills publish consumer skill output hierarchy fork type.
type SkillsPublishConsumerSkillOutputHierarchyFork struct {
	Id              string                                                        `json:"id"`
	ParentSkillId   string                                                        `json:"parent_skill_id"`
	Creator         *SkillsPublishConsumerSkillOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsPublishConsumerSkillOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                                     `json:"created_at"`
}

// SkillsPublishConsumerSkillOutputHierarchyEntity represents the skills publish consumer skill output hierarchy entity type.
type SkillsPublishConsumerSkillOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsPublishConsumerSkillOutputHierarchy represents the skills publish consumer skill output hierarchy type.
type SkillsPublishConsumerSkillOutputHierarchy struct {
	Object        string                                            `json:"object"`
	Type          string                                            `json:"type"`
	ParentSkillId *string                                           `json:"parent_skill_id,omitempty"`
	Creator       *SkillsPublishConsumerSkillOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsPublishConsumerSkillOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsPublishConsumerSkillOutputHierarchyEntity   `json:"entity"`
}

// SkillsPublishConsumerSkillOutputIntegrationsConfiguration represents the skills publish consumer skill output integrations configuration type.
type SkillsPublishConsumerSkillOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsPublishConsumerSkillOutputIntegrations represents the skills publish consumer skill output integrations type.
type SkillsPublishConsumerSkillOutputIntegrations struct {
	Object        string                                                    `json:"object"`
	Id            string                                                    `json:"id"`
	Slug          string                                                    `json:"slug"`
	Name          string                                                    `json:"name"`
	Description   *string                                                   `json:"description,omitempty"`
	Metadata      *map[string]any                                           `json:"metadata,omitempty"`
	Configuration SkillsPublishConsumerSkillOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                 `json:"created_at"`
	UpdatedAt     time.Time                                                 `json:"updated_at"`
	ArchivedAt    *time.Time                                                `json:"archived_at,omitempty"`
}

// SkillsPublishConsumerSkillOutputProviders represents the skills publish consumer skill output providers type.
type SkillsPublishConsumerSkillOutputProviders struct {
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

// SkillsPublishConsumerSkillOutput represents the skills publish consumer skill output type.
type SkillsPublishConsumerSkillOutput struct {
	Object            string                                         `json:"object"`
	Id                string                                         `json:"id"`
	Status            string                                         `json:"status"`
	Slug              string                                         `json:"slug"`
	Name              string                                         `json:"name"`
	Description       *string                                        `json:"description,omitempty"`
	ImageUrl          string                                         `json:"image_url"`
	ClientName        string                                         `json:"client_name"`
	ClientDescription *string                                        `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                `json:"client_metadata,omitempty"`
	License           *string                                        `json:"license,omitempty"`
	Compatibility     *string                                        `json:"compatibility,omitempty"`
	Metadata          map[string]any                                 `json:"metadata"`
	StoreId           string                                         `json:"store_id"`
	Hierarchy         SkillsPublishConsumerSkillOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsPublishConsumerSkillOutputIntegrations `json:"integrations"`
	Providers         []SkillsPublishConsumerSkillOutputProviders    `json:"providers"`
	CreatedAt         time.Time                                      `json:"created_at"`
	UpdatedAt         time.Time                                      `json:"updated_at"`
}

// MapSkillsPublishConsumerSkillOutputFromJSON deserializes JSON data into a SkillsPublishConsumerSkillOutput.
func MapSkillsPublishConsumerSkillOutputFromJSON(data []byte) (*SkillsPublishConsumerSkillOutput, error) {
	var v SkillsPublishConsumerSkillOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPublishConsumerSkillOutputToJSON serializes a SkillsPublishConsumerSkillOutput to JSON.
func MapSkillsPublishConsumerSkillOutputToJSON(v *SkillsPublishConsumerSkillOutput) ([]byte, error) {
	return json.Marshal(v)
}
