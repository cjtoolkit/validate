package vFile

import (
	"mime/multipart"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(value *multipart.FileHeader, hasError bool) error

func ValidateFile(file *multipart.FileHeader, rules ...ValidationRule) (*multipart.FileHeader, error) {
	value := file

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(value, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func MustValidateFile(file *multipart.FileHeader, rules ...ValidationRule) *multipart.FileHeader {
	value, err := ValidateFile(file, rules...)
	vError.Must(err)

	return value
}
