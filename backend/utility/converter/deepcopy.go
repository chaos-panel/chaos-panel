package converter

import (
	"encoding/json"
	"fmt"
)

// DeepCopyInto deep copy value into ptr. ptr must be a pointer.
func deepCopyInto[T any](ptr *T, value any) error {
	if ptr == nil {
		return fmt.Errorf("ptr is nil")
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonValue, ptr); err != nil {
		return err
	}
	return nil
}

func DeepCopy[T any](value any) (T, error) {
	var ptr T
	err := deepCopyInto(&ptr, value)
	return ptr, err
}

func ToAny[T any](value any) (T, error) {
	var ptr T
	err := deepCopyInto(&ptr, value)
	return ptr, err
}
