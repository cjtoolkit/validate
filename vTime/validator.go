package vTime

import (
	"strings"
	"time"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(format *string, src *string, value *time.Time, hasError bool) error

func validate(format string, src string, value time.Time, rules ...ValidationRule) (time.Time, error) {
	formatPtr := &format
	srcPtr := &src
	valuePtr := &value

	collector := vError.NewErrorCollector()

	time.Time{}.String()

	for _, rule := range rules {
		collector.Collect(rule(formatPtr, srcPtr, valuePtr, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func Validate(value time.Time, rules ...ValidationRule) (time.Time, error) {
	return validate("2006-01-02T15:04:05", "-", value, rules...)
}

func validateFromString(src string, location *time.Location, formats []string, rules ...ValidationRule) (time.Time, error) {
	src = strings.TrimSpace(src)

	for _, format := range formats {
		var (
			value time.Time
			err   error
		)
		if location != nil {
			value, err = time.ParseInLocation(format, src, location)
		} else {
			value, err = time.Parse(format, src)
		}
		if err == nil {
			return validate(format, src, value, rules...)
		}
	}

	return time.Time{}, vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: ParseFormat,
	}
}

func ValidateFromString(src string, location *time.Location, rules ...ValidationRule) (time.Time, error) {
	return validateFromString(
		src,
		location,
		[]string{
			"2006-01-02T15:04:05",
			"2006-01-02T15:04",
			"2006-01-02",
			"15:04:05",
			"15:04",
		},
		rules...,
	)
}

func ValidateFromStringDateOnly(src string, location *time.Location, rules ...ValidationRule) (time.Time, error) {
	return validateFromString(
		src,
		location,
		[]string{
			"2006-01-02",
		},
		rules...,
	)
}

func ValidateFromStringTimeOnly(src string, location *time.Location, rules ...ValidationRule) (time.Time, error) {
	return validateFromString(
		src,
		location,
		[]string{
			"15:04:05",
			"15:04",
		},
		rules...,
	)
}

func Must(value time.Time, err error) time.Time {
	vError.Must(err)
	return value
}
