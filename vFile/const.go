package vFile

const (
	Type = "file"

	MandatoryFormat  = "File is required"
	AcceptMimeFormat = "Invalid Type, must be {{.mimes}}"
	MaxSizeFormat    = "File size should not be bigger than {{.maxBytes}}"
)
