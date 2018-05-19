package vUint

const (
	Type = "uint"

	NotAnIntFormat  = "{{.value}} is not an integer"
	MandatoryFormat = "Must not be left empty"
	MinFormat       = "Cannot be less than '{{.min}}'"
	MaxFormat       = "Cannot be more than '{{.max}}'"
	BetweenFormat   = "Must be between '{{.min}}' and '{{.max}}"
	StepFormat      = "Must be in step of {{.step}}"
)
