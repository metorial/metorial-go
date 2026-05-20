package permissions

import (
	"encoding/json"
)

// StoresPermissionsGetOutput represents the stores permissions get output type.
type StoresPermissionsGetOutput struct {
	// Object - String representing the object's type
	Object           string   `json:"object"`
	StoreId          string   `json:"store_id"`
	HasFullAccess    bool     `json:"has_full_access"`
	Permissions      []string `json:"permissions"`
	RelevantStoreIds []string `json:"relevant_store_ids"`
	ReadableStoreIds []string `json:"readable_store_ids"`
	WritableStoreIds []string `json:"writable_store_ids"`
}

// MapStoresPermissionsGetOutputFromJSON deserializes JSON data into a StoresPermissionsGetOutput.
func MapStoresPermissionsGetOutputFromJSON(data []byte) (*StoresPermissionsGetOutput, error) {
	var v StoresPermissionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresPermissionsGetOutputToJSON serializes a StoresPermissionsGetOutput to JSON.
func MapStoresPermissionsGetOutputToJSON(v *StoresPermissionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
