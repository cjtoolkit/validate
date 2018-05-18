package vString

const (
	Type = "string"

	MandatoryFormat   = "Must not be left empty"
	PatternFormat     = "Does not match pattern '{{.pattern}}'."
	MinRuneFormat     = "Must have more than '{{.min}}' characters."
	MaxRuneFormat     = "Must have less than '{{.max}}' characters."
	BetweenRuneFormat = "Must be between {{.min}} and {{.max}} characters."
	MustMatchFormat   = "Field does not match '{{.fieldName}}'"
)
