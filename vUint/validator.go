package vUint

import (
	"strconv"
	"strings"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(src *string, value *uint64, hasError bool) error

func validate(src string, value uint64, rules ...ValidationRule) (uint64, error) {
	srcPtr := &src
	valuePtr := &value

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(srcPtr, valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func Validate(value uint64, rules ...ValidationRule) (uint64, error) {
	return validate("-", value, rules...)
}

/*
Convert to uint64 and than validate.
*/
func ValidateFromString(src string, rules ...ValidationRule) (uint64, error) {
	src = strings.TrimSpace(src)
	if src == "" {
		return validate(src, 0, rules...)
	}

	value, err := strconv.ParseUint(src, 10, 64)
	if nil != err {
		err = vError.Errors{vError.ValidationError{
			Type: Type,
			Data: map[string]interface{}{
				"value": value,
			},
			Format: NotAnIntErrorFormat,
		}}
		return value, err
	}

	return validate(src, value, rules...)
}

func Must(value uint64, err error) uint64 {
	vError.Must(err)
	return value
}
