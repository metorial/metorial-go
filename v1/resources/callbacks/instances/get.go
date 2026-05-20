package instances

import (
	"encoding/json"
	"time"
)

// CallbacksInstancesGetOutputDeployment represents the callbacks instances get output deployment type.
type CallbacksInstancesGetOutputDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksInstancesGetOutputConfig represents the callbacks instances get output config type.
type CallbacksInstancesGetOutputConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksInstancesGetOutputAuthConfig represents the callbacks instances get output auth config type.
type CallbacksInstancesGetOutputAuthConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksInstancesGetOutputTriggersProviderTriggerInputSchema represents the callbacks instances get output triggers provider trigger input schema type.
type CallbacksInstancesGetOutputTriggersProviderTriggerInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the trigger payload input shape
	Schema map[string]any `json:"schema"`
}

// CallbacksInstancesGetOutputTriggersProviderTriggerOutputSchema represents the callbacks instances get output triggers provider trigger output schema type.
type CallbacksInstancesGetOutputTriggersProviderTriggerOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the trigger delivery output shape
	Schema map[string]any `json:"schema"`
}

// CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoRegistration represents the callbacks instances get output triggers provider trigger invocation auto registration type.
type CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoRegistration struct {
	// Status - Whether automatic webhook registration is supported
	Status string `json:"status"`
}

// CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoUnregistration represents the callbacks instances get output triggers provider trigger invocation auto unregistration type.
type CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoUnregistration struct {
	// Status - Whether automatic webhook removal is supported
	Status string `json:"status"`
}

// CallbacksInstancesGetOutputTriggersProviderTriggerInvocation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CallbacksInstancesGetOutputTriggersProviderTriggerInvocation struct {
	Type *string `json:"type,omitempty"`
	// IntervalSeconds - Polling interval in seconds for polling-based triggers
	IntervalSeconds    *float64                                                                        `json:"interval_seconds,omitempty"`
	AutoRegistration   *CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoRegistration   `json:"auto_registration,omitempty"`
	AutoUnregistration *CallbacksInstancesGetOutputTriggersProviderTriggerInvocationAutoUnregistration `json:"auto_unregistration,omitempty"`
}

// CallbacksInstancesGetOutputTriggersProviderTrigger represents the callbacks instances get output triggers provider trigger type.
type CallbacksInstancesGetOutputTriggersProviderTrigger struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider trigger identifier
	Id string `json:"id"`
	// Key - Trigger key used when subscribing callbacks
	Key string `json:"key"`
	// Name - Display name of the trigger
	Name string `json:"name"`
	// Description - Trigger description
	Description  *string                                                         `json:"description,omitempty"`
	InputSchema  *CallbacksInstancesGetOutputTriggersProviderTriggerInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *CallbacksInstancesGetOutputTriggersProviderTriggerOutputSchema `json:"output_schema,omitempty"`
	Invocation   CallbacksInstancesGetOutputTriggersProviderTriggerInvocation    `json:"invocation"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Provider specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksInstancesGetOutputTriggers represents the callbacks instances get output triggers type.
type CallbacksInstancesGetOutputTriggers struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique receiver trigger identifier
	Id string `json:"id"`
	// Source - How this trigger is invoked by the provider backend
	Source string `json:"source"`
	// PollIntervalSeconds - Polling interval in seconds when the trigger uses polling
	PollIntervalSeconds *float64 `json:"poll_interval_seconds,omitempty"`
	// NextPollAt - Next scheduled poll timestamp for polling triggers
	NextPollAt *time.Time `json:"next_poll_at,omitempty"`
	// LastPolledAt - Last successful poll timestamp for polling triggers
	LastPolledAt *time.Time `json:"last_polled_at,omitempty"`
	// WebhookUrl - Provider webhook URL registered for this trigger when webhook delivery is used
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// IsWebhookRegistered - Whether webhook registration is currently active for this trigger
	IsWebhookRegistered *bool                                               `json:"is_webhook_registered,omitempty"`
	ProviderTrigger     *CallbacksInstancesGetOutputTriggersProviderTrigger `json:"provider_trigger,omitempty"`
}

// CallbacksInstancesGetOutput represents the callbacks instances get output type.
type CallbacksInstancesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback instance identifier
	Id string `json:"id"`
	// Status - Whether the callback instance is currently attached to a deployment/config pair
	Status     string                                 `json:"status"`
	Deployment CallbacksInstancesGetOutputDeployment  `json:"deployment"`
	Config     CallbacksInstancesGetOutputConfig      `json:"config"`
	AuthConfig *CallbacksInstancesGetOutputAuthConfig `json:"auth_config,omitempty"`
	// Triggers - Resolved trigger registrations for this callback instance
	Triggers []CallbacksInstancesGetOutputTriggers `json:"triggers"`
	// CreatedAt - Timestamp when the callback instance was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback instance was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbacksInstancesGetOutputFromJSON deserializes JSON data into a CallbacksInstancesGetOutput.
func MapCallbacksInstancesGetOutputFromJSON(data []byte) (*CallbacksInstancesGetOutput, error) {
	var v CallbacksInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksInstancesGetOutputToJSON serializes a CallbacksInstancesGetOutput to JSON.
func MapCallbacksInstancesGetOutputToJSON(v *CallbacksInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
