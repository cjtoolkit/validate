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

func Must(value string, err error) string {
	vError.Must(err)
	return value
}
