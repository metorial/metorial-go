package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/virtualmodeltemplates"
)

// AiGatewayVirtualModelTemplatesEndpoint provides access to ready-made virtual model definitions for common routing setups. Pick the models they ask for to create a virtual model from one.
type AiGatewayVirtualModelTemplatesEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayVirtualModelTemplatesEndpoint creates a new AiGatewayVirtualModelTemplatesEndpoint.
func NewAiGatewayVirtualModelTemplatesEndpoint(client *endpoint.Client) *AiGatewayVirtualModelTemplatesEndpoint {
	return &AiGatewayVirtualModelTemplatesEndpoint{client: client}
}

// List returns every template. The list is not paginated.
func (e *AiGatewayVirtualModelTemplatesEndpoint) List(instanceId string) (*virtualmodeltemplates.AiGatewayVirtualModelTemplatesListOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "ai-gateway", "virtual-model-templates"},
	}
	var result virtualmodeltemplates.AiGatewayVirtualModelTemplatesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
