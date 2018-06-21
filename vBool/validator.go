package vBool

import "github.com/cjtoolkit/validate/vError"

type ValidationRule func(src *string, value *bool, hasError bool) error

func validate(src string, value bool, rules ...ValidationRule) (bool, error) {
	srcPtr := &src
	valuePtr := &value

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(srcPtr, valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func Validate(value bool, rules ...ValidationRule) (bool, error) {
	return validate("-", value, rules...)
}

func ValidateFromString(src string, rules ...ValidationRule) (bool, error) {
	return validate(src, false, rules...)
}

func Must(value bool, err error) bool {
	vError.Must(err)
	return value
}
