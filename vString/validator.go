package vString

import (
	"strings"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(value *string, hasError bool) error

func Validate(value string, rules ...ValidationRule) (string, error) {
	value = strings.TrimSpace(value)
	valuePtr := &value

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func MustValidate(value string, rules ...ValidationRule) string {
	str, err := Validate(value, rules...)
	vError.Must(err)

	return str
}
