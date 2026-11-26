package skills

import (
	"encoding/json"
	"time"
)

// SkillsDuplicateOutputImage - Image shown for the skill.
type SkillsDuplicateOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsDuplicateOutputHierarchyCreatorOrganizationActorMember represents the skills duplicate output hierarchy creator organization actor member type.
type SkillsDuplicateOutputHierarchyCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDuplicateOutputHierarchyCreatorOrganizationActor struct {
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
	Member   *SkillsDuplicateOutputHierarchyCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDuplicateOutputHierarchyCreatorConsumer struct {
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

// SkillsDuplicateOutputHierarchyCreator - Actor behind a participant.
type SkillsDuplicateOutputHierarchyCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDuplicateOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDuplicateOutputHierarchyCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorMember represents the skills duplicate output hierarchy fork creator organization actor member type.
type SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor struct {
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
	Member   *SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDuplicateOutputHierarchyForkCreatorConsumer struct {
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

// SkillsDuplicateOutputHierarchyForkCreator - Actor behind a participant.
type SkillsDuplicateOutputHierarchyForkCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDuplicateOutputHierarchyForkCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorMember represents the skills duplicate output hierarchy fork original creator organization actor member type.
type SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	Member   *SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer struct {
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

// SkillsDuplicateOutputHierarchyForkOriginalCreator - Actor behind a participant.
type SkillsDuplicateOutputHierarchyForkOriginalCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDuplicateOutputHierarchyFork - Fork details. `null` unless the skill is a fork.
type SkillsDuplicateOutputHierarchyFork struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// ParentSkillId - Upstream skill the fork was created from, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// Creator - Actor behind a participant.
	Creator *SkillsDuplicateOutputHierarchyForkCreator `json:"creator,omitempty"`
	// OriginalCreator - Actor behind a participant.
	OriginalCreator *SkillsDuplicateOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsDuplicateOutputHierarchyEntity - Skill entity the versioned skill content belongs to.
type SkillsDuplicateOutputHierarchyEntity struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ParentSkillId - Upstream skill, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchy - Where the skill came from.
type SkillsDuplicateOutputHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - How the skill was created. `root` for original skills, `fork` for consumer forks, `duplicated` for copies.
	Type string `json:"type"`
	// ParentSkillId - Parent skill, by ID. `null` for root skills.
	ParentSkillId *string `json:"parent_skill_id,omitempty"`
	// Creator - Actor behind a participant.
	Creator *SkillsDuplicateOutputHierarchyCreator `json:"creator,omitempty"`
	// Fork - Fork details. `null` unless the skill is a fork.
	Fork *SkillsDuplicateOutputHierarchyFork `json:"fork,omitempty"`
	// Entity - Skill entity the versioned skill content belongs to.
	Entity SkillsDuplicateOutputHierarchyEntity `json:"entity"`
}

// SkillsDuplicateOutputIntegrationsConfiguration - Which settings integration instances may customize.
type SkillsDuplicateOutputIntegrationsConfiguration struct {
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters bool `json:"can_attach_custom_tool_filters"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters bool `json:"can_override_tool_filters"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools bool `json:"enable_callback_tools"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *bool `json:"use_integration_name_in_tool_names,omitempty"`
}

// SkillsDuplicateOutputIntegrations represents the skills duplicate output integrations type.
type SkillsDuplicateOutputIntegrations struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Which settings integration instances may customize.
	Configuration SkillsDuplicateOutputIntegrationsConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsDuplicateOutputProviders represents the skills duplicate output providers type.
type SkillsDuplicateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutput represents the skills duplicate output type.
type SkillsDuplicateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the skill is available, archived, or deleted.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - Public URL of the skill image.
	ImageUrl string `json:"image_url"`
	// Image - Image shown for the skill.
	Image SkillsDuplicateOutputImage `json:"image"`
	// ClientName - Client-facing skill name.
	ClientName string `json:"client_name"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// StoreId - Store holding the skill, by ID.
	StoreId string `json:"store_id"`
	// Hierarchy - Where the skill came from.
	Hierarchy SkillsDuplicateOutputHierarchy `json:"hierarchy"`
	// Integrations - Integrations attached to the skill.
	Integrations []SkillsDuplicateOutputIntegrations `json:"integrations"`
	// Providers - Providers attached to the skill.
	Providers []SkillsDuplicateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsDuplicateOutputFromJSON deserializes JSON data into a SkillsDuplicateOutput.
func MapSkillsDuplicateOutputFromJSON(data []byte) (*SkillsDuplicateOutput, error) {
	var v SkillsDuplicateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDuplicateOutputToJSON serializes a SkillsDuplicateOutput to JSON.
func MapSkillsDuplicateOutputToJSON(v *SkillsDuplicateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsDuplicateBody represents the skills duplicate body type.
type SkillsDuplicateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ClientName - Client-facing skill name. Defaults to `name`. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// Metadata - Custom data to identify or organize the skill.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapSkillsDuplicateBodyFromJSON deserializes JSON data into a SkillsDuplicateBody.
func MapSkillsDuplicateBodyFromJSON(data []byte) (*SkillsDuplicateBody, error) {
	var v SkillsDuplicateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDuplicateBodyToJSON serializes a SkillsDuplicateBody to JSON.
func MapSkillsDuplicateBodyToJSON(v *SkillsDuplicateBody) ([]byte, error) {
	return json.Marshal(v)
}
