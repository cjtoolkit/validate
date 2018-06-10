package vInt

import (
	"github.com/cjtoolkit/validate/vError"
)

/*
Make sure value is set, if not set the rule return a validation error

Note: will only work while validating from string
*/
func Mandatory() ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if "" == *src {
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
Optional Value, if value is not set, return nil, otherwise go though the validation rules.

Note: will only work while validating from string
*/
func Optional(rules ...ValidationRule) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if "" == *src {
			return nil
		}

		collector := vError.NewErrorCollector()

		for _, rule := range rules {
			collector.Collect(rule(src, value, collector.HasError()))
		}

		return collector.GetErrors()
	}
}

/*
If the value is not set, set it to the default value.

Note: will only work while validating from string
*/
func DefaultValue(defaultValue int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if "" == *src {
			*value = defaultValue
		}

		return nil
	}
}

/*
Minimum value, returns error if less than min.
*/
func Min(min int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if *value < min {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"min": min,
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
func Max(max int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if *value > max {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"max": max,
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
func Between(min, max int64) ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type: Type,
		Data: map[string]interface{}{
			"min": min,
			"max": max,
		},
		Format: BetweenErrorFormat,
	}, Min(min), Max(max))
}

/*
Override Error Message
*/
func OverrideErrorMsg(validationError vError.ValidationError, rules ...ValidationRule) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(src, value, hasError))
		}

		if collector.HasError() {
			return validationError
		}

		return nil
	}
}

/*
Number of step, if value not within number of steps returns an error.
*/
func Step(step int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if (*value % step) != 0 {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"step": step,
				},
				Format: StepErrorFormat,
			}
		}

		return nil
	}
}

/*
Check for matches, return error if matches is not found
*/
func Matches(matches ...int64) ValidationRule {
	m := toBoolMap(matches)
	matches = nil
	return func(src *string, value *int64, hasError bool) error {
		if m[*value] {
			return nil
		}

		return vError.ValidationError{
			Type: Type,
			Data: map[string]interface{}{
				"matches": matches,
			},
			Format: MatchesErrorFormat,
		}
	}
}
