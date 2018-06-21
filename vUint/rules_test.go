package vUint

import (
	"errors"
	"testing"

	"github.com/cjtoolkit/validate/vError"
)

func TestRules(t *testing.T) {
	values := func(src string, value uint64, hasError bool) (*string, *uint64, bool) {
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
		rule := ValidationRule(func(src *string, value *uint64, hasError bool) error {
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

	t.Run("OverrideErrorMsg", func(t *testing.T) {
		ruleWithError := ValidationRule(func(src *string, value *uint64, hasError bool) error {
			return errors.New("I am error")
		})
		ruleWithNoError := ValidationRule(func(src *string, value *uint64, hasError bool) error {
			return nil
		})

		t.Run("Has Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithError)(values("", 0, false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("No Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithNoError)(values("", 0, false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Step", func(t *testing.T) {
		t.Run("In Step", func(t *testing.T) {
			if Step(2)(values("", 4, false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Out of Step", func(t *testing.T) {
			if Step(2)(values("", 5, false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("Matches", func(t *testing.T) {
		t.Run("Match", func(t *testing.T) {
			if Matches(5)(values("", 5, false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("No Match", func(t *testing.T) {
			if Matches(5)(values("", 6, false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})
}
