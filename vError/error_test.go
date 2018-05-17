package vError

import (
	"errors"
	"testing"
)

func TestErrorCollector(t *testing.T) {
	t.Run("Collect", func(t *testing.T) {
		t.Run("No error", func(t *testing.T) {
			subject := &ErrorCollector{
				hasError: false,
				errs:     []error{},
			}

			subject.Collect(nil)

			if subject.hasError != false {
				t.Error("Should be false")
			}

			if len(subject.errs) != 0 {
				t.Error("Should be 0")
			}
		})

		t.Run("Has error", func(t *testing.T) {
			subject := &ErrorCollector{
				hasError: false,
				errs:     []error{},
			}

			subject.Collect(errors.New("I am error"))

			if subject.hasError != true {
				t.Error("Should be false")
			}

			if len(subject.errs) != 1 {
				t.Error("Should be 1")
			}
		})

		t.Run("Has collection of errors", func(t *testing.T) {
			subject := &ErrorCollector{
				hasError: false,
				errs:     []error{},
			}

			subject.Collect(Errors{
				errors.New("I am error"),
				errors.New("I am error"),
				errors.New("I am error"),
			})

			if subject.hasError != true {
				t.Error("Should be false")
			}

			if len(subject.errs) != 3 {
				t.Error("Should be 3")
			}
		})
	})

	t.Run("GetErrors", func(t *testing.T) {
		t.Run("No error", func(t *testing.T) {
			subject := &ErrorCollector{
				hasError: false,
				errs:     []error{},
			}

			if subject.GetErrors() != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Has error", func(t *testing.T) {
			subject := &ErrorCollector{
				hasError: true,
				errs:     []error{},
			}

			if subject.GetErrors() == nil {
				t.Error("Should not be nil")
			}
		})
	})
}
