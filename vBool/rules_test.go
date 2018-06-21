package vBool

import "testing"

func TestRules(t *testing.T) {
	t.Run("Mandatory", func(t *testing.T) {
		t.Run("No Error", func(t *testing.T) {
			src := ""
			value := true

			if Mandatory()(&src, &value, false) != nil {
				t.Error("Should be nil.")
			}
		})

		t.Run("Has Error", func(t *testing.T) {
			src := ""
			value := false

			if Mandatory()(&src, &value, false) == nil {
				t.Error("Should not be nil.")
			}
		})
	})

	t.Run("TrueIfMatches", func(t *testing.T) {
		t.Run("Matches", func(t *testing.T) {
			src := "match"
			value := false

			TrueIfMatches("match")(&src, &value, false)
			if value != true {
				t.Error("Should be true")
			}
		})

		t.Run("No Matches", func(t *testing.T) {
			src := "match"
			value := false

			TrueIfMatches("no-match")(&src, &value, false)
			if value != false {
				t.Error("Should be false")
			}
		})
	})
}
