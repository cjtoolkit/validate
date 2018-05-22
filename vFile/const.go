package vFile

const (
	Type = "file"

	MandatoryErrorFormat  = "File is required."
	AcceptMimeErrorFormat = "Invalid Type, must be '{{.mimes|join \"or\"}}'."
	MaxSizeErrorFormat    = "File size should not be bigger than {{.maxBytes}}."
)
