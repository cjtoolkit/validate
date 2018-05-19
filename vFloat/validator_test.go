package vFloat

import "testing"

func TestValidator(t *testing.T) {
	t.Run("ValidateFromString", func(t *testing.T) {
		t.Run("Empty String", func(t *testing.T) {
			if _, err := ValidateFromString(""); err != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Not an float", func(t *testing.T) {
			if _, err := ValidateFromString("A"); err == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("An float", func(t *testing.T) {
			if _, err := ValidateFromString("1.5"); err != nil {
				t.Error("Should be nil")
			}
		})
	})
}
