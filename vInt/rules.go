package vInt

import "github.com/cjtoolkit/validate/vError"

func Mandatory() ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if "" == *src {
			return vError.ValidationError{
				Data:   nil,
				Format: MandatoryFormat,
			}
		}

		return nil
	}
}
