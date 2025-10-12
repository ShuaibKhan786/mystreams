package utils

import "fmt"

func SanitizeNilPointer[T any](model *T) T {
	if model != nil {
		return *model
	}
	var zero T
	return zero
}

func SanitizeNilPointerUI[T any](model *T) string {
	if model != nil {
		return fmt.Sprint(*model)
	}

	return ""
}
