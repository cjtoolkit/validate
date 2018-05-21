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
				Type:   Type,
				Data:   nil,
				Format: MandatoryErrorFormat,
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
				Type: Type,
				Data: map[string]interface{}{
					"pattern": pattern.String(),
				},
				Format: PatternErrorFormat,
			}
		}

		return nil
	}
}

func MinRune(min int) ValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) < min {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"min": min,
				},
				Format: MinRuneErrorFormat,
			}
		}

		return nil
	}
}

func MaxRune(max int) ValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) > max {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"max": max,
				},
				Format: MaxRuneErrorFormat,
			}
		}

		return nil
	}
}

func BetweenRune(min, max int) ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type: Type,
		Data: map[string]interface{}{
			"min": min,
			"max": max,
		},
		Format: BetweenRuneErrorFormat,
	}, MinRune(min), MaxRune(max))
}

func MustMatch(mustMatch, fieldName string) ValidationRule {
	return func(value *string, hasError bool) error {
		if *value != mustMatch {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"fieldName": fieldName,
				},
				Format: MustMatchErrorFormat,
			}
		}

		return nil
	}
}

func Matches(matches ...string) ValidationRule {
	return func(value *string, hasError bool) error {
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

func Alpha() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaPattern)))
}

func AlphaDash() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaDashErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaDashPattern)))
}

func AlphaNumeric() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaNumericErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaNumericPattern)))
}

func CreditCard() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: CreditCardErrorFormat,
	}, Pattern(regexp.MustCompile(CreditCardPattern)))
}

func CSSColor() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: CSSColorErrorFormat,
	}, Pattern(regexp.MustCompile(CSSColorPattern)))
}

func Email() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: EmailErrorFormat,
	}, Pattern(regexp.MustCompile(EmailPattern)))
}

func IP() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPErrorFormat,
	}, Pattern(regexp.MustCompile(IPPattern)))
}

func IPV4() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPV4ErrorFormat,
	}, Pattern(regexp.MustCompile(IPV4Pattern)))
}

func IPV6() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPV6ErrorFormat,
	}, Pattern(regexp.MustCompile(IPV6Pattern)))
}

func URL() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: URLErrorFormat,
	}, Pattern(regexp.MustCompile(URLPattern)))
}

func UUID() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUIDErrorFormat,
	}, Pattern(regexp.MustCompile(UUIDPattern)))
}

func UUID3() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID3ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID3Pattern)))
}

func UUID4() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID4ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID4Pattern)))
}

func UUID5() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID5ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID5Pattern)))
}
