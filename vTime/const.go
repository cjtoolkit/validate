package vTime

const (
	Type = "time"

	ParseErrorFormat     = "Unable to Parse Time"
	MandatoryErrorFormat = "Must not be left empty"
	MinErrorFormat       = "Cannot be less than '{{.min}}'"
	MaxErrorFormat       = "Cannot be more than '{{.max}}'"
	BetweenErrorFormat   = "Must be between '{{.min}}' and '{{.max}}"
)
