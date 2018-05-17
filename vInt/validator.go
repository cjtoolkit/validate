package vInt

import (
	"strconv"
	"strings"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(src *string, value *int64, hasError bool) error

func validate(src string, value int64, rules ...ValidationRule) (int64, error) {
	srcPtr := &src
	valuePtr := &value

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(srcPtr, valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func Validate(value int64, rules ...ValidationRule) (int64, error) {
	return validate("-", value, rules...)
}

func ValidateFromString(src string, rules ...ValidationRule) (int64, error) {
	src = strings.TrimSpace(src)
	if src == "" {
		return validate(src, 0, rules...)
	}

	value, err := strconv.ParseInt(src, 10, 64)
	if nil != err {
		err = vError.ValidationError{
			Data: map[string]interface{}{
				"value": value,
			},
			Format: NotAnIntFormat,
		}
		return value, err
	}

	return validate(src, value, rules...)
}

func Must(value int64, err error) int64 {
	vError.Must(err)
	return value
}
