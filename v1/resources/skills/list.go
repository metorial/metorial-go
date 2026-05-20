package skills

import (
	"encoding/json"
	"time"
)

// SkillsListOutputItemsHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsListOutputItemsHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsListOutputItemsHierarchyCreatorOrganizationActor represents the skills list output items hierarchy creator organization actor type.
type SkillsListOutputItemsHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                        `json:"image_url"`
	Teams    []SkillsListOutputItemsHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyCreatorConsumer represents the skills list output items hierarchy creator consumer type.
type SkillsListOutputItemsHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyCreator represents the skills list output items hierarchy creator type.
type SkillsListOutputItemsHierarchyCreator struct {
	Type              string                                                  `json:"type"`
	Name              string                                                  `json:"name"`
	ImageUrl          *string                                                 `json:"image_url,omitempty"`
	Email             *string                                                 `json:"email,omitempty"`
	OrganizationActor *SkillsListOutputItemsHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsListOutputItemsHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsListOutputItemsHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsListOutputItemsHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsListOutputItemsHierarchyForkCreatorOrganizationActor represents the skills list output items hierarchy fork creator organization actor type.
type SkillsListOutputItemsHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Teams    []SkillsListOutputItemsHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyForkCreatorConsumer represents the skills list output items hierarchy fork creator consumer type.
type SkillsListOutputItemsHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyForkCreator represents the skills list output items hierarchy fork creator type.
type SkillsListOutputItemsHierarchyForkCreator struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *SkillsListOutputItemsHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsListOutputItemsHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActor represents the skills list output items hierarchy fork original creator organization actor type.
type SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                    `json:"image_url"`
	Teams    []SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyForkOriginalCreatorConsumer represents the skills list output items hierarchy fork original creator consumer type.
type SkillsListOutputItemsHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchyForkOriginalCreator represents the skills list output items hierarchy fork original creator type.
type SkillsListOutputItemsHierarchyForkOriginalCreator struct {
	Type              string                                                              `json:"type"`
	Name              string                                                              `json:"name"`
	ImageUrl          *string                                                             `json:"image_url,omitempty"`
	Email             *string                                                             `json:"email,omitempty"`
	OrganizationActor *SkillsListOutputItemsHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsListOutputItemsHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsListOutputItemsHierarchyFork represents the skills list output items hierarchy fork type.
type SkillsListOutputItemsHierarchyFork struct {
	Id              string                                             `json:"id"`
	ParentSkillId   string                                             `json:"parent_skill_id"`
	Creator         *SkillsListOutputItemsHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsListOutputItemsHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                          `json:"created_at"`
}

// SkillsListOutputItemsHierarchyEntity represents the skills list output items hierarchy entity type.
type SkillsListOutputItemsHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsListOutputItemsHierarchy represents the skills list output items hierarchy type.
type SkillsListOutputItemsHierarchy struct {
	Object        string                                 `json:"object"`
	Type          string                                 `json:"type"`
	ParentSkillId *string                                `json:"parent_skill_id,omitempty"`
	Creator       *SkillsListOutputItemsHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsListOutputItemsHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsListOutputItemsHierarchyEntity   `json:"entity"`
}

// SkillsListOutputItemsIntegrationsConfiguration represents the skills list output items integrations configuration type.
type SkillsListOutputItemsIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsListOutputItemsIntegrations represents the skills list output items integrations type.
type SkillsListOutputItemsIntegrations struct {
	Object        string                                         `json:"object"`
	Id            string                                         `json:"id"`
	Slug          string                                         `json:"slug"`
	Name          string                                         `json:"name"`
	Description   *string                                        `json:"description,omitempty"`
	Metadata      *map[string]any                                `json:"metadata,omitempty"`
	Configuration SkillsListOutputItemsIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                      `json:"created_at"`
	UpdatedAt     time.Time                                      `json:"updated_at"`
	ArchivedAt    *time.Time                                     `json:"archived_at,omitempty"`
}

// SkillsListOutputItemsProviders represents the skills list output items providers type.
type SkillsListOutputItemsProviders struct {
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

// SkillsListOutputItems represents the skills list output items type.
type SkillsListOutputItems struct {
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	Status            string                              `json:"status"`
	Slug              string                              `json:"slug"`
	Name              string                              `json:"name"`
	Description       *string                             `json:"description,omitempty"`
	ImageUrl          string                              `json:"image_url"`
	ClientName        string                              `json:"client_name"`
	ClientDescription *string                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                     `json:"client_metadata,omitempty"`
	License           *string                             `json:"license,omitempty"`
	Compatibility     *string                             `json:"compatibility,omitempty"`
	Metadata          map[string]any                      `json:"metadata"`
	StoreId           string                              `json:"store_id"`
	Hierarchy         SkillsListOutputItemsHierarchy      `json:"hierarchy"`
	Integrations      []SkillsListOutputItemsIntegrations `json:"integrations"`
	Providers         []SkillsListOutputItemsProviders    `json:"providers"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

// SkillsListOutputPagination represents the skills list output pagination type.
type SkillsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsListOutput represents the skills list output type.
type SkillsListOutput struct {
	Items      []SkillsListOutputItems    `json:"items"`
	Pagination SkillsListOutputPagination `json:"pagination"`
}

// MapSkillsListOutputFromJSON deserializes JSON data into a SkillsListOutput.
func MapSkillsListOutputFromJSON(data []byte) (*SkillsListOutput, error) {
	var v SkillsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsListOutputToJSON serializes a SkillsListOutput to JSON.
func MapSkillsListOutputToJSON(v *SkillsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsListQueryCreatedAt - Filter skill creation time by date range
type SkillsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsListQueryUpdatedAt - Filter skill last update time by date range
type SkillsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsListQuery represents the skills list query type.
type SkillsListQuery struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Search        *string  `json:"search,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Id            *any     `json:"id,omitempty"`
	SkillGroupId  *any     `json:"skill_group_id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	// CreatedAt - Filter skill creation time by date range
	CreatedAt *SkillsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill last update time by date range
	UpdatedAt *SkillsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsListQueryFromJSON deserializes JSON data into a SkillsListQuery.
func MapSkillsListQueryFromJSON(data []byte) (*SkillsListQuery, error) {
	var v SkillsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsListQueryToJSON serializes a SkillsListQuery to JSON.
func MapSkillsListQueryToJSON(v *SkillsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
