package vFloat

import (
	"strconv"
	"strings"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(src *string, value *float64, hasError bool) error

func validate(src string, value float64, rules ...ValidationRule) (float64, error) {
	srcPtr := &src
	valuePtr := &value

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(srcPtr, valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func Validate(value float64, rules ...ValidationRule) (float64, error) {
	return validate("-", value, rules...)
}

/*
Convert to float64 and than validate.
*/
func ValidateFromString(src string, rules ...ValidationRule) (float64, error) {
	src = strings.TrimSpace(src)
	if src == "" {
		return validate(src, 0, rules...)
	}

	value, err := strconv.ParseFloat(src, 64)
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

func Must(value float64, err error) float64 {
	vError.Must(err)
	return value
}
