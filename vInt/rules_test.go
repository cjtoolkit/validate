package vInt

import (
	"errors"
	"testing"
)

func TestRules(t *testing.T) {
	values := func(src string, value int64, hasError bool) (*string, *int64, bool) {
		return &src, &value, hasError
	}

	t.Run("Mandatory", func(t *testing.T) {
		t.Run("Value is empty", func(t *testing.T) {
			if Mandatory()(values("", 0, false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Mandatory()(values("-", 0, false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Optional", func(t *testing.T) {
		rule := ValidationRule(func(src *string, value *int64, hasError bool) error {
			return errors.New("I am error")
		})

		t.Run("Value is empty", func(t *testing.T) {
			if Optional(rule)(values("", 0, false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Optional(rule)(values("-", 0, false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("DefaultValue", func(t *testing.T) {
		t.Run("Value is empty", func(t *testing.T) {
			src, value, hasError := values("", 0, false)
			DefaultValue(5)(src, value, hasError)

			if *value != 5 {
				t.Error("Should be '5'")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			src, value, hasError := values("0", 0, false)
			DefaultValue(5)(src, value, hasError)

			if *value != 0 {
				t.Error("Should be '0'")
			}
		})
	})

	t.Run("Min", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if Min(5)(values("", 4, false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if Min(5)(values("", 6, false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Max", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if Max(5)(values("", 4, false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if Max(5)(values("", 6, false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})
}
