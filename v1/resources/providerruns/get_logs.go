package providerruns

import (
	"encoding/json"
	"time"
)

// ProviderRunsGetLogsOutputLogs represents the provider runs get logs output logs type.
type ProviderRunsGetLogsOutputLogs struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Timestamp - When the log entry was written.
	Timestamp time.Time `json:"timestamp"`
	// Message - Log message content.
	Message string `json:"message"`
	// OutputType - Which stream the log entry came from.
	OutputType string `json:"output_type"`
}

// ProviderRunsGetLogsOutput represents the provider runs get logs output type.
type ProviderRunsGetLogsOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ProviderRunId - Provider run the logs belong to.
	ProviderRunId string `json:"provider_run_id"`
	// Logs - Log entries for the run.
	Logs []ProviderRunsGetLogsOutputLogs `json:"logs"`
}

// MapProviderRunsGetLogsOutputFromJSON deserializes JSON data into a ProviderRunsGetLogsOutput.
func MapProviderRunsGetLogsOutputFromJSON(data []byte) (*ProviderRunsGetLogsOutput, error) {
	var v ProviderRunsGetLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderRunsGetLogsOutputToJSON serializes a ProviderRunsGetLogsOutput to JSON.
func MapProviderRunsGetLogsOutputToJSON(v *ProviderRunsGetLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}
