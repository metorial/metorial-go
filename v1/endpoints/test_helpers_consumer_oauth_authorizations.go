package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/testhelpers/consumeroauth/authorizations"
)

// TestHelpersConsumerOauthAuthorizationsEndpoint provides access to helpers for testing consumer OAuth flows.
type TestHelpersConsumerOauthAuthorizationsEndpoint struct {
	client *endpoint.Client
}

// NewTestHelpersConsumerOauthAuthorizationsEndpoint creates a new TestHelpersConsumerOauthAuthorizationsEndpoint.
func NewTestHelpersConsumerOauthAuthorizationsEndpoint(client *endpoint.Client) *TestHelpersConsumerOauthAuthorizationsEndpoint {
	return &TestHelpersConsumerOauthAuthorizationsEndpoint{client: client}
}

// TestHelpersConsumerOauthAuthorizationsEndpointCreateBody contains the request body for Create.
type TestHelpersConsumerOauthAuthorizationsEndpointCreateBody struct {
	InstanceId         string `json:"instance_id"`
	Url                string `json:"url"`
	ConsumerProfileId  string `json:"consumer_profile_id"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// Create creates a single-use test authorization token for a consumer OAuth authorize URL.
func (e *TestHelpersConsumerOauthAuthorizationsEndpoint) Create(body *TestHelpersConsumerOauthAuthorizationsEndpointCreateBody) (*authorizations.TestHelpersConsumerOauthAuthorizationsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"test-helpers", "consumer-oauth-authorizations"},
		Body: body,
	}
	var result authorizations.TestHelpersConsumerOauthAuthorizationsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
