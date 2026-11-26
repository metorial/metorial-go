package instances

import (
	"encoding/json"
	"time"
)

// AgentsInstancesGetOutputAgentClient - Client this instance connects with.
type AgentsInstancesGetOutputAgentClient struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the client connects.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// LastConnectedAt - When the client last connected. `null` when never connected.
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
}

// AgentsInstancesGetOutput represents the agents instances get output type.
type AgentsInstancesGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the agent instance connects and runs.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Client version. `null` when the instance reports no version.
	Version *string `json:"version,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AgentId - Agent this instance belongs to.
	AgentId string `json:"agent_id"`
	// AgentClient - Client this instance connects with.
	AgentClient *AgentsInstancesGetOutputAgentClient `json:"agent_client,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// LastConnectedAt - When the instance last connected. `null` when never connected.
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
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
