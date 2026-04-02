package apikeys

import (
	"encoding/json"
	"time"
)

// ApiKeysRotateOutputMachineAccessActorTeams - The teams the actor belongs to
type ApiKeysRotateOutputMachineAccessActorTeams struct {
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

// ApiKeysRotateOutputMachineAccessActor represents the api keys rotate output machine access actor type.
type ApiKeysRotateOutputMachineAccessActor struct {
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
	ImageUrl string                                       `json:"image_url"`
	Teams    []ApiKeysRotateOutputMachineAccessActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ApiKeysRotateOutputMachineAccessInstanceProject represents the api keys rotate output machine access instance project type.
type ApiKeysRotateOutputMachineAccessInstanceProject struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The project's unique identifier
	Id string `json:"id"`
	// Status - The project's status
	Status string `json:"status"`
	// Slug - The project's slug
	Slug string `json:"slug"`
	// Name - The project's name
	Name string `json:"name"`
	// OrganizationId - The organization's unique identifier
	OrganizationId string `json:"organization_id"`
	// CreatedAt - The project's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The project's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ApiKeysRotateOutputMachineAccessInstance represents the api keys rotate output machine access instance type.
type ApiKeysRotateOutputMachineAccessInstance struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The instance's unique identifier
	Id string `json:"id"`
	// Slug - The instance's slug
	Slug string `json:"slug"`
	// Name - The instance's name
	Name string `json:"name"`
	// OrganizationId - The organization's unique identifier
	OrganizationId string `json:"organization_id"`
	// Type - The instance's type
	Type string `json:"type"`
	// CreatedAt - The instance's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The instance's last update date
	UpdatedAt time.Time                                       `json:"updated_at"`
	Project   ApiKeysRotateOutputMachineAccessInstanceProject `json:"project"`
}

// ApiKeysRotateOutputMachineAccessOrganization represents the api keys rotate output machine access organization type.
type ApiKeysRotateOutputMachineAccessOrganization struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization's unique identifier
	Id string `json:"id"`
	// Type - The organization's type
	Type string `json:"type"`
	// Slug - The organization's slug
	Slug string `json:"slug"`
	// Name - The organization's name
	Name string `json:"name"`
	// ImageUrl - The organization's image URL
	ImageUrl string `json:"image_url"`
	// CreatedAt - The organization's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ApiKeysRotateOutputMachineAccessUser represents the api keys rotate output machine access user type.
type ApiKeysRotateOutputMachineAccessUser struct {
	// Object - Type of the object, fixed as 'user'
	Object string `json:"object"`
	// Id - The user's unique identifier
	Id string `json:"id"`
	// Status - The user's status
	Status string `json:"status"`
	// Type - The user's type
	Type string `json:"type"`
	// Email - The user's email address
	Email string `json:"email"`
	// Name - The user's full name
	Name string `json:"name"`
	// FirstName - The user's first name
	FirstName string `json:"first_name"`
	// LastName - The user's last name
	LastName string `json:"last_name"`
	// ImageUrl - The user's image URL
	ImageUrl string `json:"image_url"`
	// CreatedAt - The user's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The user's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ApiKeysRotateOutputMachineAccess represents the api keys rotate output machine access type.
type ApiKeysRotateOutputMachineAccess struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The machineAccess's unique identifier
	Id string `json:"id"`
	// Status - The machineAccess's status
	Status string `json:"status"`
	// Type - The machineAccess's type
	Type string `json:"type"`
	// Name - The machineAccess's name
	Name string `json:"name"`
	// LastUsedAt - The machineAccess's last used date
	LastUsedAt time.Time `json:"last_used_at"`
	// CreatedAt - The machineAccess's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The machineAccess's last update date
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt - The machineAccess's deletion date
	DeletedAt    time.Time                                     `json:"deleted_at"`
	Actor        *ApiKeysRotateOutputMachineAccessActor        `json:"actor,omitempty"`
	Instance     *ApiKeysRotateOutputMachineAccessInstance     `json:"instance,omitempty"`
	Organization *ApiKeysRotateOutputMachineAccessOrganization `json:"organization,omitempty"`
	User         *ApiKeysRotateOutputMachineAccessUser         `json:"user,omitempty"`
}

// ApiKeysRotateOutput represents the api keys rotate output type.
type ApiKeysRotateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The apiKey's unique identifier
	Id string `json:"id"`
	// Status - The apiKey's status
	Status string `json:"status"`
	// SecretRedacted - The apiKey's secret, redacted
	SecretRedacted string `json:"secret_redacted"`
	// Secret - The apiKey's secret
	Secret *string `json:"secret,omitempty"`
	// Type - The apiKey's type
	Type string `json:"type"`
	// Name - The apiKey's name
	Name string `json:"name"`
	// Description - The apiKey's description
	Description *string `json:"description,omitempty"`
	// IpFilters - List of allowed IP addresses or CIDR ranges for this API key
	IpFilters     []string                         `json:"ip_filters"`
	MachineAccess ApiKeysRotateOutputMachineAccess `json:"machine_access"`
	// DeletedAt - The apiKey's deletion date
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// LastUsedAt - The apiKey's last usage date
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// ExpiresAt - The apiKey's expiration date
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// CreatedAt - The apiKey's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The apiKey's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// MapApiKeysRotateOutputFromJSON deserializes JSON data into a ApiKeysRotateOutput.
func MapApiKeysRotateOutputFromJSON(data []byte) (*ApiKeysRotateOutput, error) {
	var v ApiKeysRotateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapApiKeysRotateOutputToJSON serializes a ApiKeysRotateOutput to JSON.
func MapApiKeysRotateOutputToJSON(v *ApiKeysRotateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ApiKeysRotateBody represents the api keys rotate body type.
type ApiKeysRotateBody struct {
	CurrentExpiresAt *time.Time `json:"current_expires_at,omitempty"`
}

// MapApiKeysRotateBodyFromJSON deserializes JSON data into a ApiKeysRotateBody.
func MapApiKeysRotateBodyFromJSON(data []byte) (*ApiKeysRotateBody, error) {
	var v ApiKeysRotateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapApiKeysRotateBodyToJSON serializes a ApiKeysRotateBody to JSON.
func MapApiKeysRotateBodyToJSON(v *ApiKeysRotateBody) ([]byte, error) {
	return json.Marshal(v)
}
