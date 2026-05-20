package authorizations

import (
	"encoding/json"
	"time"
)

// TestHelpersConsumerOauthAuthorizationsCreateOutput represents the test helpers consumer oauth authorizations create output type.
type TestHelpersConsumerOauthAuthorizationsCreateOutput struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// MapTestHelpersConsumerOauthAuthorizationsCreateOutputFromJSON deserializes JSON data into a TestHelpersConsumerOauthAuthorizationsCreateOutput.
func MapTestHelpersConsumerOauthAuthorizationsCreateOutputFromJSON(data []byte) (*TestHelpersConsumerOauthAuthorizationsCreateOutput, error) {
	var v TestHelpersConsumerOauthAuthorizationsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapTestHelpersConsumerOauthAuthorizationsCreateOutputToJSON serializes a TestHelpersConsumerOauthAuthorizationsCreateOutput to JSON.
func MapTestHelpersConsumerOauthAuthorizationsCreateOutputToJSON(v *TestHelpersConsumerOauthAuthorizationsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// TestHelpersConsumerOauthAuthorizationsCreateBody represents the test helpers consumer oauth authorizations create body type.
type TestHelpersConsumerOauthAuthorizationsCreateBody struct {
	InstanceId         string `json:"instance_id"`
	Url                string `json:"url"`
	ConsumerProfileId  string `json:"consumer_profile_id"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// MapTestHelpersConsumerOauthAuthorizationsCreateBodyFromJSON deserializes JSON data into a TestHelpersConsumerOauthAuthorizationsCreateBody.
func MapTestHelpersConsumerOauthAuthorizationsCreateBodyFromJSON(data []byte) (*TestHelpersConsumerOauthAuthorizationsCreateBody, error) {
	var v TestHelpersConsumerOauthAuthorizationsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapTestHelpersConsumerOauthAuthorizationsCreateBodyToJSON serializes a TestHelpersConsumerOauthAuthorizationsCreateBody to JSON.
func MapTestHelpersConsumerOauthAuthorizationsCreateBodyToJSON(v *TestHelpersConsumerOauthAuthorizationsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
