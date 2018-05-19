package vTime

import "testing"

func TestValidator(t *testing.T) {
	t.Run("ValidateFromString", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			if _, err := ValidateFromStringTimeOnly("15:04:05", nil); err != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Invalid", func(t *testing.T) {
			if _, err := ValidateFromStringTimeOnly("AA:BB:CC", nil); err == nil {
				t.Error("Should not be nil")
			}
		})
	})
}
