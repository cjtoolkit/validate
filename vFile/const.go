package vFile

const (
	Type = "file"

	MandatoryErrorFormat  = "File is required"
	AcceptMimeErrorFormat = "Invalid Type, must be {{.mimes}}"
	MaxSizeErrorFormat    = "File size should not be bigger than {{.maxBytes}}"
)
