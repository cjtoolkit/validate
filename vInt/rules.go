package vInt

import "github.com/cjtoolkit/validate/vError"

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

func DefaultValue(defaultValue int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		if "" == *src {
			*value = defaultValue
		}

		return nil
	}
}

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

func Matches(matches ...int64) ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
		for _, match := range matches {
			if *value == match {
				return nil
			}
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
