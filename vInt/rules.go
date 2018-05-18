package vInt

import "github.com/cjtoolkit/validate/vError"

func Mandatory() ValidationRule {
	return func(src *string, value *int64, hasError bool) error {
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
				Format: MinFormat,
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
				Format: MaxFormat,
			}
		}

		return nil
	}
}
