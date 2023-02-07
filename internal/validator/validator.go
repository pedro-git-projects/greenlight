package validator

import (
	"reflect"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$/")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// IsValid returns true if the Errors map is empty
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

// AddError will add an error message to a key if it is not in use
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to a map iff the validation is not ok
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In checks if a given value is contained in a list
func In[T any](value T, list ...T) bool {
	for i := range list {
		if reflect.DeepEqual(value, list[i]) {
			return true
		}
	}
	return false
}

// UniqueKeys returns true iff all keys in a slice are unique
func UniqueKeys[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, v := range values {
		uniqueValues[v] = true
	}

	return len(values) == len(uniqueValues)
}
