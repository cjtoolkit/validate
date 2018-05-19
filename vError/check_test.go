package vError

import (
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {
	t.Run("CheckErr", func(t *testing.T) {
		t.Run("has no error", func(t *testing.T) {
			if CheckErr(nil, nil, nil) != true {
				t.Error("Should be true")
			}
		})

		t.Run("has error", func(t *testing.T) {
			if CheckErr(nil, errors.New("I am error"), nil) != false {
				t.Error("Should be false")
			}
		})
	})

	t.Run("CheckBool", func(t *testing.T) {
		t.Run("all true", func(t *testing.T) {
			if CheckBool(true, true, true) != true {
				t.Error("Should be true")
			}
		})

		t.Run("one is false", func(t *testing.T) {
			if CheckBool(true, false, true) != false {
				t.Error("Should be false")
			}
		})
	})

	t.Run("Must", func(t *testing.T) {
		t.Run("No error, therefore does not panic", func(t *testing.T) {
			Must(nil)
		})

		t.Run("Error, therefore panic", func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("Recover should not be nil")
				}
			}()

			Must(errors.New("I am error"))
		})
	})
}