package typing

import (
	"encoding/json"
)

// ChatsChannelsTypingStartOutput represents the chats channels typing start output type.
type ChatsChannelsTypingStartOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Started - Whether the typing indicator was shown
	Started bool `json:"started"`
}

// MapChatsChannelsTypingStartOutputFromJSON deserializes JSON data into a ChatsChannelsTypingStartOutput.
func MapChatsChannelsTypingStartOutputFromJSON(data []byte) (*ChatsChannelsTypingStartOutput, error) {
	var v ChatsChannelsTypingStartOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsTypingStartOutputToJSON serializes a ChatsChannelsTypingStartOutput to JSON.
func MapChatsChannelsTypingStartOutputToJSON(v *ChatsChannelsTypingStartOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsChannelsTypingStartBody represents the chats channels typing start body type.
type ChatsChannelsTypingStartBody struct {
	// ThreadId - Show the typing indicator within this thread
	ThreadId *string `json:"thread_id,omitempty"`
	// Status - Provider-specific typing status text
	Status *string `json:"status,omitempty"`
}

// MapChatsChannelsTypingStartBodyFromJSON deserializes JSON data into a ChatsChannelsTypingStartBody.
func MapChatsChannelsTypingStartBodyFromJSON(data []byte) (*ChatsChannelsTypingStartBody, error) {
	var v ChatsChannelsTypingStartBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsTypingStartBodyToJSON serializes a ChatsChannelsTypingStartBody to JSON.
func MapChatsChannelsTypingStartBodyToJSON(v *ChatsChannelsTypingStartBody) ([]byte, error) {
	return json.Marshal(v)
}
