package skills

import (
	"encoding/json"
	"time"
)

// SkillsShareOutputImage - Image shown for the skill.
type SkillsShareOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsShareOutputHierarchyCreatorOrganizationActorMember represents the skills share output hierarchy creator organization actor member type.
type SkillsShareOutputHierarchyCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsShareOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                    `json:"image_url"`
	Member   *SkillsShareOutputHierarchyCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsShareOutputHierarchyCreatorConsumer struct {
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

// SkillsShareOutputHierarchyCreator - Actor behind a participant.
type SkillsShareOutputHierarchyCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsShareOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsShareOutputHierarchyCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyForkCreatorOrganizationActorMember represents the skills share output hierarchy fork creator organization actor member type.
type SkillsShareOutputHierarchyForkCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyForkCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsShareOutputHierarchyForkCreatorOrganizationActor struct {
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
	Member   *SkillsShareOutputHierarchyForkCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsShareOutputHierarchyForkCreatorConsumer struct {
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

// SkillsShareOutputHierarchyForkCreator - Actor behind a participant.
type SkillsShareOutputHierarchyForkCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsShareOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsShareOutputHierarchyForkCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember represents the skills share output hierarchy fork original creator organization actor member type.
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	Member   *SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsShareOutputHierarchyForkOriginalCreatorConsumer struct {
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

// SkillsShareOutputHierarchyForkOriginalCreator - Actor behind a participant.
type SkillsShareOutputHierarchyForkOriginalCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsShareOutputHierarchyForkOriginalCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyFork - Fork details. `null` unless the skill is a fork.
type SkillsShareOutputHierarchyFork struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// ParentSkillId - Upstream skill the fork was created from, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// Creator - Actor behind a participant.
	Creator *SkillsShareOutputHierarchyForkCreator `json:"creator,omitempty"`
	// OriginalCreator - Actor behind a participant.
	OriginalCreator *SkillsShareOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsShareOutputHierarchyEntity - Skill entity the versioned skill content belongs to.
type SkillsShareOutputHierarchyEntity struct {
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

// SkillsShareOutputHierarchy - Where the skill came from.
type SkillsShareOutputHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - How the skill was created. `root` for original skills, `fork` for consumer forks, `duplicated` for copies.
	Type string `json:"type"`
	// ParentSkillId - Parent skill, by ID. `null` for root skills.
	ParentSkillId *string `json:"parent_skill_id,omitempty"`
	// Creator - Actor behind a participant.
	Creator *SkillsShareOutputHierarchyCreator `json:"creator,omitempty"`
	// Fork - Fork details. `null` unless the skill is a fork.
	Fork *SkillsShareOutputHierarchyFork `json:"fork,omitempty"`
	// Entity - Skill entity the versioned skill content belongs to.
	Entity SkillsShareOutputHierarchyEntity `json:"entity"`
}

// SkillsShareOutputIntegrationsConfiguration - Which settings integration instances may customize.
type SkillsShareOutputIntegrationsConfiguration struct {
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

// SkillsShareOutputIntegrations represents the skills share output integrations type.
type SkillsShareOutputIntegrations struct {
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
	Configuration SkillsShareOutputIntegrationsConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsShareOutputProviders represents the skills share output providers type.
type SkillsShareOutputProviders struct {
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

// SkillsShareOutput represents the skills share output type.
type SkillsShareOutput struct {
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
	Image SkillsShareOutputImage `json:"image"`
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
	Hierarchy SkillsShareOutputHierarchy `json:"hierarchy"`
	// Integrations - Integrations attached to the skill.
	Integrations []SkillsShareOutputIntegrations `json:"integrations"`
	// Providers - Providers attached to the skill.
	Providers []SkillsShareOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsShareOutputFromJSON deserializes JSON data into a SkillsShareOutput.
func MapSkillsShareOutputFromJSON(data []byte) (*SkillsShareOutput, error) {
	var v SkillsShareOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsShareOutputToJSON serializes a SkillsShareOutput to JSON.
func MapSkillsShareOutputToJSON(v *SkillsShareOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsShareBody represents the skills share body type.
type SkillsShareBody struct {
	// ConsumerProfileIds - Consumers to share the skill with, by ID.
	ConsumerProfileIds *[]string `json:"consumer_profile_ids,omitempty"`
	// OrganizationMemberIds - Must be omitted. Organization member access cannot be changed.
	OrganizationMemberIds *[]string `json:"organization_member_ids,omitempty"`
	// Permission - For consumers with inherited edit and share access, `write` and `owner` add portal visibility. Consumers must own the skill to grant ownership. Only admins can remove or downgrade an owner.
	Permission string `json:"permission"`
}

// MapSkillsShareBodyFromJSON deserializes JSON data into a SkillsShareBody.
func MapSkillsShareBodyFromJSON(data []byte) (*SkillsShareBody, error) {
	var v SkillsShareBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsShareBodyToJSON serializes a SkillsShareBody to JSON.
func MapSkillsShareBodyToJSON(v *SkillsShareBody) ([]byte, error) {
	return json.Marshal(v)
}
