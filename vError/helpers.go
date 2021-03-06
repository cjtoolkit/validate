package vError

import (
	"fmt"
	"strings"
)

/*
Check for errors, return true if there is no error, otherwise false
*/
func CheckErr(errs ...error) bool {
	for _, err := range errs {
		if nil != err {
			return false
		}
	}

	return true
}

/*
Check bools, return true if all is true, otherwise false
*/
func CheckBool(boolValues ...bool) bool {
	for _, boolValue := range boolValues {
		if !boolValue {
			return false
		}
	}

	return true
}

/*
Panic if error is not nil.
*/
func Must(err error) {
	if nil != err {
		panic(err)
	}
}

/*
Clean out duplicate error message.
*/
func CleanError(err error) error {
	switch err := err.(type) {
	case Errors:
		dupeList := map[string]bool{}
		cleanErr := Errors{}

		for _, e := range err {
			if ve, ok := e.(ValidationError); ok {
				key := ve.Type + ve.Format
				if !dupeList[key] {
					dupeList[key] = true
					cleanErr = append(cleanErr, ve)
				}
			}
		}

		return cleanErr
	}

	return err
}

/*
Merge Multiple errors into one.
*/
func MergeErrors(errs ...error) error {
	collector := NewErrorCollector()
	for _, err := range errs {
		collector.Collect(err)
	}

	return CleanError(collector.GetErrors())
}

func join(article string, values interface{}) string {
	strs := []string{}

	switch values := values.(type) {
	case []string:
		strs = values
	case []int64:
		for _, value := range values {
			strs = append(strs, fmt.Sprint(value))
		}
	case []float64:
		for _, value := range values {
			strs = append(strs, fmt.Sprint(value))
		}
	case []uint64:
		for _, value := range values {
			strs = append(strs, fmt.Sprint(value))
		}
	}

	strsLen := len(strs)
	switch strsLen {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	return strings.Join(strs[:strsLen-1], ", ") + " " + article + " " + strs[strsLen-1]
}
