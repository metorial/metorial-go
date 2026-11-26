package items

import (
	"encoding/json"
	"time"
)

// StoresItemsModifyOutputItemsFileCreatedByOrganizationActorMember represents the stores items modify output items file created by organization actor member type.
type StoresItemsModifyOutputItemsFileCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams struct {
	// Id - The team ID
	Id string `json:"id"`
	// Name - The team name
	Name string `json:"name"`
	// Slug - The team slug
	Slug string `json:"slug"`
	// AssignmentId - The team assignment ID
	AssignmentId string `json:"assignment_id"`
	// CreatedAt - The team assignment creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The team assignment last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsFileCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresItemsModifyOutputItemsFileCreatedByOrganizationActor struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Type - The organization member's type
	Type string `json:"type"`
	// OrganizationId - The organization member's organization ID
	OrganizationId string `json:"organization_id"`
	// Name - The organization member's name
	Name string `json:"name"`
	// Email - The organization member's email
	Email *string `json:"email,omitempty"`
	// ImageUrl - The organization member's image URL
	ImageUrl string                                                            `json:"image_url"`
	Member   *StoresItemsModifyOutputItemsFileCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresItemsModifyOutputItemsFileCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsFileCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresItemsModifyOutputItemsFileCreatedByConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsFileCreatedBy - Actor that uploaded the file.
type StoresItemsModifyOutputItemsFileCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresItemsModifyOutputItemsFileCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresItemsModifyOutputItemsFileCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresItemsModifyOutputItemsFile - File at this item. `null` unless `kind` is `file`.
type StoresItemsModifyOutputItemsFile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` after the file is deleted.
	Status string `json:"status"`
	// FileName - Name of the uploaded file.
	FileName string `json:"file_name"`
	// FileSize - Size of the file in bytes.
	FileSize float64 `json:"file_size"`
	// FileType - MIME type of the file.
	FileType string `json:"file_type"`
	// Title - Label shown for the file.
	Title string `json:"title"`
	// Purpose - Why the file was uploaded. Controls what it can be used for.
	Purpose string `json:"purpose"`
	// CreatedBy - Actor that uploaded the file.
	CreatedBy *StoresItemsModifyOutputItemsFileCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorMember represents the stores items modify output items document created by organization actor member type.
type StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams - The teams the actor belongs to
type StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams struct {
	// Id - The team ID
	Id string `json:"id"`
	// Name - The team name
	Name string `json:"name"`
	// Slug - The team slug
	Slug string `json:"slug"`
	// AssignmentId - The team assignment ID
	AssignmentId string `json:"assignment_id"`
	// CreatedAt - The team assignment creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The team assignment last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Type - The organization member's type
	Type string `json:"type"`
	// OrganizationId - The organization member's organization ID
	OrganizationId string `json:"organization_id"`
	// Name - The organization member's name
	Name string `json:"name"`
	// Email - The organization member's email
	Email *string `json:"email,omitempty"`
	// ImageUrl - The organization member's image URL
	ImageUrl string                                                                `json:"image_url"`
	Member   *StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorMember `json:"member,omitempty"`
	Teams    []StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedByConsumer - Consumer details. `null` unless `type` is `consumer`.
type StoresItemsModifyOutputItemsDocumentCreatedByConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItemsDocumentCreatedBy - Actor that created the document.
type StoresItemsModifyOutputItemsDocumentCreatedBy struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *StoresItemsModifyOutputItemsDocumentCreatedByOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *StoresItemsModifyOutputItemsDocumentCreatedByConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// StoresItemsModifyOutputItemsDocument - Document at this item. `null` unless `kind` is `document`.
type StoresItemsModifyOutputItemsDocument struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` after the document is deleted.
	Status string `json:"status"`
	// Title - Label shown for the document.
	Title string `json:"title"`
	// Content - Text body of the document.
	Content string `json:"content"`
	// FileId - File holding the document content.
	FileId string `json:"file_id"`
	// ParentDocumentId - Document this was cloned from. `null` for top-level documents.
	ParentDocumentId *string `json:"parent_document_id,omitempty"`
	// CurrentVersionId - Latest version of the document, by ID.
	CurrentVersionId *string `json:"current_version_id,omitempty"`
	// CreatedBy - Actor that created the document.
	CreatedBy *StoresItemsModifyOutputItemsDocumentCreatedBy `json:"created_by,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutputItems represents the stores items modify output items type.
type StoresItemsModifyOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Kind - What the item holds: `file`, `document`, or `directory`.
	Kind string `json:"kind"`
	// Path - Path of the item within the store.
	Path string `json:"path"`
	// StoreId - Store the item belongs to.
	StoreId string `json:"store_id"`
	// DirectoryId - Parent directory, by ID. `null` for top-level items.
	DirectoryId *string `json:"directory_id,omitempty"`
	// File - File at this item. `null` unless `kind` is `file`.
	File *StoresItemsModifyOutputItemsFile `json:"file,omitempty"`
	// Document - Document at this item. `null` unless `kind` is `document`.
	Document *StoresItemsModifyOutputItemsDocument `json:"document,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresItemsModifyOutput represents the stores items modify output type.
type StoresItemsModifyOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Items - Resulting store items.
	Items []StoresItemsModifyOutputItems `json:"items"`
}

// MapStoresItemsModifyOutputFromJSON deserializes JSON data into a StoresItemsModifyOutput.
func MapStoresItemsModifyOutputFromJSON(data []byte) (*StoresItemsModifyOutput, error) {
	var v StoresItemsModifyOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsModifyOutputToJSON serializes a StoresItemsModifyOutput to JSON.
func MapStoresItemsModifyOutputToJSON(v *StoresItemsModifyOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresItemsModifyBodyOperations represents the stores items modify body operations type.
type StoresItemsModifyBodyOperations struct {
	// Type - Operation to perform.
	Type *string `json:"type,omitempty"`
	// ItemId - Store item the operation applies to, by ID.
	ItemId *string `json:"itemId,omitempty"`
	// FileId - File to add to the store, by ID.
	FileId *string `json:"fileId,omitempty"`
	// DocumentId - Document to add to the store, by ID.
	DocumentId *string `json:"documentId,omitempty"`
	// Path - Path the operation applies to.
	Path *string `json:"path,omitempty"`
}

// StoresItemsModifyBody represents the stores items modify body type.
type StoresItemsModifyBody struct {
	// Operations - Operations to apply to the store.
	Operations []StoresItemsModifyBodyOperations `json:"operations"`
}

// MapStoresItemsModifyBodyFromJSON deserializes JSON data into a StoresItemsModifyBody.
func MapStoresItemsModifyBodyFromJSON(data []byte) (*StoresItemsModifyBody, error) {
	var v StoresItemsModifyBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresItemsModifyBodyToJSON serializes a StoresItemsModifyBody to JSON.
func MapStoresItemsModifyBodyToJSON(v *StoresItemsModifyBody) ([]byte, error) {
	return json.Marshal(v)
}
