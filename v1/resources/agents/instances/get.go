package instances

import (
	"encoding/json"
	"time"
)

// AgentsInstancesGetOutputAgentClient represents the agents instances get output agent client type.
type AgentsInstancesGetOutputAgentClient struct {
	Object          string     `json:"object"`
	Id              string     `json:"id"`
	Type            string     `json:"type"`
	Name            string     `json:"name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
}

// AgentsInstancesGetOutput represents the agents instances get output type.
type AgentsInstancesGetOutput struct {
	Object          string                               `json:"object"`
	Id              string                               `json:"id"`
	Type            string                               `json:"type"`
	Name            string                               `json:"name"`
	Version         *string                              `json:"version,omitempty"`
	Description     *string                              `json:"description,omitempty"`
	AgentId         string                               `json:"agent_id"`
	AgentClient     *AgentsInstancesGetOutputAgentClient `json:"agent_client,omitempty"`
	CreatedAt       time.Time                            `json:"created_at"`
	UpdatedAt       time.Time                            `json:"updated_at"`
	LastConnectedAt *time.Time                           `json:"last_connected_at,omitempty"`
}

// MapAgentsInstancesGetOutputFromJSON deserializes JSON data into a AgentsInstancesGetOutput.
func MapAgentsInstancesGetOutputFromJSON(data []byte) (*AgentsInstancesGetOutput, error) {
	var v AgentsInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsInstancesGetOutputToJSON serializes a AgentsInstancesGetOutput to JSON.
func MapAgentsInstancesGetOutputToJSON(v *AgentsInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
