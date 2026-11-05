// Package nullable provides a value that can be sent as an explicit JSON null.
//
// Request fields that accept null in addition to a value use a pointer to
// Nullable: nil omits the field, Null sends null, and Of sends the value.
package nullable

import "encoding/json"

// Nullable holds either a value or an explicit null.
type Nullable[T any] struct {
	value *T
}

// Of returns a Nullable that is sent as v.
func Of[T any](v T) *Nullable[T] {
	return &Nullable[T]{value: &v}
}

// Null returns a Nullable that is sent as JSON null.
func Null[T any]() *Nullable[T] {
	return &Nullable[T]{}
}

// IsNull reports whether n holds an explicit null.
func (n Nullable[T]) IsNull() bool {
	return n.value == nil
}

// Value returns the held value and whether one is present.
func (n Nullable[T]) Value() (T, bool) {
	if n.value == nil {
		var zero T
		return zero, false
	}
	return *n.value, true
}

// MarshalJSON implements json.Marshaler.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.value = nil
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	n.value = &v
	return nil
}
