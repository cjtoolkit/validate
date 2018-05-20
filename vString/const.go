package vString

const (
	Type = "string"

	MandatoryErrorFormat   = "Must not be left empty"
	PatternErrorFormat     = "Does not match pattern '{{.pattern}}'."
	MinRuneErrorFormat     = "Must have more than '{{.min}}' characters."
	MaxRuneErrorFormat     = "Must have less than '{{.max}}' characters."
	BetweenRuneErrorFormat = "Must be between {{.min}} and {{.max}} characters."
	MustMatchErrorFormat   = "Field does not match '{{.fieldName}}'"
)
