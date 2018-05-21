package vString

const (
	Type = "string"

	MandatoryErrorFormat    = "Must not be left empty."
	PatternErrorFormat      = "Does not match pattern '{{.pattern}}'."
	MinRuneErrorFormat      = "Must have more than '{{.min}}' characters."
	MaxRuneErrorFormat      = "Must have less than '{{.max}}' characters."
	BetweenRuneErrorFormat  = "Must be between {{.min}} and {{.max}} characters."
	MustMatchErrorFormat    = "Field does not match '{{.fieldName}}'."
	MatchesErrorFormat      = "Field does not match '{{.matches}}'."
	AlphaErrorFormat        = "Should only be alpha characters."
	AlphaDashErrorFormat    = "Should only be alpha characters with underscore and dash."
	AlphaNumericErrorFormat = "Should only be alpha numeric characters."
	CreditCardErrorFormat   = "Not a valid credit card number."
	CSSColorErrorFormat     = "Not valid as css color."
	EmailErrorFormat        = "Not a valid email address."
	IPErrorFormat           = "Not a valid IP address."
	IPV4ErrorFormat         = "Not a valid IPv4 address."
	IPV6ErrorFormat         = "Not a valid IPv6 address."
	URLErrorFormat          = "Not a valid URL."
	UUIDErrorFormat         = "Not a valid UUID."
	UUID3ErrorFormat        = "Not a valid UUID (version 3)."
	UUID4ErrorFormat        = "Not a valid UUID (version 4)."
	UUID5ErrorFormat        = "Not a valid UUID (version 5)."
)
