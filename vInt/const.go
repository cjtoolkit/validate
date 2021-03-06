package vInt

const (
	Type = "int"

	NotAnIntErrorFormat  = "{{.value}} is not an integer."
	MandatoryErrorFormat = "Must not be left empty."
	MinErrorFormat       = "Cannot be less than '{{.min}}'."
	MaxErrorFormat       = "Cannot be more than '{{.max}}'."
	BetweenErrorFormat   = "Must be between '{{.min}}' and '{{.max}}'."
	StepErrorFormat      = "Must be in step of '{{.step}}'."
	MatchesErrorFormat   = "Is not in the list, valid matches are '{{.matches|join \"and\"}}'."
)
