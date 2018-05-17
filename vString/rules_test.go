package vString

import (
	"errors"
	"regexp"
	"testing"

	"github.com/cjtoolkit/validate/vError"
)

func TestRules(t *testing.T) {
	values := func(value string, hasError bool) (*string, bool) { return &value, hasError }

	t.Run("Mandatory", func(t *testing.T) {
		t.Run("Value is empty", func(t *testing.T) {
			if Mandatory()(values("", false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Mandatory()(values("-", false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Optional", func(t *testing.T) {
		rule := ValidationRule(func(value *string, hasError bool) error {
			return errors.New("I am error")
		})

		t.Run("Value is empty", func(t *testing.T) {
			if Optional(rule)(values("", false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Optional(rule)(values("-", false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("Pattern", func(t *testing.T) {
		pattern := regexp.MustCompile("([0-9]+)")

		t.Run("It's matches", func(t *testing.T) {
			if Pattern(pattern)(values("123", false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("It's does not match", func(t *testing.T) {
			if Pattern(pattern)(values("abc", false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("MinRune", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if MinRune(5)(values("abc", false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if MinRune(5)(values("abcdef", false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("MaxRune", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if MaxRune(5)(values("abc", false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if MaxRune(5)(values("abcdef", false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("MustMatch", func(t *testing.T) {
		t.Run("Match", func(t *testing.T) {
			if MustMatch("abc", "name")(values("abc", false)) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("No Match", func(t *testing.T) {
			if MustMatch("abc", "name")(values("def", false)) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("OverrideErrorMsg", func(t *testing.T) {
		ruleWithError := ValidationRule(func(value *string, hasError bool) error {
			return errors.New("I am error")
		})
		ruleWithNoError := ValidationRule(func(value *string, hasError bool) error {
			return nil
		})

		t.Run("Has Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithError)(values("", false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("No Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithNoError)(values("", false)) != nil {
				t.Error("Should be nil")
			}
		})
	})
}
