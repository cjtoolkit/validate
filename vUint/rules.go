package vUint

import "github.com/cjtoolkit/validate/vError"

/*
Make sure value is set, if not set the rule return a validation error

Note: will only work while validating from string
*/
func Mandatory() ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
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
	return func(src *string, value *uint64, hasError bool) error {
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
func DefaultValue(defaultValue uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if "" == *src {
			*value = defaultValue
		}

		return nil
	}
}

/*
Minimum value, returns error if less than min.
*/
func Min(min uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
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
func Max(max uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
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
func Between(min, max uint64) ValidationRule {
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
	return func(src *string, value *uint64, hasError bool) error {
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
func Step(step uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
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
