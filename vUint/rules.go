package vUint

import "github.com/cjtoolkit/validate/vError"

func Mandatory() ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if "" == *src {
			return vError.ValidationError{
				Type:   Type,
				Data:   nil,
				Format: MandatoryFormat,
			}
		}

		return nil
	}
}

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

func DefaultValue(defaultValue uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if "" == *src {
			*value = defaultValue
		}

		return nil
	}
}

func Min(min uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if *value < min {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"min": min,
				},
				Format: MinFormat,
			}
		}

		return nil
	}
}

func Max(max uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if *value > max {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"max": max,
				},
				Format: MaxFormat,
			}
		}

		return nil
	}
}

func Between(min, max uint64) ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type: Type,
		Data: map[string]interface{}{
			"min": min,
			"max": max,
		},
		Format: BetweenFormat,
	}, Min(min), Max(max))
}

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

func Step(step uint64) ValidationRule {
	return func(src *string, value *uint64, hasError bool) error {
		if (*value % step) != 0 {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"step": step,
				},
				Format: StepFormat,
			}
		}

		return nil
	}
}
