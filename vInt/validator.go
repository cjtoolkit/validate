package vInt

import (
	"sort"
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

func ValidateMultiple(values []int64, rules ...ValidationRule) ([]int64, error) {
	collector := vError.NewErrorCollector()

	for _, value := range values {
		_, err := Validate(value, rules...)
		collector.Collect(err)
	}

	return values, vError.CleanError(collector.GetErrors())
}

/*
Convert to int64 and than validate.
*/
func ValidateFromString(src string, rules ...ValidationRule) (int64, error) {
	src = strings.TrimSpace(src)
	if src == "" {
		return validate(src, 0, rules...)
	}

	value, err := strconv.ParseInt(src, 10, 64)
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

/*
Convert to int64 and than validate. (Multiple)
*/
func ValidateFromStringMultiple(srcs []string, rules ...ValidationRule) ([]int64, error) {
	collector := vError.NewErrorCollector()
	m := map[int64]bool{}

	for _, src := range srcs {
		value, err := ValidateFromString(src, rules...)
		m[value] = true
		collector.Collect(err)
	}

	cleanValues := sortInt64{}
	for key, _ := range m {
		cleanValues = append(cleanValues, key)
	}
	sort.Sort(cleanValues)

	return []int64(cleanValues), vError.CleanError(collector.GetErrors())
}

func Must(value int64, err error) int64 {
	vError.Must(err)
	return value
}

func MustMultiple(values []int64, err error) []int64 {
	vError.Must(err)
	return values
}
