package vTime

import (
	"time"

	"github.com/cjtoolkit/validate/vError"
)

func OverrideErrorMsg(validationError vError.ValidationError, rules ...ValidationRule) ValidationRule {
	return func(format *string, src *string, value *time.Time, hasError bool) error {
		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(format, src, value, hasError))
		}

		if collector.HasError() {
			return validationError
		}

		return nil
	}
}

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
