package instances

import (
	"encoding/json"
	"time"
)

// AgentsInstancesListOutputItemsAgentClient represents the agents instances list output items agent client type.
type AgentsInstancesListOutputItemsAgentClient struct {
	Object          string     `json:"object"`
	Id              string     `json:"id"`
	Type            string     `json:"type"`
	Name            string     `json:"name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
}

// AgentsInstancesListOutputItems represents the agents instances list output items type.
type AgentsInstancesListOutputItems struct {
	Object          string                                     `json:"object"`
	Id              string                                     `json:"id"`
	Type            string                                     `json:"type"`
	Name            string                                     `json:"name"`
	Version         *string                                    `json:"version,omitempty"`
	Description     *string                                    `json:"description,omitempty"`
	AgentId         string                                     `json:"agent_id"`
	AgentClient     *AgentsInstancesListOutputItemsAgentClient `json:"agent_client,omitempty"`
	CreatedAt       time.Time                                  `json:"created_at"`
	UpdatedAt       time.Time                                  `json:"updated_at"`
	LastConnectedAt *time.Time                                 `json:"last_connected_at,omitempty"`
}

// AgentsInstancesListOutputPagination represents the agents instances list output pagination type.
type AgentsInstancesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// AgentsInstancesListOutput represents the agents instances list output type.
type AgentsInstancesListOutput struct {
	Items      []AgentsInstancesListOutputItems    `json:"items"`
	Pagination AgentsInstancesListOutputPagination `json:"pagination"`
}

// MapAgentsInstancesListOutputFromJSON deserializes JSON data into a AgentsInstancesListOutput.
func MapAgentsInstancesListOutputFromJSON(data []byte) (*AgentsInstancesListOutput, error) {
	var v AgentsInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsInstancesListOutputToJSON serializes a AgentsInstancesListOutput to JSON.
func MapAgentsInstancesListOutputToJSON(v *AgentsInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AgentsInstancesListQueryCreatedAt - Filter agent instance creation time by date range
type AgentsInstancesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for agent instance creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for agent instance creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsInstancesListQueryUpdatedAt - Filter agent instance last update time by date range
type AgentsInstancesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for agent instance last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for agent instance last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsInstancesListQuery represents the agents instances list query type.
type AgentsInstancesListQuery struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Type          *any     `json:"type,omitempty"`
	Id            *any     `json:"id,omitempty"`
	AgentClientId *any     `json:"agent_client_id,omitempty"`
	// CreatedAt - Filter agent instance creation time by date range
	CreatedAt *AgentsInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter agent instance last update time by date range
	UpdatedAt *AgentsInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapAgentsInstancesListQueryFromJSON deserializes JSON data into a AgentsInstancesListQuery.
func MapAgentsInstancesListQueryFromJSON(data []byte) (*AgentsInstancesListQuery, error) {
	var v AgentsInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsInstancesListQueryToJSON serializes a AgentsInstancesListQuery to JSON.
func MapAgentsInstancesListQueryToJSON(v *AgentsInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
