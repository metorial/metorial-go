package deployments

import (
	"encoding/json"
	"time"
)

// CustomProvidersDeploymentsGetLogsOutputStepsLogs represents the custom providers deployments get logs output steps logs type.
type CustomProvidersDeploymentsGetLogsOutputStepsLogs struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Timestamp - When the log line was emitted.
	Timestamp time.Time `json:"timestamp"`
	// Message - Log line text.
	Message string `json:"message"`
}

// CustomProvidersDeploymentsGetLogsOutputSteps - Step of the deployment.
type CustomProvidersDeploymentsGetLogsOutputSteps struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Name of the step.
	Name string `json:"name"`
	// Type - What kind of work the step does.
	Type string `json:"type"`
	// Status - Step execution status.
	Status string `json:"status"`
	// Logs - Log lines for this step.
	Logs []CustomProvidersDeploymentsGetLogsOutputStepsLogs `json:"logs"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// StartedAt - When the step started. `null` while `status` is `pending`.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// EndedAt - When the step ended. `null` until the step finishes.
	EndedAt *time.Time `json:"ended_at,omitempty"`
}

// CustomProvidersDeploymentsGetLogsOutput represents the custom providers deployments get logs output type.
type CustomProvidersDeploymentsGetLogsOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// CustomProviderDeploymentId - Deployment these logs belong to.
	CustomProviderDeploymentId string `json:"custom_provider_deployment_id"`
	// Steps - Steps of the deployment, in order.
	Steps []CustomProvidersDeploymentsGetLogsOutputSteps `json:"steps"`
}

// MapCustomProvidersDeploymentsGetLogsOutputFromJSON deserializes JSON data into a CustomProvidersDeploymentsGetLogsOutput.
func MapCustomProvidersDeploymentsGetLogsOutputFromJSON(data []byte) (*CustomProvidersDeploymentsGetLogsOutput, error) {
	var v CustomProvidersDeploymentsGetLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersDeploymentsGetLogsOutputToJSON serializes a CustomProvidersDeploymentsGetLogsOutput to JSON.
func MapCustomProvidersDeploymentsGetLogsOutputToJSON(v *CustomProvidersDeploymentsGetLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}
