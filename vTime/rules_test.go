package vTime

import (
	"errors"
	"testing"
	"time"

	"github.com/cjtoolkit/validate/vError"
)

func TestRules(t *testing.T) {
	values := func(format, src string, value time.Time, hasError bool) (*string, *string, *time.Time, bool) {
		return &format, &src, &value, hasError
	}

	t.Run("OverrideErrorMsg", func(t *testing.T) {
		ruleWithError := ValidationRule(func(format *string, src *string, value *time.Time, hasError bool) error {
			return errors.New("I am error")
		})
		ruleWithNoError := ValidationRule(func(format *string, src *string, value *time.Time, hasError bool) error {
			return nil
		})

		t.Run("Has Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithError)(values("", "", time.Time{}, false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("No Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithNoError)(values("", "", time.Time{}, false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Mandatory", func(t *testing.T) {
		t.Run("Value is empty", func(t *testing.T) {
			if Mandatory()(values("", "", time.Time{}, false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Mandatory()(values("-", "-", time.Time{}, false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Min", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if Min(time.Now())(values("", "", time.Now().Add(-(2*time.Second)), false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if Min(time.Now())(values("", "", time.Now().Add(2*time.Second), false)) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("Max", func(t *testing.T) {
		t.Run("Less", func(t *testing.T) {
			if Max(time.Now())(values("", "", time.Now().Add(2*time.Second), false)) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("More", func(t *testing.T) {
			if Max(time.Now())(values("", "", time.Now().Add(-(2*time.Second)), false)) != nil {
				t.Error("Should be nil")
			}
		})
	})
}
