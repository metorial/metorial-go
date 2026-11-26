package instances

import (
	"encoding/json"
	"time"
)

// ChatInstancesGetOutputIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesGetOutputIdentity struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// UserId - The provider-side user or app id this connection is authorized as
	UserId string `json:"user_id"`
	// Name - Display name of the authorized identity
	Name string `json:"name"`
	// Username - Username or handle of the authorized identity
	Username string `json:"username"`
	// ProviderType - The kind of identity as classified by the provider, e.g. a human user vs. an app/bot
	ProviderType string `json:"provider_type"`
	// Email - Email address of the authorized identity, if available
	Email *string `json:"email,omitempty"`
	// ImageUrl - Avatar or image URL of the authorized identity, if available
	ImageUrl *string `json:"image_url,omitempty"`
}

// ChatInstancesGetOutput represents the chat instances get output type.
type ChatInstancesGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the instance. Archived instances stay listable.
	Status string `json:"status"`
	// ChatConnectionId - The chat connection this instance belongs to
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Identity - Account this instance is authorized as on the chat provider. `null` until observed.
	Identity *ChatInstancesGetOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatInstancesGetOutputFromJSON deserializes JSON data into a ChatInstancesGetOutput.
func MapChatInstancesGetOutputFromJSON(data []byte) (*ChatInstancesGetOutput, error) {
	var v ChatInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesGetOutputToJSON serializes a ChatInstancesGetOutput to JSON.
func MapChatInstancesGetOutputToJSON(v *ChatInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
