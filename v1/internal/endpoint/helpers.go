package endpoint

import "encoding/json"

// StructToQuery converts a struct with json tags to a query parameter map.
// It uses json marshaling to respect json tags and omitempty.
func StructToQuery(v any) map[string]any {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	return m
}

// StructToBody converts a struct to a request body map.
func StructToBody(v any) map[string]any {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	return m
}
