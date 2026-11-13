package consumers

import (
	"encoding/json"
)

// ConsumersCreateMemberSessionOutput represents the consumers create member session output type.
type ConsumersCreateMemberSessionOutput struct {
	AccessToken   string  `json:"access_token"`
	CatalogItemId *string `json:"catalog_item_id,omitempty"`
}

// MapConsumersCreateMemberSessionOutputFromJSON deserializes JSON data into a ConsumersCreateMemberSessionOutput.
func MapConsumersCreateMemberSessionOutputFromJSON(data []byte) (*ConsumersCreateMemberSessionOutput, error) {
	var v ConsumersCreateMemberSessionOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersCreateMemberSessionOutputToJSON serializes a ConsumersCreateMemberSessionOutput to JSON.
func MapConsumersCreateMemberSessionOutputToJSON(v *ConsumersCreateMemberSessionOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersCreateMemberSessionBody represents the consumers create member session body type.
type ConsumersCreateMemberSessionBody struct {
	IntegrationId *string `json:"integration_id,omitempty"`
}

// MapConsumersCreateMemberSessionBodyFromJSON deserializes JSON data into a ConsumersCreateMemberSessionBody.
func MapConsumersCreateMemberSessionBodyFromJSON(data []byte) (*ConsumersCreateMemberSessionBody, error) {
	var v ConsumersCreateMemberSessionBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersCreateMemberSessionBodyToJSON serializes a ConsumersCreateMemberSessionBody to JSON.
func MapConsumersCreateMemberSessionBodyToJSON(v *ConsumersCreateMemberSessionBody) ([]byte, error) {
	return json.Marshal(v)
}
