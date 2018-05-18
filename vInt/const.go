package vInt

const (
	Type = "int"

	NotAnIntFormat  = "{{.value}} is not an integer"
	MandatoryFormat = "Must not be left empty"
	MinFormat       = "Cannot be less than '{{.min}}'"
	MaxFormat       = "Cannot be more than '{{.max}}'"
)
