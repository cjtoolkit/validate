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

	t.Run("CleanError", func(t *testing.T) {
		errs := Errors{
			ValidationError{
				Type: "abc",
			},
			ValidationError{
				Type: "def",
			},
			ValidationError{
				Type: "abc",
			},
			ValidationError{
				Type: "def",
			},
		}

		if cap(CleanError(errs).(Errors)) != 2 {
			t.Error("Should be 2")
		}
	})

	t.Run("join", func(t *testing.T) {
		if join("and", []string{}) != "" {
			t.Error("Should be empty")
		}

		if join("and", []string{"Myself"}) != "Myself" {
			t.Error("Should be 'Myself'")
		}

		if join("and", []string{"Me", "Myself", "I"}) != "Me, Myself and I" {
			t.Error("Should be 'Me, Myself and I'")
		}

		if join("and", []int64{-4, -5, -6}) != "-4, -5 and -6" {
			t.Error("Should be '-4, -5 and -6'")
		}

		if join("and", []float64{1.4, 1.5, 1.6}) != "1.4, 1.5 and 1.6" {
			t.Error("Should be '1.4, 1.5 and 1.6'")
		}

		if join("and", []uint64{4, 5, 6}) != "4, 5 and 6" {
			t.Error("Should be '4, 5 and 6'")
		}

		if join("and", []int{4, 5, 6}) != "" {
			t.Error("Should be empty")
		}
	})
}
