package vFile

import (
	"mime/multipart"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(value *multipart.FileHeader, hasError bool) error

func Validate(file *multipart.FileHeader, rules ...ValidationRule) (*multipart.FileHeader, error) {
	value := file

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(value, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func MustValidate(file *multipart.FileHeader, rules ...ValidationRule) *multipart.FileHeader {
	value, err := Validate(file, rules...)
	vError.Must(err)

	return value
}
