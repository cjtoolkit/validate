package vTime

import (
	"time"

	"github.com/cjtoolkit/validate/vError"
)

/*
Override Error Message
*/
func OverrideErrorMsg(validationError vError.ValidationError, rules ...ValidationRule) ValidationRule {
	return func(format *string, src *string, value *time.Time, hasError bool) error {
		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(format, src, value, hasError || collector.HasError()))
		}

		if collector.HasError() {
			return validationError
		}

		return nil
	}
}

/*
Make sure value is set, if not set the rule return a validation error

Note: will only work while validating from string
*/
func Mandatory() ValidationRule {
	return func(format *string, src *string, value *time.Time, hasError bool) error {
		if *src == "" {
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
Minimum value, returns error if less than min.
*/
func Min(min time.Time) ValidationRule {
	return func(format *string, src *string, value *time.Time, hasError bool) error {
		if (*value).Before(min) {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"min": min.Format(*format),
				},
				Format: MinErrorFormat,
			}
		}

		return nil
	}
}

/*
Maximum value, returns error if more than max.
*/
func Max(max time.Time) ValidationRule {
	return func(format *string, src *string, value *time.Time, hasError bool) error {
		if (*value).After(max) {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"max": max.Format(*format),
				},
				Format: MaxErrorFormat,
			}
		}

		return nil
	}
}

/*
Between Minimum and Maximum value
*/
func Between(min, max time.Time) ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type: Type,
		Data: map[string]interface{}{
			"min": min,
			"max": max,
		},
		Format: BetweenErrorFormat,
	}, Min(min), Max(max))
}
