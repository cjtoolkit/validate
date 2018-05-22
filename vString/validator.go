package vString

import (
	"sort"
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

func ValidateMultiple(values []string, rules ...ValidationRule) ([]string, error) {
	m := map[string]bool{}

	collector := vError.NewErrorCollector()

	for _, value := range values {
		value, err := Validate(value, rules...)
		m[value] = true
		collector.Collect(err)
	}

	cleanValues := []string{}
	for key, _ := range m {
		cleanValues = append(cleanValues, key)
	}
	sort.Strings(cleanValues)

	return cleanValues, vError.CleanError(collector.GetErrors())
}

func Must(value string, err error) string {
	vError.Must(err)
	return value
}

func MustMultiple(values []string, err error) []string {
	vError.Must(err)
	return values
}
