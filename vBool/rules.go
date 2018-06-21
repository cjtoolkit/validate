package vBool

import "github.com/cjtoolkit/validate/vError"

/*
Make sure value is true. Leave this rule to last.
*/
func Mandatory() ValidationRule {
	return func(src *string, value *bool, hasError bool) error {
		if false == *value {
			return vError.ValidationError{
				Type:   Type,
				Data:   nil,
				Format: MandatoryErrorFormat,
			}
		}

		return nil
	}
}

/*
If the src matches, set value to true.
*/
func TrueIfMatches(match string) ValidationRule {
	return func(src *string, value *bool, hasError bool) error {
		if match == *src {
			*value = true
		}

		return nil
	}
}
