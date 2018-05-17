package vInt

import "testing"

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
}
