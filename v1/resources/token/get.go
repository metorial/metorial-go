package token

import (
	"encoding/json"
)

// TokenGetOutputOrganization - Organization the token is scoped to.
type TokenGetOutputOrganization struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name of the organization.
	Slug string `json:"slug"`
}

// TokenGetOutputInstance - Instance the token is scoped to. `null` for organization-level tokens.
type TokenGetOutputInstance struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name of the instance.
	Slug string `json:"slug"`
	// ProjectId - Project this instance belongs to.
	ProjectId string `json:"project_id"`
}

// TokenGetOutputProject - Project the token’s instance belongs to. `null` without an instance.
type TokenGetOutputProject struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name of the project.
	Slug string `json:"slug"`
}

// TokenGetOutputActor - Organization actor the token acts as.
type TokenGetOutputActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Whether the actor is an organization member or a machine identity.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
}

// TokenGetOutputMember - Organization member the token belongs to, if any.
type TokenGetOutputMember struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// TokenGetOutputUser - User the token belongs to, for user session tokens.
type TokenGetOutputUser struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// TokenGetOutput represents the token get output type.
type TokenGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - What kind of credential this is, for example `user_auth_token` or `oauth_access_token`.
	Type string `json:"type"`
	// Organization - Organization the token is scoped to.
	Organization *TokenGetOutputOrganization `json:"organization,omitempty"`
	// Instance - Instance the token is scoped to. `null` for organization-level tokens.
	Instance *TokenGetOutputInstance `json:"instance,omitempty"`
	// Project - Project the token’s instance belongs to. `null` without an instance.
	Project *TokenGetOutputProject `json:"project,omitempty"`
	// Actor - Organization actor the token acts as.
	Actor *TokenGetOutputActor `json:"actor,omitempty"`
	// Member - Organization member the token belongs to, if any.
	Member *TokenGetOutputMember `json:"member,omitempty"`
	// User - User the token belongs to, for user session tokens.
	User *TokenGetOutputUser `json:"user,omitempty"`
}

// MapTokenGetOutputFromJSON deserializes JSON data into a TokenGetOutput.
func MapTokenGetOutputFromJSON(data []byte) (*TokenGetOutput, error) {
	var v TokenGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapTokenGetOutputToJSON serializes a TokenGetOutput to JSON.
func MapTokenGetOutputToJSON(v *TokenGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
