package vString

import (
	"regexp"
	"unicode/utf8"

	"sort"

	"github.com/cjtoolkit/validate/vError"
)

/*
Make sure value is set, if not set the rule return a validation error
*/
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

/*
Optional Value, if value is not set, return nil, otherwise go though the validation rules.
*/
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

/*
Validate Pattern
*/
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

/*
Validate minimum number of characters
*/
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

/*
Validate maximum number of characters
*/
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

/*
Validate minimum and maximum number of characters
*/
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

/*
Must match field.
*/
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

/*
Check for matches, return error if matches is not found
*/
func Matches(matches ...string) ValidationRule {
	m := toBoolMap(matches)
	sort.Strings(matches)
	return func(value *string, hasError bool) error {
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

/*
Override Error Message
*/
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

/*
Validate Alpha
*/
func Alpha() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaPattern)))
}

/*
Validate AlphaDash
*/
func AlphaDash() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaDashErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaDashPattern)))
}

/*
Validate AlphaNumeric
*/
func AlphaNumeric() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: AlphaNumericErrorFormat,
	}, Pattern(regexp.MustCompile(AlphaNumericPattern)))
}

/*
Validate CreditCard
*/
func CreditCard() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: CreditCardErrorFormat,
	}, Pattern(regexp.MustCompile(CreditCardPattern)))
}

/*
Validate CSSColor
*/
func CSSColor() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: CSSColorErrorFormat,
	}, Pattern(regexp.MustCompile(CSSColorPattern)))
}

/*
Validate Email
*/
func Email() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: EmailErrorFormat,
	}, Pattern(regexp.MustCompile(EmailPattern)))
}

/*
Validate IP
*/
func IP() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPErrorFormat,
	}, Pattern(regexp.MustCompile(IPPattern)))
}

/*
Validate IPV4
*/
func IPV4() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPV4ErrorFormat,
	}, Pattern(regexp.MustCompile(IPV4Pattern)))
}

/*
Validate IPV6
*/
func IPV6() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: IPV6ErrorFormat,
	}, Pattern(regexp.MustCompile(IPV6Pattern)))
}

/*
Validate URL
*/
func URL() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: URLErrorFormat,
	}, Pattern(regexp.MustCompile(URLPattern)))
}

/*
Validate UUID
*/
func UUID() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUIDErrorFormat,
	}, Pattern(regexp.MustCompile(UUIDPattern)))
}

/*
Validate UUID3
*/
func UUID3() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID3ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID3Pattern)))
}

/*
Validate UUID4
*/
func UUID4() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID4ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID4Pattern)))
}

/*
Validate UUID5
*/
func UUID5() ValidationRule {
	return OverrideErrorMsg(vError.ValidationError{
		Type:   Type,
		Data:   nil,
		Format: UUID5ErrorFormat,
	}, Pattern(regexp.MustCompile(UUID5Pattern)))
}
