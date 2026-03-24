package token

import (
	"encoding/json"
)

// TokenGetOutputOrganization represents the token get output organization type.
type TokenGetOutputOrganization struct {
	// Object - String representing the organization's type
	Object string `json:"object"`
	// Id - The organization's unique identifier
	Id string `json:"id"`
	// Name - The organization's name
	Name string `json:"name"`
	// Slug - The organization's slug
	Slug string `json:"slug"`
}

// TokenGetOutputInstance represents the token get output instance type.
type TokenGetOutputInstance struct {
	// Object - String representing the instance's type
	Object string `json:"object"`
	// Id - The instance's unique identifier
	Id string `json:"id"`
	// Name - The instance's name
	Name string `json:"name"`
	// Slug - The instance's slug
	Slug string `json:"slug"`
	// ProjectId - The instance's project ID
	ProjectId string `json:"project_id"`
}

// TokenGetOutputProject represents the token get output project type.
type TokenGetOutputProject struct {
	// Object - String representing the project's type
	Object string `json:"object"`
	// Id - The project's unique identifier
	Id string `json:"id"`
	// Name - The project's name
	Name string `json:"name"`
	// Slug - The project's slug
	Slug string `json:"slug"`
}

// TokenGetOutputActor represents the token get output actor type.
type TokenGetOutputActor struct {
	// Object - String representing the organization actor's type
	Object string `json:"object"`
	// Id - The organization actor's unique identifier
	Id string `json:"id"`
	// Type - The organization actor's type
	Type string `json:"type"`
	// Name - The organization actor's name
	Name string `json:"name"`
}

// TokenGetOutputMember represents the token get output member type.
type TokenGetOutputMember struct {
	// Object - String representing the organization member's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Name - The organization member's name
	Name string `json:"name"`
}

// TokenGetOutputUser represents the token get output user type.
type TokenGetOutputUser struct {
	// Object - String representing the user's type
	Object string `json:"object"`
	// Id - The user's unique identifier
	Id string `json:"id"`
	// Name - The user's name
	Name string `json:"name"`
}

// TokenGetOutput represents the token get output type.
type TokenGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The token's type
	Type         string                      `json:"type"`
	Organization *TokenGetOutputOrganization `json:"organization,omitempty"`
	Instance     *TokenGetOutputInstance     `json:"instance,omitempty"`
	Project      *TokenGetOutputProject      `json:"project,omitempty"`
	Actor        *TokenGetOutputActor        `json:"actor,omitempty"`
	Member       *TokenGetOutputMember       `json:"member,omitempty"`
	User         *TokenGetOutputUser         `json:"user,omitempty"`
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
