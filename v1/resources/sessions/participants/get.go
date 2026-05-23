package participants

import (
	"encoding/json"
	"time"
)

// SessionsParticipantsGetOutputData - Participant payload data
type SessionsParticipantsGetOutputData struct {
	// Identifier - Participant-specific identifier within the payload
	Identifier string `json:"identifier"`
	// Name - Participant-specific display name within the payload
	Name string `json:"name"`
}

// SessionsParticipantsGetOutput represents the sessions participants get output type.
type SessionsParticipantsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session participant identifier
	Id string `json:"id"`
	// Type - Participant type
	Type string `json:"type"`
	// Identifier - Participant identifier
	Identifier string `json:"identifier"`
	// Name - Display name
	Name string `json:"name"`
	// Data - Participant payload data
	Data SessionsParticipantsGetOutputData `json:"data"`
	// ProviderId - Provider ID if associated
	ProviderId      *string `json:"provider_id,omitempty"`
	ConnectionType  *string `json:"connection_type,omitempty"`
	AgentId         *string `json:"agent_id,omitempty"`
	AgentInstanceId *string `json:"agent_instance_id,omitempty"`
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	IdentityId      *string `json:"identity_id,omitempty"`
	AgentActorId    *string `json:"agent_actor_id,omitempty"`
	AgentClientId   *string `json:"agent_client_id,omitempty"`
	ConsumerId      *string `json:"consumer_id,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
}

// MapSessionsParticipantsGetOutputFromJSON deserializes JSON data into a SessionsParticipantsGetOutput.
func MapSessionsParticipantsGetOutputFromJSON(data []byte) (*SessionsParticipantsGetOutput, error) {
	var v SessionsParticipantsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsParticipantsGetOutputToJSON serializes a SessionsParticipantsGetOutput to JSON.
func MapSessionsParticipantsGetOutputToJSON(v *SessionsParticipantsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
