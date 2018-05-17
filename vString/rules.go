package vString

import (
	"regexp"
	"unicode/utf8"

	"github.com/cjtoolkit/validate/vError"
)

func Mandatory() ValidationRule {
	return func(value *string, hasError bool) error {
		if "" == *value {
			return vError.ValidationError{
				Data:   nil,
				Format: MandatoryFormat,
			}
		}

		return nil
	}
}

func Optional(rules ...ValidationRule) ValidationRule {
	return func(value *string, hasError bool) error {
		if "" == *value {
			return nil
		}

		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(value, collector.HasError()))
		}

		return collector.GetErrors()
	}
}

func Pattern(pattern *regexp.Regexp) ValidationRule {
	return func(value *string, hasError bool) error {
		if !pattern.MatchString(*value) {
			return vError.ValidationError{
				Data: map[string]interface{}{
					"pattern": pattern.String(),
				},
				Format: PatternFormat,
			}
		}

		return nil
	}
}

func MinRune(min int) ValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) < min {
			return vError.ValidationError{
				Data: map[string]interface{}{
					"min": min,
				},
				Format: MinRuneFormat,
			}
		}

		return nil
	}
}

func MaxRune(max int) ValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) > max {
			return vError.ValidationError{
				Data: map[string]interface{}{
					"max": max,
				},
				Format: MaxRuneFormat,
			}
		}

		return nil
	}
}

func MustMatch(mustMatch, fieldName string) ValidationRule {
	return func(value *string, hasError bool) error {
		if *value != mustMatch {
			return vError.ValidationError{
				Data: map[string]interface{}{
					"fieldName": fieldName,
				},
				Format: MustMatchFormat,
			}
		}

		return nil
	}
}

func OverrideErrorMsg(validationError vError.ValidationError, rules ...ValidationRule) ValidationRule {
	return func(value *string, hasError bool) error {
		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(value, hasError))
		}

		if collector.HasError() {
			return validationError
		}

		return nil
	}
}

func OverrideValue(overrideValue string) ValidationRule {
	return func(value *string, hasError bool) error {
		*value = overrideValue

		return nil
	}
}
