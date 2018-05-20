package vTime

const (
	Type = "time"

	ParseErrorFormat     = "Unable to parse time."
	MandatoryErrorFormat = "Must not be left empty."
	MinErrorFormat       = "Cannot be less than '{{.min}}'."
	MaxErrorFormat       = "Cannot be more than '{{.max}}'."
	BetweenErrorFormat   = "Must be between '{{.min}}' and '{{.max}}'."
)
