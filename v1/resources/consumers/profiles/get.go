package profiles

import (
	"encoding/json"
	"time"
)

// ConsumersProfilesGetOutputGroupsGroup represents the consumers profiles get output groups group type.
type ConsumersProfilesGetOutputGroupsGroup struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	SsoGroupIds []string  `json:"sso_group_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ConsumersProfilesGetOutputGroups represents the consumers profiles get output groups type.
type ConsumersProfilesGetOutputGroups struct {
	Object      string                                `json:"object"`
	Group       ConsumersProfilesGetOutputGroupsGroup `json:"group"`
	AssignedVia string                                `json:"assigned_via"`
}

// ConsumersProfilesGetOutput represents the consumers profiles get output type.
type ConsumersProfilesGetOutput struct {
	Object     string                              `json:"object"`
	Id         string                              `json:"id"`
	Name       string                              `json:"name"`
	Email      string                              `json:"email"`
	ImageUrl   string                              `json:"image_url"`
	Groups     *[]ConsumersProfilesGetOutputGroups `json:"groups,omitempty"`
	ConsumerId string                              `json:"consumer_id"`
	CreatedAt  time.Time                           `json:"created_at"`
	UpdatedAt  time.Time                           `json:"updated_at"`
}

// MapConsumersProfilesGetOutputFromJSON deserializes JSON data into a ConsumersProfilesGetOutput.
func MapConsumersProfilesGetOutputFromJSON(data []byte) (*ConsumersProfilesGetOutput, error) {
	var v ConsumersProfilesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesGetOutputToJSON serializes a ConsumersProfilesGetOutput to JSON.
func MapConsumersProfilesGetOutputToJSON(v *ConsumersProfilesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
